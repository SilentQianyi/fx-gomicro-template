package server

import (
	"chat/internal/bootstrap/rabbitmq"
	"context"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"os"

	"chat/internal/bootstrap/asynq"
)

var Module = fx.Options(
	fx.Provide(NewClient),
	fx.Provide(NewServer),
)

type GRPCManager struct {
	fx.In
	Server server.Server
	Client client.Client
}

func NewClient() client.Client {
	return grpc.NewClient()
}

func NewServer(db *mongo.Client, rpc GRPCManager, schedule asynq.Manager, group rabbitmq.Group, lifecycle fx.Lifecycle, log logger.Logger) (micro.Service, error) {
	opts, err := group.Init(lifecycle)
	if err != nil {
		log.Logf(logger.ErrorLevel, "NewServer group init error! err[ %s ]", err.Error())
		return nil, err
	}

	name, version := os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION")
	opts = append(opts, micro.Name(name))
	opts = append(opts, micro.Version(version))
	opts = append(opts, micro.Client(rpc.Client))
	opts = append(opts, micro.Server(rpc.Server))
	opts = append(opts, micro.AfterStart(func() error {
		if err = schedule.Start(); err != nil {
			return err
		}
		println("🎉 service started successfully!")
		return nil
	}))
	opts = append(opts, micro.AfterStop(func() error {
		err = schedule.Stop()
		err = multierr.Append(err, db.Disconnect(context.TODO()))
		println("👋 service shutdown.")
		return err
	}))
	srv := micro.NewService(
		opts...,
	)
	return srv, nil
}
