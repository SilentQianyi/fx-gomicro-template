package model

import (
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
)

type RabbitMqQueueConfig struct {
	Topic        string `json:"topic"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Durable      bool   `json:"durable"`
	AckOnSuccess bool   `json:"ack_on_success"`
}

type RabbitMqExchangeConfig struct {
	ExchangeName  string `json:"exchangeName"`
	ExchangeType  string `json:"exchangeType"`
	Durable       bool   `json:"durable"`
	AckOnSuccess  bool   `json:"ackOnSuccess"`
	PrefetchCount int    `json:"prefetch_count"`
}

type RabbitMqConfig struct {
	Uri string `json:"uri"`
}

func GetRabbitMqConfig(conf config.Config, log logger.Logger) *RabbitMqConfig {
	config := &RabbitMqConfig{}
	err := conf.Get("consumer").Scan(config)
	if err != nil {
		log.Logf(logger.ErrorLevel, "NewRabbitMq consumer config error! err[%s]", err.Error())
		return nil
	}
	return config
}

func GetRabbitMqExchangeConfig(conf config.Config, log logger.Logger, exchange string) *RabbitMqExchangeConfig {
	config := &RabbitMqExchangeConfig{}
	err := conf.Get("consumer", exchange).Scan(config)
	if err != nil {
		log.Logf(logger.ErrorLevel, "NewRabbitMq exchange config error! err[%s]", err.Error())
		return nil
	}
	return config
}

func GetRabbitMqQueueConfig(conf config.Config, log logger.Logger, exchange string, queue string) *RabbitMqQueueConfig {
	config := &RabbitMqQueueConfig{}
	err := conf.Get("consumer", exchange, queue).Scan(config)
	if err != nil {
		log.Logf(logger.ErrorLevel, "NewRabbitMq queue config error! err[%s]", err.Error())
		return nil
	}
	return config
}
