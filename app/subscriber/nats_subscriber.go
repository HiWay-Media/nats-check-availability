package subscriber

import (
	"time"

	"github.com/HiWay-Media/nats-check-availability/env"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type NatsSubscriber struct {
	configuration *env.Configuration
	logger        *zap.SugaredLogger
	nc            *nats.EncodedConn
}

func NewNatsSubscriber(configuration *env.Configuration, logger *zap.SugaredLogger, nc *nats.EncodedConn) *NatsSubscriber {
	s := &NatsSubscriber{
		configuration: configuration,
		logger:        logger,
		nc:            nc,
	}
	if config.NATS_CHECK_PUB_SUB{
		go s.Subscribe()
	}
	return s
}

func (s *NatsSubscriber) Subscribe() {
	//
	s.logger.Debugf("[JetStreamSubscriber] Subscribe %v", s)
	for {
		_, err := s.nc.Subscribe(s.configuration.NATS_DEFAULT_STREAM, s.handle)
		if err != nil {
			s.logger.Errorf("[NatsSubscriber]failed to subscribe to nats subject: %s!", s.configuration.NATS_DEFAULT_STREAM)
			time.Sleep(time.Second * 3)
			continue
		}
		return
	}
}

func (s *NatsSubscriber) handle(request *any) {
	s.logger.Infof("received message: %v", request)
}
