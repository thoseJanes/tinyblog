package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
)


func RequestId(c *gin.Context) {
	requestId := c.Request.Header.Get(core.XRequestIdKey)
	if requestId == "" {
		requestId = uuid.New().String()
	}

	c.Writer.Header().Set(core.XRequestIdKey, requestId)
	c.Set(core.XRequestIdKey, requestId)
	c.Next()
}