package panel

import "github.com/google/uuid"

type Panel struct {
	UUID  uuid.UUID `db:"uuid"`
	Owner uuid.UUID `db:"owner"`
	Mac   string    `db:"mac"`
	Rev   int       `db:"rev"`
	Host  string    `db:"host"`
}
