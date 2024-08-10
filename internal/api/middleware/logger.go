package middleware

import (
	"time"

	"exchange-backend/internal/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.FromContext(c)
		if traceId := c.GetHeader("X-Amzn-Trace-Id"); traceId != "" {
			logger = logger.With(zap.String("x-amzn-trace-id", traceId))
		}

		requestId := c.GetHeader("X-Request-Id")
		if requestId == "" {
			requestId = ksuid.New().String()
		}
		logger = logger.With(zap.String("x-request-id", requestId))

		log.WithContext(c, logger)

		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path
		// Process request
		c.Next()

		// Stop timer
		param.Latency = time.Since(start)
		param.BodySize = c.Writer.Size()
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		log.FromContext(c).Sugar().
			With("latency", param.Latency).
			With("client_ip", param.ClientIP).
			With("method", param.Method).
			With("status", param.StatusCode).
			With("error", param.ErrorMessage).
			With("body_size", param.BodySize).
			With("path", param.Path).
			Infof("access-out")
	}
}
