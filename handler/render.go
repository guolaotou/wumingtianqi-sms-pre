package handler

import (
	"github.com/gin-gonic/gin"
	"wumingtianqi/utils/errnum"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(context *gin.Context, err error, data interface{}) {
	code, message := errnum.DecodeErr(err)
	context.Writer.Header().Set("Content-Type", "application/json")

	response := &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	context.JSON(200, *response)
}
