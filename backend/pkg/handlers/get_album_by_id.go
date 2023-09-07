package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/gitops-for-devs/pkg/data"
)

// GetAlbumByID                godoc
// @Summary      Get single album by id
// @Description  Returns the album whose ISBN value matches the ID.
// @Tags         albums
// @Produce      json
// @Param        id  path      string  true  "search album by ID"
// @Success      200  {object}  models.Album
// @Router       /albums/{id} [get]
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range data.Albums {
		if a.ID == id {
			c.IndentedJSON(200, a)
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "album not found"})
}
