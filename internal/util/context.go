package util

import (
	"context"
	"go-micro.dev/v4/metadata"
	"go.uber.org/zap"
)

const (
	traceIdKey = "x-trace-id"
	agentIdKey = "x-agent-id"
)

func CreateTraceableContext() context.Context {
	md := metadata.Metadata{}
	md.Set(traceIdKey, GetUUID())
	return metadata.NewContext(context.TODO(), md)
}

func getValue(ctx context.Context, key string) string {
	md, ok := metadata.FromContext(ctx)
	if ok {
		id, ok := md.Get(key)
		if ok {
			return id
		}
	}
	return ""
}

func GetTraceId(ctx context.Context) string {
	return getValue(ctx, traceIdKey)
}

func GetAgentId(ctx context.Context) string {
	return getValue(ctx, agentIdKey)
}

func WithTraceField(ctx context.Context) zap.Field {
	return zap.String(traceIdKey, GetTraceId(ctx))
}

func WithAgentField(ctx context.Context) zap.Field {
	return zap.String(agentIdKey, GetAgentId(ctx))
}
