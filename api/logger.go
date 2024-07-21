package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ZerologMiddleware logs requests using zerolog
func ZerologMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Info().
			Str("client_ip", c.ClientIP()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", status).
			Str("latency", latency.String()).
			Str("user_agent", c.Request.UserAgent()).
			Str("error", c.Errors.ByType(gin.ErrorTypePrivate).String()).
			Msg("HTTP request")
	}
}
