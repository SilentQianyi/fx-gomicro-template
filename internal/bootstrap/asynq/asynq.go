package asynq

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
	"go.uber.org/fx"
	"time"
)

var Module = fx.Options(
	fx.Provide(NewClient),
	fx.Provide(NewScheduler),
	fx.Provide(NewServer),
)

func NewClient(loader config.Config) (*asynq.Client, error) {
	redisOpt, err := applyRedisUri(loader)
	if err != nil {
		return nil, err
	}
	return asynq.NewClient(*redisOpt), nil
}

func NewScheduler(loader config.Config, log logger.Logger) (*asynq.Scheduler, error) {
	redisOpt, err := applyRedisUri(loader)
	if err != nil {
		return nil, err
	}
	timeZone := loader.Get("asynq", "scheduler", "zone").String("Asia/Shanghai")
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}
	scheduler := asynq.NewScheduler(*redisOpt, &asynq.SchedulerOpts{
		Location: loc,
		Logger:   newAsynqLogger(log),
		LogLevel: asynq.InfoLevel,
	})
	return scheduler, nil
}

func NewServer(loader config.Config, log logger.Logger) (*asynq.Server, error) {
	redisOpt, err := applyRedisUri(loader)
	if err != nil {
		return nil, err
	}
	server := asynq.NewServer(*redisOpt, asynq.Config{
		Logger:   newAsynqLogger(log),
		LogLevel: asynq.InfoLevel,
	})
	return server, nil
}

func applyRedisUri(loader config.Config) (*asynq.RedisClientOpt, error) {
	redisUri := loader.Get("redis", "uri").String("redis://localhost:6379")
	opt, err := redis.ParseURL(redisUri)
	if err != nil {
		return nil, err
	}
	redisOpt := &asynq.RedisClientOpt{
		Network:      opt.Network,
		Addr:         opt.Addr,
		Username:     opt.Username,
		Password:     opt.Password,
		DB:           opt.DB,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		TLSConfig:    opt.TLSConfig,
	}
	return redisOpt, nil
}
