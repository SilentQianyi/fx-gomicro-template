package producer

import (
	"chat/internal/bootstrap/rabbitmq"
	"chat/internal/domain/constant"
	"chat/internal/domain/model"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
)

type helloWorldProducer struct {
	rabbitmq.Producer
}

func NewHelloWorldProducer(conf config.Config, log logger.Logger) (rabbitmq.Publisher, error) {
	producer := &helloWorldProducer{}
	producer.RabbitMqConf = model.GetRabbitMqConfig(conf, log)
	producer.ExchangeConf = model.GetRabbitMqExchangeConfig(conf, log, constant.HelloWorldExchangeName)
	producer.QueueConf = model.GetRabbitMqQueueConfig(conf, log, constant.HelloWorldExchangeName, constant.HelloWorldQueueName)

	return producer, nil
}
