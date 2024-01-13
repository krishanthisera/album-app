package handlers

import "github.com/gin-gonic/gin"

type ResponseHandler interface {
	SendResponse(status int, body interface{})
}

type JSONResponse struct {
	Context *gin.Context
}

type XMLResponse struct {
	Context *gin.Context
}

func (x *XMLResponse) SendResponse(code int, body interface{}) {
	x.Context.XML(code, body)
}

func (grh *JSONResponse) SendResponse(status int, body interface{}) {
	grh.Context.IndentedJSON(status, body)
}
