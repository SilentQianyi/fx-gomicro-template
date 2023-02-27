package asynq

import (
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"go.uber.org/multierr"
)

type Manager struct {
	fx.In
	Client    *asynq.Client
	Scheduler *asynq.Scheduler
	Server    *asynq.Server
	Mux       *asynq.ServeMux
}

func (m Manager) Start() error {
	err := m.Scheduler.Start()
	return multierr.Append(err, m.Server.Start(m.Mux))
}

func (m Manager) Stop() error {
	m.Scheduler.Shutdown()
	err := m.Client.Close()
	m.Server.Stop()
	return err
}
