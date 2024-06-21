package config

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"first_move/config/provider"
)

type App struct {
	env     *provider.EnvProvider
	db      *sql.DB
	redis   *redis.Client
	rootDir string
	logger  *zap.Logger
	queue   *provider.AsynqProvider
}

func (app *App) EnvVars() *provider.EnvProvider {
	return app.env
}

func (app *App) DB() *sql.DB {
	if app.db == nil {
		app.db = provider.NewDbProvider(app.env)
	}
	return app.db
}

func (app *App) Redis() *redis.Client {
	if app.redis == nil {
		app.redis = provider.NewRedisProvider(app.env)
	}
	return app.redis
}

func (app *App) Logger() *zap.Logger {
	if app.logger == nil {
		app.logger = provider.NewLoggerProvider(app.env)
	}
	return app.logger
}

func (app *App) Queue() *provider.AsynqProvider {
	if app.queue == nil {
		app.queue = provider.NewQueueProvider(app.Redis())
	}
	return app.queue
}

func NewApp() *App {
	app := App{}

	app.env = provider.NewEnvProvider()

	return &app
}
