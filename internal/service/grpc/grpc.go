package grpc

import (
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"time"
)

var Module = fx.Options(
	fx.Provide(NewServer),
	fx.Invoke(registerHelloWorldHandler),
)

func NewServer(conf config.Config, log logger.Logger) server.Server {
	port := conf.Get("server", "port").Int(10016)
	return grpc.NewServer(
		server.Address(fmt.Sprintf(":%d", port)),
		server.WrapHandler(handleError),
		server.WrapHandler(requestLog),
		server.WithLogger(log),
	)
}

func validateRequest(req interface{}) error {
	reqVal := reflect.ValueOf(req)
	validateMethod := reqVal.MethodByName("Validate")
	if !validateMethod.IsValid() || validateMethod.IsZero() {
		return nil
	}
	results := validateMethod.Call(nil)
	if results != nil && len(results) == 1 && !results[0].IsZero() {
		return results[0].Interface().(error)
	}
	return nil
}

func getTraceId(ctx context.Context) string {
	md, ok := metadata.FromContext(ctx)
	if ok {
		value, ok := md.Get("traceId")
		if ok {
			if value != "" {
				return value
			}
		}
	}
	return primitive.NewObjectID().Hex()
}

func requestLog(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		// 记录请求参数 耗时 错误信息等数据
		fields := map[string]any{
			"method":    req.Method(),
			"request":   req.Body(),
			"traceId":   getTraceId(ctx),
			"startTime": start.Format(time.RFC3339),
		}
		logger.Fields(fields).Log(logger.InfoLevel, "rpc request start:")
		err := handlerFunc(ctx, req, rsp)
		fields["requestTime"] = time.Now().Sub(start)
		if err != nil {
			logger.Fields(fields).Logf(logger.ErrorLevel, "rpc request error: %+v", err)
		} else {
			logger.Fields(fields).Log(logger.InfoLevel, "rpc request end.")
		}
		return err
	}
}

func handleError(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp any) error {
		if err := validateRequest(req); err != nil {
			return err
		}
		err := handlerFunc(ctx, req, rsp)
		if err == nil {
			return nil
		}
		_, ok := status.FromError(err)
		if ok {
			return err
		}
		if err == mongo.ErrNoDocuments {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
}
