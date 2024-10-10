package entity

import (
	"DynamicLED/internal/domain/constant"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	UUID     uuid.UUID     `json:"uuid"`
	UserUUID uuid.UUID     `json:"userUUID"`
	Login    string        `json:"login"`
	Role     constant.Role `json:"role"`
	Expire   time.Time     `json:"expire"`
}
