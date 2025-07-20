package middleware_test

import (
	"github.com/NeF2le/common-lib-golang/logger"
	"github.com/NeF2le/common-lib-golang/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGinLoggingMiddleware(t *testing.T) {
	logger_ := logger.NewZapLogger(true)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware.GinLoggingMiddleware(logger_))
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
}
