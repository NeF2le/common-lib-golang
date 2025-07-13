package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LogFormatterParams struct {
	Request      *http.Request
	TimeStamp    time.Time
	Latency      time.Duration
	StatusCode   int
	Method       string
	Path         string
	ErrorMessage string
	Keys         map[string]any
	RequestID    string
}

func LoggingMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		params := &LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		params.TimeStamp = time.Now()
		params.Latency = params.TimeStamp.Sub(start)

		params.Method = params.Request.Method
		params.StatusCode = c.Writer.Status()
		params.ErrorMessage = c.Errors.String()

		params.Path = params.Request.URL.Path
		if params.Request.URL.RawQuery != "" {
			params.Path += "?" + params.Request.URL.RawQuery
		}

		logger.Infow(
			"method", params.Method,
			"path", params.Path,
			"statusCode", params.StatusCode,
			"latency", params.Latency,
			"errorMessage", params.ErrorMessage,
			"params", params.Keys,
		)
	}
}
