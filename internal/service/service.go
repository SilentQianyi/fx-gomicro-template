package service

import (
	"chat/internal/service/biz"
	"chat/internal/service/grpc"
	"go.uber.org/fx"
)

var Module = fx.Options(
	biz.Module,
	grpc.Module,
)
