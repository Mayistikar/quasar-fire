package router

import (
	"github.com/gin-gonic/gin"
	"quasar.fire.com/pkg/healthcheck"
	"quasar.fire.com/pkg/swagger"
	"quasar.fire.com/pkg/topsecret"
)

// Router interface for router implementation
type Router interface {
	Run(addr ...string) error
}

// NewRouter create a new router instance with all routes using gin
func NewRouter(routes *RoutesGroup) Router {
	path := "api/quasar-fire/v1"

	// Setting router with gin
	router := gin.Default()

	// Creating routes group
	groupV1 := router.Group(path)

	// Registering routes
	routes.HealthCheck.RegisterRoutes(groupV1)
	routes.TopSecret.RegisterRoutes(groupV1)
	routes.Swagger.Register(groupV1)

	return router
}

// RoutesGroup for unify all routes
type RoutesGroup struct {
	HealthCheck healthcheck.Routes
	TopSecret   topsecret.Routes
	Swagger     swagger.Routes
}
