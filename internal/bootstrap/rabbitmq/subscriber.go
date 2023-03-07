package rabbitmq

import (
	"chat/internal/domain/model"
	"chat/internal/util"
	"context"
	"errors"
	"fmt"
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/metadata"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	defaultPrefetchCount = 15
)

type Subscriber interface {
	Handle(context.Context, broker.Event) error
	Subscribe(lifecycle fx.Lifecycle) (broker.Broker, error)
}
type Consumer struct {
	RabbitMqConf *model.RabbitMqConfig
	ExchangeConf *model.RabbitMqExchangeConfig
	QueueConf    *model.RabbitMqQueueConfig
	broker       broker.Broker
	Consume      func(context.Context, broker.Event) error
}

func (c *Consumer) init(lifecycle fx.Lifecycle) error {
	if c.ExchangeConf.PrefetchCount == 0 {
		c.ExchangeConf.PrefetchCount = defaultPrefetchCount
	}
	opts := []broker.Option{broker.Addrs(c.RabbitMqConf.Uri), rabbitmq.PrefetchCount(c.ExchangeConf.PrefetchCount)}
	if c.ExchangeConf.ExchangeType == "" {
		opts = append(opts, rabbitmq.WithoutExchange())
	} else {
		opts = append(opts, rabbitmq.ExchangeName(c.ExchangeConf.ExchangeName), rabbitmq.ExchangeType(rabbitmq.MQExchangeType(c.ExchangeConf.ExchangeType)))
	}
	c.broker = rabbitmq.NewBroker(opts...)
	err := c.broker.Init()
	if err != nil {
		return err
	}
	if err = c.broker.Connect(); err != nil {
		return err
	}
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return c.broker.Disconnect()
		},
	})
	return nil
}

func (c *Consumer) Subscribe(lifecycle fx.Lifecycle) (broker.Broker, error) {
	if err := c.init(lifecycle); err != nil {
		return nil, err
	}
	if c.Consume == nil {
		return nil, errors.New("this consumer may not implemented Consume function")
	}
	handler := func(e broker.Event) error {
		zap.L().Info("receive raw message", zap.Any("header", e.Message().Header), zap.String("body", string(e.Message().Body)))
		ctx, header := util.CreateTraceableContext(), e.Message().Header
		return c.Consume(metadata.NewContext(ctx, header), e)
	}

	_, err := c.broker.Subscribe(c.QueueConf.Topic, handler, c.getSubscribeOptions()...)
	if err != nil {
		zap.L().Error(fmt.Sprintf("c.broker.Subscribe error[ %s ]", err.Error()))
		return nil, err
	}
	return c.broker, nil
}

func (c *Consumer) getSubscribeOptions() []broker.SubscribeOption {
	opts := []broker.SubscribeOption{broker.Queue(c.QueueConf.Name)}
	if c.QueueConf.Type != "" {
		opts = append(opts, rabbitmq.QueueArguments(map[string]any{"x-queue-type": c.QueueConf.Type}))
	}
	if c.QueueConf.Durable {
		opts = append(opts, rabbitmq.DurableQueue())
	}
	if c.QueueConf.AckOnSuccess {
		opts = append(opts, rabbitmq.AckOnSuccess())
	}
	return opts
}
