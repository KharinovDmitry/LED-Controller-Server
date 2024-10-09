package user

import (
	"DynamicLED/internal/domain/constant"
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID     `db:"uuid"`
	Login    string        `db:"login"`
	Password string        `db:"password"`
	Role     constant.Role `db:"role"`
}
