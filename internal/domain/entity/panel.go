package entity

import "github.com/google/uuid"

type Panel struct {
	UUID  uuid.UUID
	Owner uuid.UUID
	Mac   string
	Rev   int
	Host  string
}

type PanelPosition struct {
	X int
	Y int
}

type PanelTask struct {
	Position PanelPosition
	Color    string
}
