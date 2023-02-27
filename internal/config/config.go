package config

import (
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/env"
	"go-micro.dev/v4/config/source/file"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewServiceInfo),
)

func NewConfig() (config.Config, error) {
	conf, err := config.NewConfig(
		config.WithSource(file.NewSource(file.WithPath("./conf/config.yaml"))),
		config.WithSource(env.NewSource()),
		config.WithReader(
			json.NewReader( // json reader for internal config merge
				reader.WithEncoder(yaml.NewEncoder()),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	config.DefaultConfig = conf
	return conf, nil
}

func NewServiceInfo(conf config.Config) (*ServiceInfo, error) {
	var info ServiceInfo
	err := conf.Get("service").Scan(&info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

type ServiceInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Env     string `json:"env"`
}

func (s *ServiceInfo) Development() bool {
	return s.Env == "" || s.Env == "dev" || s.Env == "default"
}
