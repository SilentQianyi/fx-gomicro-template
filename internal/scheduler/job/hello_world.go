package schedulerJob

import (
	enumScheduler "chat/internal/domain/enum/scheduler"
	"github.com/hibiken/asynq"
	"time"
)

type HelloWorldJob struct {
	ScheduleTask
}

func newContactMetricsJob() ScheduleTask {
	return &HelloWorldJob{}
}

func (job *HelloWorldJob) Task() *asynq.Task {
	return asynq.NewTask(string(enumScheduler.HelloWorldJob), nil, asynq.Unique(time.Minute))
}

func (job *HelloWorldJob) CronSpec() string {
	return "4 19 * * ?"
}

func (job *HelloWorldJob) Name() string {
	return "HELLO_WORLD_JOB"
}
