package service

import (
	"chat/internal/service/consumer"
	"chat/internal/service/grpc"
	"go.uber.org/fx"
)

var Module = fx.Options(
	grpc.Module,
	consumer.Module,
)
