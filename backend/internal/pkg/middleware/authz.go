package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
	"github.com/thoseJanes/tinyblog/pkg/auth"
)


type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}


func Authz(a *auth.Authz) gin.HandlerFunc {
	return func(c *gin.Context){
		sub := c.GetHeader(core.XUsernameKey)
		obj := c.Request.URL.Path
		act := c.Request.Method
		if allow, _ := a.Authorize(sub, obj, act); !allow{
			core.WriteResponse(c, errno.ErrUnauthorized, nil)
			c.Abort()
		}else{
			c.Next()
		}
	}
}