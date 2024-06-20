package provider

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/redis/go-redis/v9"
)

func newRedisOpts(env *EnvProvider) redis.Options {
	redisUrl := env.redisHost + ":" + env.redisPort

	opts := redis.Options{
		Addr:     redisUrl,
		Password: env.redisPassword,
	}

	// always enable TLS unless in local or test ENVs
	if (env.appEnv != "local") && (env.appEnv != "test") {
		opts.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return opts
}

func NewRedisProvider(env *EnvProvider) *redis.Client {
	opts := newRedisOpts(env)
	rdb := redis.NewClient(&opts)

	err := rdb.Ping(context.Background()).Err()

	if err != nil {
		log.Fatalf("Unable to initialize redis connection: %s", opts.Addr)
	}

	return rdb
}
