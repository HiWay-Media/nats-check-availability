package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Configuration struct {
	AppEnv                 	string 	`env:"APP_ENV"`
	LogLevel               	string 	`env:"LOG_LEVEL"`
	DbUsername             	string 	`env:"DB_USERNAME"`
	DbPassword             	string 	`env:"DB_PASSWORD"`
	DbHost                 	string 	`env:"DB_HOST"`
	DbPort                 	string 	`env:"DB_PORT"`
	DbIdleConn             	string 	`env:"DB_IDLE_CONN"`
	DbMaxConn              	string 	`env:"DB_MAX_CONN"`
	DbName                 	string 	`env:"DB_NAME"`
	NATS_SERVERS           	string 	`env:"NATS_SERVERS"`
	NATS_CHECK_PUB_SUB 		bool 	`env:"NATS_CHECK_PUB_SUB"`
	NATS_JETSTREAM_SERVERS 	string 	`env:"NATS_JETSTREAM_SERVERS"`
	NATS_CHECK_JETSTREAM 	bool 	`env:"NATS_CHECK_JETSTREAM"`
	NATS_REPLICA_SIZE      	string 	`env:"NATS_REPLICA_SIZE"`
	NATS_DEFAULT_CLUSTER   	string 	`env:"NATS_DEFAULT_CLUSTER"`
	NATS_DEFAULT_STREAM    	string 	`env:"NATS_DEFAULT_STREAM"`
}

func GetEnvConfig() *Configuration {
	setupEnv()
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}

func setupEnv() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		return
	}

	err := godotenv.Load(fmt.Sprintf("./env/.env.%s", strings.ToLower(appEnv)))
	if err == nil {
		return
	}

	_ = godotenv.Load()
}
