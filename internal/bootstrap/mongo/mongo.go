package mongo

import (
	"context"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMongoClient),
	fx.Provide(NewMongoDatabase),
)

type dbConfig struct {
	Uri  string `json:"uri"`
	Ping bool   `json:"ping,omitempty"`
}

func NewMongoClient(loader config.Config, log logger.Logger) (*mongo.Client, error) {
	var conf dbConfig
	err := loader.Get("mongo").Scan(&conf)
	if err != nil {
		return nil, err
	}
	cli, err := mongo.NewClient(options.Client().ApplyURI(conf.Uri))
	if err != nil {
		return nil, err
	}
	if err := cli.Connect(context.TODO()); err != nil {
		return nil, err
	}
	log.Logf(logger.InfoLevel, "connected to [%s]", conf.Uri)
	if conf.Ping {
		if err := cli.Ping(context.TODO(), readpref.Primary()); err != nil {
			return nil, err
		}
	}
	return cli, nil
}

func NewMongoDatabase(loader config.Config, cli *mongo.Client) *mongo.Database {
	name := loader.Get("mongo", "database").String("test")
	return cli.Database(name)
}
