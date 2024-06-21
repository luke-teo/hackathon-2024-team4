// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package oapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Score Summary of user behavior on that day
type Score struct {
	CurrentScore      int                `json:"currentScore"`
	Date              openapi_types.Date `json:"date"`
	Mean              int                `json:"mean"`
	StandardDeviation int                `json:"standardDeviation"`
	ZScore            float32            `json:"zScore"`
}

// N400Error defines model for 400Error.
type N400Error struct {
	Data         *map[string]interface{} `json:"data"`
	ErrorCode    string                  `json:"errorCode"`
	ErrorMessage string                  `json:"errorMessage"`
}

// N500Error defines model for 500Error.
type N500Error struct {
	Data         *map[string]interface{} `json:"data"`
	ErrorCode    string                  `json:"errorCode"`
	ErrorMessage string                  `json:"errorMessage"`
}

// GetScoresByUserIDJSONBody defines parameters for GetScoresByUserID.
type GetScoresByUserIDJSONBody interface{}

// GetScoresByUserIDParams defines parameters for GetScoresByUserID.
type GetScoresByUserIDParams struct {
	// StartDate Start date
	StartDate openapi_types.Date `form:"startDate" json:"startDate"`

	// EndDate End date
	EndDate openapi_types.Date `form:"endDate" json:"endDate"`
}

// PostUploadCsvMultipartBody defines parameters for PostUploadCsv.
type PostUploadCsvMultipartBody struct {
	File     *openapi_types.File `json:"file,omitempty"`
	Filename *string             `json:"filename,omitempty"`
}

// GetScoresByUserIDJSONRequestBody defines body for GetScoresByUserID for application/json ContentType.
type GetScoresByUserIDJSONRequestBody GetScoresByUserIDJSONBody

