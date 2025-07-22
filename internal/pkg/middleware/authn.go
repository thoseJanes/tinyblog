package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
	"github.com/thoseJanes/tinyblog/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
		}else{
			c.Set(core.XUsernameKey, username)
			c.Next()
		}
	}
}