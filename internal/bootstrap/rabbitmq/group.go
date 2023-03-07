package rabbitmq

import (
	"go-micro.dev/v4"
	"go.uber.org/fx"
)

type Group struct {
	fx.In
	Subscribers []Subscriber `group:"subscriber"`
	Publishers  []Publisher  `group:"publisher"`
}

func (g Group) Init(lifecycle fx.Lifecycle) ([]micro.Option, error) {
	opts := make([]micro.Option, 0)
	for _, subscriber := range g.Subscribers {
		brk, err := subscriber.Subscribe(lifecycle)
		if err != nil {
			return nil, err
		}
		opts = append(opts, micro.Broker(brk))
	}
	for _, publisher := range g.Publishers {
		brk, err := publisher.Publish(lifecycle)
		if err != nil {
			return nil, err
		}
		opts = append(opts, micro.Broker(brk))
	}
	return opts, nil
}
