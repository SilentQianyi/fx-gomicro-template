package rabbitmq

import (
	"chat/internal/domain/event"
	"chat/internal/domain/model"
	"chat/internal/util"
	"context"
	"encoding/json"
	rabbitmq2 "github.com/go-micro/plugins/v4/broker/rabbitmq"
	"go-micro.dev/v4/broker"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Publisher interface {
	Publish(lifecycle fx.Lifecycle) (broker.Broker, error)
}
type Producer struct {
	RabbitMqConf *model.RabbitMqConfig
	ExchangeConf *model.RabbitMqExchangeConfig
	QueueConf    *model.RabbitMqQueueConfig
	broker       broker.Broker
}

func (p *Producer) init(lifecycle fx.Lifecycle) error {
	if p.ExchangeConf.PrefetchCount == 0 {
		p.ExchangeConf.PrefetchCount = defaultPrefetchCount
	}
	opts := []broker.Option{broker.Addrs(p.RabbitMqConf.Uri)}
	if p.ExchangeConf.ExchangeType == "" {
		opts = append(opts, rabbitmq2.WithoutExchange())
	} else {
		opts = append(opts, rabbitmq2.ExchangeName(p.ExchangeConf.ExchangeName), rabbitmq2.ExchangeType(rabbitmq2.MQExchangeType(p.ExchangeConf.ExchangeType)))
	}
	if p.ExchangeConf.Durable {
		opts = append(opts, rabbitmq2.DurableExchange())
	}
	p.broker = rabbitmq2.NewBroker(opts...)
	err := p.broker.Init()
	if err != nil {
		return err
	}
	if err = p.broker.Connect(); err != nil {
		return err
	}
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return p.broker.Disconnect()
		},
	})
	return nil
}

func (p *Producer) Publish(lifecycle fx.Lifecycle) (broker.Broker, error) {
	if err := p.init(lifecycle); err != nil {
		return nil, err
	}
	return p.broker, nil
}

func (p *Producer) getSubscribeOptions() []broker.SubscribeOption {
	opts := []broker.SubscribeOption{broker.Queue(p.QueueConf.Name)}
	if p.QueueConf.Type != "" {
		opts = append(opts, rabbitmq2.QueueArguments(map[string]any{"x-queue-type": p.QueueConf.Type}))
	}
	if p.QueueConf.Durable {
		opts = append(opts, rabbitmq2.DurableQueue())
	}
	if p.QueueConf.AckOnSuccess {
		opts = append(opts, rabbitmq2.AckOnSuccess())
	}
	return opts
}

func (p *Producer) SendMessage(ctx context.Context, msg *event.HelloWorldMessageEvent) error {
	data, err := json.Marshal(msg)
	if err != nil {
		zap.L().Error("marshal message error:", util.WithTraceField(ctx), util.WithAgentField(ctx), zap.Error(err))
		return err
	}
	zap.L().Info("send message:", util.WithTraceField(ctx), util.WithAgentField(ctx), zap.String("data", string(data)))
	return p.broker.Publish(
		p.QueueConf.Topic,
		&broker.Message{
			Header: map[string]string{
				"x-trace-id": util.GetTraceId(ctx),
				"x-agent-id": util.GetAgentId(ctx),
			},
			Body: data,
		},
	)
}
