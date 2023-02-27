package handler

import (
	"chat/internal/common"
	enumScheduler "chat/internal/domain/enum/scheduler"
	"github.com/hibiken/asynq"
	"go-micro.dev/v4/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	common.ProvideComponents(
		"schedule_task_handler",
		NewHelloWorldJobHandler,
	),
	fx.Provide(NewServeMux),
)

type ScheduleTaskHandler interface {
	asynq.Handler
	Pattern() enumScheduler.ScheduleTaskType
}

type ScheduleTaskHandlerGroup struct {
	fx.In
	Items []ScheduleTaskHandler `group:"schedule_task_handler"`
}

func NewServeMux(group ScheduleTaskHandlerGroup) *asynq.ServeMux {
	mux := asynq.NewServeMux()
	for _, item := range group.Items {
		logger.Infof("register scheduler task handler: (%s)", item.Pattern())
		mux.Handle(string(item.Pattern()), item)
	}
	return mux
}
