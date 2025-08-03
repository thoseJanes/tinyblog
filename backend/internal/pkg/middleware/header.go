package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func NoCache(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Writer.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}


func CORS(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Writer.Header().Set("Allow", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}else{
		c.Next()
	}
}

func Secure(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("X-Frame-Options", "DENY")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000") // 365*24*60*60 = 31536000
	}
}