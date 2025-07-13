package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.Request.Header.Get("X-Request-Id")
		if rid == "" {
			rid = uuid.NewString()
		}
		c.Set("X-Request-Id", rid)
		c.Request.Header.Set("X-Request-Id", rid)
		c.Next()
	}
}
