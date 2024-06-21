package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"first_move/config"
	"first_move/generated/oapi"
	"first_move/internal/webserver/handler"
	"first_move/internal/webserver/middleware"
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

	baseURL := ""
	serverOptions := oapi.StrictHTTPServerOptions{}
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
