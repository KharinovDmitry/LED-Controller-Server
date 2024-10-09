package entity

import (
	"DynamicLED/internal/domain/constant"
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID
	Login    string
	Password string
	Role     constant.Role
}
