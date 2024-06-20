package provider

import (
	"log"
	"os"
	"strconv"
)

type EnvProvider struct {
	appEnv           string
	serverPort       string
	databaseUrl      string
	databaseMaxConns int
	redisHost        string
	redisPort        string
	redisPassword    string
	logLevel         string
	sentryDsn        string
	sentryEnv        string
}

func (e *EnvProvider) AppEnv() string {
	return e.appEnv
}

func (e *EnvProvider) ServerPort() string {
	return e.serverPort
}

func NewEnvProvider(rootDir string) *EnvProvider {
	fallbackLookupEnv := func(key string, fallback string) string {
		value, exists := os.LookupEnv(key)
		if !exists {
			return fallback
		}
		return value
	}

	requireLookupEnv := func(key string) string {
		value, exists := os.LookupEnv(key)
		if !exists {
			log.Fatalf("[#kpkgirhm] '%s'", key)
		}
		return value
	}

	appServer := fallbackLookupEnv("APP_ENV", "local")
	serverPort := fallbackLookupEnv("SERVER_PORT", "3000")

	databaseUrl := requireLookupEnv("DATABASE_URL")
	redisHost := requireLookupEnv("REDIS_HOST")
	redisPort := requireLookupEnv("REDIS_PORT")
	redisPassword := fallbackLookupEnv("REDIS_PASSWORD", "")
	logLevel := requireLookupEnv("LOG_LEVEL")
	sentryDsn := fallbackLookupEnv("SENTRY_DSN", "")
	sentryEnv := fallbackLookupEnv("SENTRY_ENV", "local")

	databaseMaxConnsString := fallbackLookupEnv("DATABASE_MAX_CONNS", "5")
	parsedDatabaseMaxConns, err := strconv.Atoi(databaseMaxConnsString)
	if err != nil {
		log.Fatalf("[#la94puvm] '%s'", err)
	}

	envProvider := EnvProvider{
		appEnv:           appServer,
		serverPort:       serverPort,
		databaseMaxConns: parsedDatabaseMaxConns,
		databaseUrl:      databaseUrl,
		redisHost:        redisHost,
		redisPort:        redisPort,
		redisPassword:    redisPassword,
		logLevel:         logLevel,
		sentryDsn:        sentryDsn,
		sentryEnv:        sentryEnv,
	}

	return &envProvider
}
