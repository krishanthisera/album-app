package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/gitops-for-devs/pkg/handlers"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/albums", handlers.GetAlbums)
		v1.GET("/album/:id", handlers.GetAlbumByID)
		v1.POST("/album", handlers.PostAlbums)
	}
}
