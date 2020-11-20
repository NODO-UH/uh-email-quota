// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/stdevAdrianPaez/uh-email-quota/models"
	"github.com/stdevAdrianPaez/uh-email-quota/quota"
	"github.com/stdevAdrianPaez/uh-email-quota/restapi/operations"
)

//go:generate swagger generate server --target ../../uh-email-quota --name GestionEmailPlugin --spec ../swagger-spec/swagger.yml --principal interface{}

func configureFlags(api *operations.GestionEmailPluginAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GestionEmailPluginAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.UserQuotaHandler = operations.UserQuotaHandlerFunc(func(params operations.UserQuotaParams) middleware.Responder {
		user := params.UserEmail.String()
		quota, err := quota.GetUserQuota(user)
		if err != nil {
			message := err.Error()
			return operations.NewUserQuotaDefault(500).WithPayload(&models.Error{Code: 500, Message: &message})
		}
		return operations.NewUserQuotaOK().WithPayload(&models.Quota{Value: &quota.Value, Limit: &quota.Limit})
	})

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
