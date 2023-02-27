package common

import (
	"fmt"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

func ProvideComponents(group string, constructors ...interface{}) fx.Option {
	opts := lo.Map(constructors, func(constructor interface{}, _ int) fx.Option {
		return fx.Provide(fx.Annotate(constructor, fx.ResultTags(fmt.Sprintf(`group:"%s"`, group))))
	})
	return fx.Options(opts...)
}
