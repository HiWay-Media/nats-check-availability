package js_subscriber

import (
	"context"

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
	if configuration.NATS_CHECK_JETSTREAM {
		go s.Subscribe()
	}
	return s
}

func (s *JetStreamSubscriber) Subscribe() {
	s.logger.Debugf("[JetStreamSubscriber] Subscribe %v", s)
	/*js, err :=  s.js.Stream(context.Background(), s.configuration.NATS_DEFAULT_STREAM)
	if err != nil {
		s.logger.Fatalf("failed to get stream %s: %v", s.configuration.NATS_DEFAULT_STREAM, err)
		return
	}*/
	accountInfo, err := s.js.AccountInfo(context.Background())
	if err != nil {
		s.logger.Fatalf("failed to get account info: %v", err)
		return
	}
	s.logger.Infof("Account Info: %+v", accountInfo)
}
