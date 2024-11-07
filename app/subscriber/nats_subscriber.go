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

func NewNatsSubscriber(nc *nats.EncodedConn, configuration *env.Configuration, logger *zap.SugaredLogger) *NatsSubscriber {
	s := &NatsSubscriber{
		configuration: configuration,
		logger:        logger,
		nc:            nc,
	}
	//go s.Subscribe()
	return s
}

func (s *NatsSubscriber) Subscribe() {
	for {
		_, err := s.nc.Subscribe(s.configuration.NATS_DEFAULT_STREAM, s.handle)
		if err != nil {
			s.logger.Errorf("failed to subscribe to encoding progress topic!")
			time.Sleep(time.Second * 3)
			continue
		}
		return
	}
}

func (s *NatsSubscriber) handle(request *any) {
	s.logger.Infof("received message: %v", request)
}
