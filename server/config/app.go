package config

import (
	"database/sql"
	"go_chi_template/config/provider"
	"path"
	"runtime"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type App struct {
	env     *provider.EnvProvider
	db      *sql.DB
	redis   *redis.Client
	rootDir string
	logger  *zap.Logger
	queue   *provider.AsynqProvider
	sentry  *sentryhttp.Handler
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

func (app *App) Sentry() *sentryhttp.Handler {
	if app.sentry == nil {
		app.sentry = provider.NewSentryProvider(app.env)
	}
	return app.sentry
}

func (app *App) setRootDir() {
	_, b, _, _ := runtime.Caller(0)
	app.rootDir = path.Join(path.Dir(b), "..")
}

func (app *App) UseTestDB() {
	app.db = provider.NewTestDbProvider(app.env)
}

func (app *App) UseTestQueue() {
	app.queue = provider.NewTestQueueProvider(app.env)
}

func NewApp() *App {
	app := App{}

	app.setRootDir()
	app.env = provider.NewEnvProvider(app.rootDir)
	provider.NewValidationProvider()
	app.Sentry()

	return &app
}
