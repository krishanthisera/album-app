package main

import (
	_ "github.com/krishanthisera/gitops-for-devs/cmd/api/docs"
	"github.com/krishanthisera/gitops-for-devs/pkg/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
    config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
    router.Use(cors.New(config))
	
	routes.SetupRoutes(router)
	
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080")
}
