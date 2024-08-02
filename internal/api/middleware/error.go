package middleware

import (
	"exchange-backend/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			logger := log.FromContext(c).Sugar()
			for _, e := range c.Errors {
				logger.Errorf("Error: %s", e.Error())
			}
		}()

		c.Next()
	}
}
