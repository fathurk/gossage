package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrorBodyFormat     = "Wrong body format"
	ErrorInsertDocument = "Failed to insert document"
)

type GenericResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Responser(c *gin.Context, Code int, Message string) {
	c.JSON(Code, &GenericResponse{
		Code:    Code,
		Status:  http.StatusText(Code),
		Message: Message,
	})
}
