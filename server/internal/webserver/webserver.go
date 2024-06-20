package webserver

import (
	"fmt"
	"go_chi_template/config"
	"go_chi_template/generated/oapi"
	"go_chi_template/internal/app/enum"
	"go_chi_template/internal/webserver/handler"
	"go_chi_template/internal/webserver/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/jwtauth"
)

type Webserver struct {
	router     *chi.Mux
	serverAddr string
}

func (ws *Webserver) Router() *chi.Mux {
	return ws.router
}

func NewWebserver(app *config.App) *Webserver {
	handler := handler.NewHandler(app)
	serverAddr := ":" + app.EnvVars().ServerPort()

	r := chi.NewRouter()
	r.Use(middleware.NewLoggerMiddleware(app.Logger()))
	r.Use(chimiddleware.Recoverer)
	r.Use(app.Sentry().Handle)
	// r.Use(middleware.NewAuthMiddleware(app))
	r.Get("/health", handler.GetHealth)

	baseURL := ""
	serverOptions := oapi.StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			errorEnum := enum.InternalRequestHandlerErrorEnum()
			middleware.ErrorResponseHandler(app.Logger(), w, r, err, errorEnum.Code)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			errorEnum := enum.InternalResponseHandlerErrorEnum()
			middleware.ErrorResponseHandler(app.Logger(), w, r, err, errorEnum.Code)
		},
	}
	strictHandler := oapi.NewStrictHandlerWithOptions(
		handler,
		[]oapi.StrictMiddlewareFunc{},
		serverOptions,
	)
	oapi.HandlerFromMuxWithBaseURL(strictHandler, r, baseURL)
	return &Webserver{
		router:     r,
		serverAddr: serverAddr,
	}
}

func (ws *Webserver) Start() {

	log.Print("WebServer listening on " + ws.serverAddr)

	s := &http.Server{
		Handler: ws.router,
		Addr:    ws.serverAddr,
	}

	log.Fatal(s.ListenAndServe())
}

func (ws *Webserver) PrintRoutes() {
	err := chi.Walk(
		ws.router,
		func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
			return nil
		},
	)

	if err != nil {
		log.Panicln(err)
	}
}
