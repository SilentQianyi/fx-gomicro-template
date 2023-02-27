package handler

import (
	enumScheduler "chat/internal/domain/enum/scheduler"
	"context"
	"github.com/hibiken/asynq"
	"go-micro.dev/v4/logger"
	"time"
)

type HelloWorldJobHandler struct {
	client *asynq.Client
	//repo   db.HelloWorldRepository
	//service biz.MetricsService
	log *logger.Helper
}

func NewHelloWorldJobHandler(
	client *asynq.Client,
	//repo db.HelloWorldRepository,
	log logger.Logger,
) ScheduleTaskHandler {

	return &HelloWorldJobHandler{
		client: client,
		//repo:   repo,
		log: logger.NewHelper(log),
	}
}

func (h *HelloWorldJobHandler) Pattern() enumScheduler.ScheduleTaskType {
	return enumScheduler.HelloWorldJob
}

func (h *HelloWorldJobHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {
	h.log.Infof("start process hello world job in(%+v) ...", time.Now())
	//h.repo.GetHelloWorldByLanguage(ctx, )

	return nil
}
