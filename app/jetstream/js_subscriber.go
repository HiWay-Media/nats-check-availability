package js_subscriber

import (
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

func NewJetstreamSubscriber(configuration *env.Configuration, logger *zap.SugaredLogger,  js jetstream.JetStream) *JetStreamSubscriber{
    s := &JetStreamSubscriber{
		configuration: configuration,
		logger:        logger,
		js:            js,
	}
	go s.Subscribe()
	return s
}


func (s *JetStreamSubscriber) Subscribe() {
	for {
        s.logger.Debugf("Subscribe %v", s)
    }
}