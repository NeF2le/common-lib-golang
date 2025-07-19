package middleware

import (
	"fmt"
	"github.com/NeF2le/common-lib-golang/logger"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"time"
)

type LogFormatterParams struct {
	RequestID   string
	HandlerName string
	StatusCode  int
	TimeStamp   time.Time
	Latency     time.Duration
}

func LoggingMiddleware(logger_ logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		params := &LogFormatterParams{
			RequestID:   c.Request.Header.Get("X-Request-Id"),
			StatusCode:  c.Writer.Status(),
			HandlerName: c.HandlerName(),
			Latency:     time.Since(start),
			TimeStamp:   time.Now(),
		}

		requestPath := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			requestPath += "?" + c.Request.URL.RawQuery
		}
		logMsg := fmt.Sprintf("%s %s", c.Request.Method, requestPath)
		logFields := structs.Map(params)

		if params.StatusCode >= 400 {
			logger_.Error(logMsg, logFields)
		} else {
			logger_.Info(logMsg, logFields)
		}
	}
}
