package scheduler

import (
	"chat/internal/scheduler/handler"
	schedulerJob "chat/internal/scheduler/job"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	schedulerJob.Module,
)
