package consumer

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(fx.Annotate(NewHelloWorldConsumer, fx.ResultTags(`group:"subscriber"`))),
)
