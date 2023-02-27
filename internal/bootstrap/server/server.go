package server

import (
	"context"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
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

func NewServer(db *mongo.Client, rpc GRPCManager, schedule asynq.Manager) (micro.Service, error) {
	name, version := os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION")
	srv := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Client(rpc.Client),
		micro.Server(rpc.Server),
		micro.AfterStart(func() error {
			if err := schedule.Start(); err != nil {
				return err
			}
			println("🎉 service started successfully!")
			return nil
		}),
		micro.AfterStop(func() error {
			err := schedule.Stop()
			err = multierr.Append(err, db.Disconnect(context.TODO()))
			println("👋 service shutdown.")
			return err
		}),
	)
	return srv, nil
}
