package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swaggerGin "github.com/swaggo/gin-swagger"
	_ "quasar.fire.com/pkg/swagger/docs" //documentation
	"quasar.fire.com/pkg/topsecret"
)

// Routes for swagger routes
type Routes struct{}

// NewRoutes create a new swagger routes instance
func NewRoutes() Routes {
	return Routes{}
}

// Register to register the routes
// Next line is for swagger documentation
// @title QUASAR-FIRE API Rest
// @version 1.0
// @description API for QUASAR-FIRE.
// @contact.name Anderson Rodriguez
// @contact.url https://www.linkedin.com/in/anderson-rodriguez-cer%C3%B3n-22aa28155/
// @contact.email andersonrodriguezce@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/quasar-fire/v1
//
//go:generate go run github.com/swaggo/swag/cmd/swag@v1.8.4 init --parseDependency=true -o docs -g routes.go
func (r *Routes) Register(group *gin.RouterGroup) {
	// Next line is for swagger documentation
	_ = topsecret.Handler{}
	group.GET("/swagger/*any", swaggerGin.WrapHandler(swaggerFiles.Handler))
}
