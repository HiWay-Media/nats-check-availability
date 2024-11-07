//go:build wireinject
// +build wireinject

package deps

import (
	"github.com/HiWay-Media/hwm-go-utils/nats_helper"
	"github.com/HiWay-Media/nats-check-availability/app"
	"github.com/HiWay-Media/nats-check-availability/app/subscriber"
	"github.com/HiWay-Media/nats-check-availability/env"
	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

func InjectApp(config *env.Configuration, logger *zap.SugaredLogger) (*app.App, error) {
	wire.Build(
		//brokers
		NewNatsClient,
		//NewNatsJetStreamClient,
		//subscriber
		subscriber.NewNatsSubscriber,
		// wire
		wire.Struct(new(app.App), "*"),
	)

	return nil, nil
}

func NewNatsClient(configuration *env.Configuration, logger *zap.SugaredLogger) (*nats.EncodedConn, error) {
	nc, err := nats_helper.NewNatsConn(configuration.NATS_SERVERS, logger)
	if err != nil {
		logger.Fatalf("failed to connect to nats server %s: %v", configuration.NATS_SERVERS, err)
		return nil, err
	}
	return nc, nil
}

func NewNatsJetStreamClient(nc *nats.EncodedConn, logger *zap.SugaredLogger, configuration *env.Configuration) (jetstream.JetStream, error) {
	js, err := nats_helper.NewNatsJetStream(nc, logger)
	if err != nil {
		logger.Fatalf("failed to connect to nats server %s: %v", configuration.NATS_SERVERS, err)
		return nil, err
	}

	return js, nil
}
