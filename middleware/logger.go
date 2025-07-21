package middleware

import (
	"bytes"
	"fmt"
	"github.com/NeF2le/common-lib-golang/logger"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseBodyWriter) Write(data []byte) (int, error) {
	r.body.Write(data)
	return r.ResponseWriter.Write(data)
}

type LogFormatterParams struct {
	RequestID   string
	HandlerName string
	StatusCode  int
	TimeStamp   time.Time
	Latency     time.Duration
	Body        string
	method      string
	path        string
	host        string
}

func logRequest(params *LogFormatterParams, logger_ logger.Logger) {
	logMsg := fmt.Sprintf("HTTP request %s %s on %s", params.method, params.path, params.host)
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

		bodyBuf := &bytes.Buffer{}
		writer := &responseBodyWriter{c.Writer, bodyBuf}
		c.Writer = writer

		c.Next()

		skipPaths := []string{
			"/metrics",
			"/health",
			"/healthz",
			"/healthcheck",
			"/ping",
		}
		skippedPath := false
		for _, path := range skipPaths {
			if strings.HasSuffix(c.Request.URL.Path, path) {
				skippedPath = true
				break
			}
		}
		if skippedPath {
			return
		}

		params := &LogFormatterParams{
			RequestID:   c.Request.Header.Get("X-Request-Id"),
			StatusCode:  c.Writer.Status(),
			HandlerName: c.HandlerName(),
			Latency:     time.Since(start),
			TimeStamp:   time.Now(),
			Body:        bodyBuf.String(),
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
