// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/prologic/bitcask"

	petServer "github.com/CyrusJavan/terraform-provider-example/example-server/pet"

	"github.com/CyrusJavan/terraform-provider-example/example-server/restapi/operations"
	"github.com/CyrusJavan/terraform-provider-example/example-server/restapi/operations/pet"
)

//go:generate swagger generate server --target ../../example-server --name Pet --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.PetAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PetAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	db, _ := bitcask.Open("/tmp/petdb")
	ps := petServer.PetServer{
		DB: db,
	}

	api.PetAddPetHandler = pet.AddPetHandlerFunc(func(params pet.AddPetParams) middleware.Responder {
		return ps.AddPet(params.Body)
	})
	api.PetDeletePetHandler = pet.DeletePetHandlerFunc(func(params pet.DeletePetParams) middleware.Responder {
		return ps.DeletePet(params.PetID)
	})
	api.PetGetPetByIDHandler = pet.GetPetByIDHandlerFunc(func(params pet.GetPetByIDParams) middleware.Responder {
		return ps.GetPetByID(params.PetID)
	})
	api.PetUpdatePetHandler = pet.UpdatePetHandlerFunc(func(params pet.UpdatePetParams) middleware.Responder {
		return ps.AddPet(params.Body)
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
	return handler
}
