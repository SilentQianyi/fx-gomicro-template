package schedulerJob

import (
	"chat/internal/common"
	"github.com/hibiken/asynq"
	"go-micro.dev/v4/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	common.ProvideComponents(
		"schedule_task",
		newContactMetricsJob,
	),
	fx.Invoke(registerScheduleTask),
)

type ScheduleTask interface {
	CronSpec() string
	Name() string
	Task() *asynq.Task
}

type ScheduleTaskGroup struct {
	fx.In
	Items []ScheduleTask `group:"schedule_task"`
}

func registerScheduleTask(scheduler *asynq.Scheduler, group ScheduleTaskGroup) error {
	for _, item := range group.Items {
		logger.Infof("start register %s...", item.Name())
		entryId, err := scheduler.Register(item.CronSpec(), item.Task())
		if err != nil {
			return err
		}
		logger.Infof("register task(%s)", entryId)
	}
	return nil
}
