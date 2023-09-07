package main

import (
	_ "github.com/krishanthisera/gitops-for-devs/cmd/api/docs"
	"github.com/krishanthisera/gitops-for-devs/pkg/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Album API
// @version         1.0
// @description     A album management service API in Go using Gin framework.

// @contact.name   Krishan Thisera
// @contact.url    https://github.com/krishanthisera/gitops-for-devs/issues
// @contact.email  @krishanthisera

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	router := routes.SetupRoutes()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080")
}
