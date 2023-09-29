package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
	"os"
	"quasar.fire.com/pkg/healthcheck"
	"quasar.fire.com/pkg/router"
	"quasar.fire.com/pkg/swagger"
	"quasar.fire.com/pkg/topsecret"
)

func main() {

	// Database instance
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Error opening database: %v", err)
	}

	// Migrations
	if err := db.AutoMigrate(&topsecret.Records{}); err != nil {
		logrus.Fatalf("Error migrating database: %v", err)
	}

	// Healthcheck
	healthcheckHandler := healthcheck.NewHandler()
	healthcheckRoutes := healthcheck.NewRoutes(healthcheckHandler)

	// TopSecret
	topsecretRepository := topsecret.NewRepository(db)
	topsecretService := topsecret.NewService(topsecretRepository)
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

	// Get port from environment
	port := os.Getenv("PORT")

	// Run server
	r := router.NewRouter(routes)
	logrus.Fatal(r.Run(fmt.Sprintf(":%v", port)))
}
