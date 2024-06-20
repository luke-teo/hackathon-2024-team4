package provider

import (
	"log"

	"go.uber.org/zap"
)

func NewLoggerProvider(env *EnvProvider) *zap.Logger {
	var logger *zap.Logger
	var err error

	if env.logLevel == "debug" || env.logLevel == "test" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Panicln(err)
	}

	return logger
}
