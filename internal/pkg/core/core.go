package core

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
)


const(
	XRequestIdKey = "X-Request-Id"
	XUsernameKey = "X-Username"
)


type ErrResponse struct{
	Code string `json:"code"`
	Message string `json:"message"`
}


func WriteResponse(c *gin.Context, err error, data interface{}){
	httpCode := errno.Ok.HttpCode
	if err != nil {
		var errCode, message string
		httpCode, errCode, message = errno.Decode(err)
		data = ErrResponse{
			Code: errCode,
			Message: message,
		}
	}

	c.JSON(httpCode, data)
}
