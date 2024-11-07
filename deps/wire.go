//go:build wireinject
// +build wireinject

package deps

import (
	"github.com/HiWay-Media/nats-check-availability/app"
	"github.com/HiWay-Media/nats-check-availability/env"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InjectApp(config *env.Configuration, logger *zap.SugaredLogger) (*app.App, error) {
	wire.Build(

		// wire
		wire.Struct(new(app.App), "*"),
	)

	return nil, nil
}
