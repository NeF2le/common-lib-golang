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
	method      string
	path        string
	host        string
}

func logRequest(params *LogFormatterParams, logger_ logger.Logger) {
	logMsg := fmt.Sprintf("%s %s", params.method, params.path)
	logFields := structs.Map(params)

	if params.StatusCode >= 400 {
		logger_.Error(logMsg, logFields)
	} else {
		logger_.Info(logMsg, logFields)
	}
}

func GinLoggingMiddleware(logger_ logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		params := &LogFormatterParams{
			RequestID:   c.Request.Header.Get("X-Request-Id"),
			StatusCode:  c.Writer.Status(),
			HandlerName: c.HandlerName(),
			Latency:     time.Since(start),
			TimeStamp:   time.Now(),
			method:      c.Request.Method,
			path:        c.FullPath(),
			host:        c.Request.Host,
		}

		if c.Request.URL.RawQuery != "" {
			params.path += "?" + c.Request.URL.RawQuery
		}

		logRequest(params, logger_)
	}
}
