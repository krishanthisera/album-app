package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/gitops-for-devs/pkg/data"
)

// GetAlbumByID                godoc
//
//	@Summary		Get single album by id
//	@Description	Returns the album whose ISBN value matches the ID.
//	@Tags			albums
//	@Accept			application/json
//	@accept			application/xml
//	@Produce		application/json
//	@produce		application/xml
//	@Param			id	path		string	true	"search album by ID"
//	@Success		200	{object}	models.Album
//	@Router			/album/{id} [get]
//
//	GetAlbumByID retrieves an album by its ID from the database and sends the response to the client.
//	It accepts a gin.Context object as a parameter, which contains the HTTP request and response information.
//	The ID of the album to retrieve is extracted from the request parameters.
//	The response format is determined based on the "Accept" header of the request.
//	If the "Accept" header is set to "application/xml", the response will be in XML format.
//	Otherwise, the response will be in JSON format.
//	If the album is found in the database, a 200 OK response is sent with the album details.
//	If the album is not found, a 404 Not Found response is sent with an error message.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var responseHandler ResponseHandler

	if c.Request.Header.Get("Accept") == "application/xml" {
		responseHandler = &XMLResponse{Context: c}
	} else {
		responseHandler = &JSONResponse{Context: c}
	}

	for _, a := range data.Albums {
		if a.ID == id {
			responseHandler.SendResponse(http.StatusOK, a)
			return
		}
	}
	responseHandler.SendResponse(http.StatusNotFound, gin.H{"message": "album not found"})
}
