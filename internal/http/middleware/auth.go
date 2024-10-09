package middleware

import (
	"DynamicLED/internal/domain/service"
	"github.com/gin-gonic/gin"
)

func Auth(auth service.Auth) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
