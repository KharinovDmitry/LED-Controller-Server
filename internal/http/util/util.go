package util

import (
	"DynamicLED/internal/domain/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetClaimsFromContext(c *gin.Context) (login string, role string, userUUID uuid.UUID, err error) {
	loginCtx, ok := c.Get(constant.ContextLogin)
	if !ok {
		return "", "", uuid.Nil, fmt.Errorf("login is required")
	}

	login, ok = loginCtx.(string)
	if !ok {
		return "", "", uuid.Nil, fmt.Errorf("login is invalid")
	}

	roleCtx, ok := c.Get(constant.ContextRole)
	if !ok {
		return "", "", uuid.Nil, fmt.Errorf("role is required")
	}

	role, ok = roleCtx.(string)
	if !ok {
		return "", "", uuid.Nil, fmt.Errorf("role is invalid")
	}

	userUUIDCtx, ok := c.Get(constant.ContextUserUUID)
	if !ok {
		return "", "", uuid.Nil, fmt.Errorf("user uuid is required")
	}

	userUUID, ok = userUUIDCtx.(uuid.UUID)
	if !ok {
		return "", "", uuid.Nil, fmt.Errorf("role is invalid")
	}

	return login, role, userUUID, nil
}
