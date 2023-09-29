package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"quasar.fire.com/pkg/healthcheck"
	"quasar.fire.com/pkg/router"
	"quasar.fire.com/pkg/swagger"
	"quasar.fire.com/pkg/topsecret"
)

func main() {
	//	fmt.Println(GetLocation(100, 115.5, 142.7))
	//	fmt.Println(GetMessage(
	//		[]string{"", "este", "es", "un", "mensaje"},
	//		[]string{"este", "", "un", "mensaje"},
	//		[]string{"", "", "es", "", "mensaje"},
	//	))

	// Healthcheck
	healthcheckHandler := healthcheck.NewHandler()
	healthcheckRoutes := healthcheck.NewRoutes(healthcheckHandler)

	// TopSecret
	topsecretService := topsecret.NewService()
	topsecretHandler := topsecret.NewHandler(topsecretService)
	topsecretRoutes := topsecret.NewRoutes(topsecretHandler)

	// Swagger
	swaggerRoutes := swagger.NewRoutes()

	// Routes
	routes := &router.RoutesGroup{
		HealthCheck: healthcheckRoutes,
		TopSecret:   topsecretRoutes,
		Swagger:     swaggerRoutes,
	}

	// Run server
	r := router.NewRouter(routes)
	port := os.Getenv("PORT")
	logrus.Fatal(r.Run(fmt.Sprintf(":%v", port)))
}
