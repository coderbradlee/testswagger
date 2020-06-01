// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/lzxm160/testswagger/handler"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/lzxm160/testswagger/restapi/operations"
	"github.com/lzxm160/testswagger/restapi/operations/get"
	"github.com/lzxm160/testswagger/restapi/operations/update"
)

//go:generate swagger generate server --target ../../testswagger --name Did --spec ../swagger.yml

func configureFlags(api *operations.DidAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DidAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//if api.GetGetHandler == nil {
	api.GetGetHandler = get.GetHandlerFunc(func(params get.GetParams) middleware.Responder {
		return middleware.NotImplemented("operation get.Get has not yet been implemented")
	})
	//}
	fmt.Println("before nil", api.UpdateUpdateHandler)
	//if api.UpdateUpdateHandler == nil {
	api.UpdateUpdateHandler = update.UpdateHandlerFunc(func(params update.UpdateParams) middleware.Responder {
		ret := handler.UpdateHandler(params)
		return handler.UpdateResponse{ret}
	})
	//}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