// PostUploadCsvMultipartRequestBody defines body for PostUploadCsv for multipart/form-data ContentType.
type PostUploadCsvMultipartRequestBody PostUploadCsvMultipartBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetScoresByUserIDWithBody request with any body
	GetScoresByUserIDWithBody(ctx context.Context, userId string, params *GetScoresByUserIDParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GetScoresByUserID(ctx context.Context, userId string, params *GetScoresByUserIDParams, body GetScoresByUserIDJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostUploadCsvWithBody request with any body
	PostUploadCsvWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetScoresByUserIDWithBody(ctx context.Context, userId string, params *GetScoresByUserIDParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetScoresByUserIDRequestWithBody(c.Server, userId, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetScoresByUserID(ctx context.Context, userId string, params *GetScoresByUserIDParams, body GetScoresByUserIDJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetScoresByUserIDRequest(c.Server, userId, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUploadCsvWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUploadCsvRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetScoresByUserIDRequest calls the generic GetScoresByUserID builder with application/json body
func NewGetScoresByUserIDRequest(server string, userId string, params *GetScoresByUserIDParams, body GetScoresByUserIDJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetScoresByUserIDRequestWithBody(server, userId, params, "application/json", bodyReader)
}

// NewGetScoresByUserIDRequestWithBody generates requests for GetScoresByUserID with any type of body
func NewGetScoresByUserIDRequestWithBody(server string, userId string, params *GetScoresByUserIDParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/scores/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "startDate", runtime.ParamLocationQuery, params.StartDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "endDate", runtime.ParamLocationQuery, params.EndDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPostUploadCsvRequestWithBody generates requests for PostUploadCsv with any type of body
func NewPostUploadCsvRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/upload_csv")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetScoresByUserIDWithBodyWithResponse request with any body
	GetScoresByUserIDWithBodyWithResponse(ctx context.Context, userId string, params *GetScoresByUserIDParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetScoresByUserIDResponse, error)

	GetScoresByUserIDWithResponse(ctx context.Context, userId string, params *GetScoresByUserIDParams, body GetScoresByUserIDJSONRequestBody, reqEditors ...RequestEditorFn) (*GetScoresByUserIDResponse, error)

	// PostUploadCsvWithBodyWithResponse request with any body
	PostUploadCsvWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUploadCsvResponse, error)
}

type GetScoresByUserIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Scores []Score `json:"scores"`
		UserId string  `json:"userId"`
	}
	JSON400 *N400Error
	JSON500 *N500Error
}

// Status returns HTTPResponse.Status
func (r GetScoresByUserIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetScoresByUserIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostUploadCsvResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *N400Error
	JSON500      *N500Error
}

// Status returns HTTPResponse.Status
func (r PostUploadCsvResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostUploadCsvResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetScoresByUserIDWithBodyWithResponse request with arbitrary body returning *GetScoresByUserIDResponse
func (c *ClientWithResponses) GetScoresByUserIDWithBodyWithResponse(ctx context.Context, userId string, params *GetScoresByUserIDParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetScoresByUserIDResponse, error) {
	rsp, err := c.GetScoresByUserIDWithBody(ctx, userId, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetScoresByUserIDResponse(rsp)
}

func (c *ClientWithResponses) GetScoresByUserIDWithResponse(ctx context.Context, userId string, params *GetScoresByUserIDParams, body GetScoresByUserIDJSONRequestBody, reqEditors ...RequestEditorFn) (*GetScoresByUserIDResponse, error) {
	rsp, err := c.GetScoresByUserID(ctx, userId, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetScoresByUserIDResponse(rsp)
}

// PostUploadCsvWithBodyWithResponse request with arbitrary body returning *PostUploadCsvResponse
func (c *ClientWithResponses) PostUploadCsvWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUploadCsvResponse, error) {
	rsp, err := c.PostUploadCsvWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUploadCsvResponse(rsp)
}

// ParseGetScoresByUserIDResponse parses an HTTP response from a GetScoresByUserIDWithResponse call
func ParseGetScoresByUserIDResponse(rsp *http.Response) (*GetScoresByUserIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetScoresByUserIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Scores []Score `json:"scores"`
			UserId string  `json:"userId"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostUploadCsvResponse parses an HTTP response from a PostUploadCsvWithResponse call
func ParsePostUploadCsvResponse(rsp *http.Response) (*PostUploadCsvResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostUploadCsvResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /scores/{userId})
	GetScoresByUserID(w http.ResponseWriter, r *http.Request, userId string, params GetScoresByUserIDParams)

	// (POST /upload_csv)
	PostUploadCsv(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /scores/{userId})
func (_ Unimplemented) GetScoresByUserID(w http.ResponseWriter, r *http.Request, userId string, params GetScoresByUserIDParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /upload_csv)
func (_ Unimplemented) PostUploadCsv(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetScoresByUserID operation middleware
func (siw *ServerInterfaceWrapper) GetScoresByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "userId", chi.URLParam(r, "userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScoresByUserIDParams

	// ------------- Required query parameter "startDate" -------------

	if paramValue := r.URL.Query().Get("startDate"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "startDate"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "startDate", r.URL.Query(), &params.StartDate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "startDate", Err: err})
		return
	}

	// ------------- Required query parameter "endDate" -------------

	if paramValue := r.URL.Query().Get("endDate"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "endDate"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "endDate", r.URL.Query(), &params.EndDate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "endDate", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetScoresByUserID(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostUploadCsv operation middleware
func (siw *ServerInterfaceWrapper) PostUploadCsv(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUploadCsv(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/scores/{userId}", wrapper.GetScoresByUserID)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/upload_csv", wrapper.PostUploadCsv)
	})

	return r
}

type N400ErrorJSONResponse struct {
	Data         *map[string]interface{} `json:"data"`
	ErrorCode    string                  `json:"errorCode"`
	ErrorMessage string                  `json:"errorMessage"`
}

type N500ErrorJSONResponse struct {
	Data         *map[string]interface{} `json:"data"`
	ErrorCode    string                  `json:"errorCode"`
	ErrorMessage string                  `json:"errorMessage"`
}

type GetScoresByUserIDRequestObject struct {
	UserId string `json:"userId"`
	Params GetScoresByUserIDParams
	Body   *GetScoresByUserIDJSONRequestBody
}

type GetScoresByUserIDResponseObject interface {
	VisitGetScoresByUserIDResponse(w http.ResponseWriter) error
}

type GetScoresByUserID200JSONResponse struct {
	Scores []Score `json:"scores"`
	UserId string  `json:"userId"`
}

func (response GetScoresByUserID200JSONResponse) VisitGetScoresByUserIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetScoresByUserID400JSONResponse struct{ N400ErrorJSONResponse }

func (response GetScoresByUserID400JSONResponse) VisitGetScoresByUserIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetScoresByUserID500JSONResponse struct{ N500ErrorJSONResponse }

func (response GetScoresByUserID500JSONResponse) VisitGetScoresByUserIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostUploadCsvRequestObject struct {
	Body *multipart.Reader
}

type PostUploadCsvResponseObject interface {
	VisitPostUploadCsvResponse(w http.ResponseWriter) error
}

type PostUploadCsv200Response struct {
}

func (response PostUploadCsv200Response) VisitPostUploadCsvResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostUploadCsv400JSONResponse struct{ N400ErrorJSONResponse }

func (response PostUploadCsv400JSONResponse) VisitPostUploadCsvResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostUploadCsv500JSONResponse struct{ N500ErrorJSONResponse }

func (response PostUploadCsv500JSONResponse) VisitPostUploadCsvResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /scores/{userId})
	GetScoresByUserID(ctx context.Context, request GetScoresByUserIDRequestObject) (GetScoresByUserIDResponseObject, error)

	// (POST /upload_csv)
	PostUploadCsv(ctx context.Context, request PostUploadCsvRequestObject) (PostUploadCsvResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetScoresByUserID operation middleware
func (sh *strictHandler) GetScoresByUserID(w http.ResponseWriter, r *http.Request, userId string, params GetScoresByUserIDParams) {
	var request GetScoresByUserIDRequestObject

	request.UserId = userId
	request.Params = params

	var body GetScoresByUserIDJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetScoresByUserID(ctx, request.(GetScoresByUserIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetScoresByUserID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetScoresByUserIDResponseObject); ok {
		if err := validResponse.VisitGetScoresByUserIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostUploadCsv operation middleware
func (sh *strictHandler) PostUploadCsv(w http.ResponseWriter, r *http.Request) {
	var request PostUploadCsvRequestObject

	if reader, err := r.MultipartReader(); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode multipart body: %w", err))
		return
	} else {
		request.Body = reader
	}

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostUploadCsv(ctx, request.(PostUploadCsvRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUploadCsv")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostUploadCsvResponseObject); ok {
		if err := validResponse.VisitPostUploadCsvResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
