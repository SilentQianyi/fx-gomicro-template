package consumer

import (
	"chat/internal/bootstrap/rabbitmq"
	"chat/internal/domain/constant"
	"chat/internal/domain/model"
	"context"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
)

type helloWorldConsumer struct {
	rabbitmq.Consumer
}

func NewHelloWorldConsumer(conf config.Config, log logger.Logger) (rabbitmq.Subscriber, error) {
	consumer := &helloWorldConsumer{}
	consumer.Consume = consumer.Handle
	consumer.RabbitMqConf = model.GetRabbitMqConfig(conf, log)
	consumer.ExchangeConf = model.GetRabbitMqExchangeConfig(conf, log, constant.HelloWorldExchangeName)
	consumer.QueueConf = model.GetRabbitMqQueueConfig(conf, log, constant.HelloWorldExchangeName, constant.HelloWorldQueueName)

	return consumer, nil
}

func (c *helloWorldConsumer) Handle(ctx context.Context, event broker.Event) error {
	// TODO: 处理消息
	return nil
}
