package test

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/config/provider"
	"go_chi_template/generated/oapi"
	"go_chi_template/internal/webserver"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mock HttpClient to work against generated openapi client
type fakeHttpServer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type fakeHttpClient struct {
	server fakeHttpServer
}

func (c *fakeHttpClient) Do(r *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	c.server.ServeHTTP(rr, r)
	return rr.Result(), nil
}

func init() {
	// TODO: figure out if there is an elegant way to initalize txn DB without going through app
	provider.RegisterTestTxDb()
}

func SetupTestApp(t *testing.T) *config.App {
	app := config.NewApp()
	app.UseTestDB()
	app.UseTestQueue()

	return app
}

func SetupTestHttpClient(app *config.App) *fakeHttpClient {
	ws := webserver.NewWebserver(app)

	fakeClient := fakeHttpClient{
		server: ws.Router(),
	}

	return &fakeClient
}

func SetupTestOpenAPIClient(app *config.App) *oapi.ClientWithResponses {

	fakeClient := SetupTestHttpClient(app)
	client, err := oapi.NewClientWithResponses("", oapi.WithHTTPClient(fakeClient))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func SetupAuthHeaders(app *config.App) oapi.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		// TODO: placeholder, but you should add your test setup to generate auth headers here
		// req.Header.Add("Authorization", val)
		return nil
	}
}
