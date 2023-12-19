package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/gitops-for-devs/pkg/data"
	"github.com/krishanthisera/gitops-for-devs/pkg/models"
)

// PostAlbum             godoc
// @Summary		Store a new album
// @Description	Takes an album JSON and store.
// @Tags			albums
// @Accept			application/json
// @accept			application/xml
// @Produce		application/json
// @produce		application/xml
// @Param			album	body		models.Album	true	"Album JSON"
// @Success		201		{object}	models.Album
// @Router			/album [post]
//
// PostAlbums handles the HTTP POST request for creating new albums.
// It checks the content type of the request and binds the request body to a new album object.
// If the content type is JSON, it uses the JSONResponse handler to send the response.
// If the content type is XML, it uses the XMLResponse handler to send the response.
// If the content type is neither JSON nor XML, it sends a response with status code 204 (No Content).
// Finally, it appends the new album to the data.Albums slice and sends a response with status code 201 (Created).
func PostAlbums(c *gin.Context) {

	var responseHandler ResponseHandler

	var newAlbum models.Album

	contentType := c.Request.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		responseHandler = &JSONResponse{Context: c}
		if err := c.BindJSON(&newAlbum); err != nil {
			println(err.Error())
			responseHandler.SendResponse(http.StatusBadRequest, gin.H{"message": err})
			return
		}
	case "application/xml":
		responseHandler = &XMLResponse{Context: c}
		if err := c.BindXML(&newAlbum); err != nil {
			println(err.Error())
			responseHandler.SendResponse(http.StatusBadRequest, gin.H{"message": err})
			return
		}
	default:
		responseHandler = &JSONResponse{Context: c}
		responseHandler.SendResponse(http.StatusNoContent, gin.H{"message": "invalid content type"})
	}

	data.Albums = append(data.Albums, newAlbum)

	responseHandler.SendResponse(http.StatusCreated, newAlbum)
}
