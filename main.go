package main

import (
	app_log "github.com/HiWay-Media/nats-check-availability/app/lib/log"
	"github.com/HiWay-Media/nats-check-availability/deps"
	"github.com/HiWay-Media/nats-check-availability/env"
)

func main() {
	//
	config := env.GetEnvConfig()
	logger := app_log.NewLogger(config.LogLevel)
	app, err := deps.InjectApp(config, logger)
	if err != nil {
		logger.Fatal(err)
	}
	//
	err = app.Routes(config, logger).Run()
	logger.Error(err)
}
