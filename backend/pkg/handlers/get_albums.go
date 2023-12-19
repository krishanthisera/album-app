package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/gitops-for-devs/pkg/data"
)

// GetAlbums     godoc
// @Summary		Get album array
// @Description	Responds with the list of albums as JSON.
// @Tags			albums
// @Success		200	{array}	models.Album
// @Router			/albums [get]
//
// GetAlbums is a handler function that retrieves albums and sends the response based on the Accept header.
// If the Accept header is "application/xml", it sends the response in XML format.
// Otherwise, it sends the response in JSON format.
func GetAlbums(c *gin.Context) {
	var responseHandler ResponseHandler

	if c.Request.Header.Get("Accept") == "application/xml" {
		responseHandler = &XMLResponse{Context: c}
	} else {
		responseHandler = &JSONResponse{Context: c}
	}
	responseHandler.SendResponse(http.StatusOK, data.Albums)
}
