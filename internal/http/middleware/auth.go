package middleware

import (
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/service"
	"DynamicLED/internal/http/dto"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
)

func Auth(auth service.Auth) func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.GetHeader(constant.AuthorizationHeader)
		parts := strings.Split(authorization, " ")
		if len(parts) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenType := parts[0]
		token := parts[1]

		if tokenType != constant.BearerTokenType {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := auth.ParseClaims(c.Request.Context(), token)
		if err != nil {
			if errors.Is(err, service.ErrTokenExpired) ||
				errors.Is(err, service.ErrInvalidCredentials) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewErrorResponse(err.Error()))
				return
			}

			slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Set(constant.ContextUserUUID, claims.UserUUID)
		c.Set(constant.ContextLogin, claims.Login)
		c.Set(constant.ContextRole, claims.Role)
	}
}
