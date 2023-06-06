package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrorBodyFormat       = "Wrong body format"
	ErrorInsertDocument   = "Failed to insert document"
	ErrorInvalidEmail     = "Invalid email format"
	ErrorHasBeenUsedEmail = "Email has been used"
)

type GenericResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Responser(c *gin.Context, Code int, Message string, Data interface{}) {
	if Data == nil {
		c.JSON(Code, &GenericResponse{
			Code:    Code,
			Status:  http.StatusText(Code),
			Message: Message,
			Data:    nil,
		})
		return
	}
	c.JSON(Code, &GenericResponse{
		Code:    Code,
		Status:  http.StatusText(Code),
		Message: Message,
		Data:    Data,
	})
}
