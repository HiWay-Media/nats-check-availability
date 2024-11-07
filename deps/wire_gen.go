// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package deps

import (
	"github.com/HiWay-Media/hwm-go-utils/nats_helper"
	"github.com/HiWay-Media/nats-check-availability/app"
	"github.com/HiWay-Media/nats-check-availability/app/jetstream"
	"github.com/HiWay-Media/nats-check-availability/app/subscriber"
	"github.com/HiWay-Media/nats-check-availability/env"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func InjectApp(config *env.Configuration, logger *zap.SugaredLogger) (*app.App, error) {
	encodedConn, err := NewNatsClient(config, logger)
	if err != nil {
		return nil, err
	}
	natsSubscriber := subscriber.NewNatsSubscriber(config, logger, encodedConn)
	jetStream, err := NewNatsJetStreamClient(encodedConn, logger, config)
	if err != nil {
		return nil, err
	}
	jetStreamSubscriber := js_subscriber.NewJetstreamSubscriber(config, logger, jetStream)
	appApp := &app.App{
		NATSSubscriber:      natsSubscriber,
		JetStreamSubscriber: jetStreamSubscriber,
	}
	return appApp, nil
}

// wire.go:

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
