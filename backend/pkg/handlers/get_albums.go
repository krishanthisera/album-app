package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/gitops-for-devs/pkg/data"
)

// GetAlbums     godoc
// @Summary      Get album array
// @Description  Responds with the list of albums as JSON.
// @Tags         albums
// @Produce      json
// @Success      200  {array}  models.Album
// @Router       /albums [get]
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(200, data.Albums)
}
