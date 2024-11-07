package js_subscriber

import (
	"github.com/HiWay-Media/nats-check-availability/env"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

type Subscriber interface {
	Subscribe()
}

type JetStreamSubscriber struct {
	configuration *env.Configuration
	logger        *zap.SugaredLogger
	js            jetstream.JetStream
}

func NewJetstreamSubscriber(configuration *env.Configuration, logger *zap.SugaredLogger, js jetstream.JetStream) *JetStreamSubscriber {
	s := &JetStreamSubscriber{
		configuration: configuration,
		logger:        logger,
		js:            js,
	}
	if configuration.NATS_CHECK_JETSTREAM{
		go s.Subscribe()
	}
	return s
}

func (s *JetStreamSubscriber) Subscribe() {
	s.logger.Debugf("[JetStreamSubscriber] Subscribe %v", s)
	for {
	}
}
