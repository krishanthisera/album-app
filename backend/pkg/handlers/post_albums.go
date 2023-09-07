package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/gitops-for-devs/pkg/data"
	"github.com/krishanthisera/gitops-for-devs/pkg/models"
)

// PostAlbum             godoc
// @Summary      Store a new album
// @Description  Takes an album JSON and store.
// @Tags         albums
// @Produce      json
// @Param        album  body      models.Album  true  "Album JSON"
// @Success      201   {object}  models.Album
// @Router       /album [post]
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		println(err.Error())
		return
	}

	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(201, newAlbum)
}
