package entity

import (
	"github.com/google/uuid"
)

type Panel struct {
	UUID  uuid.UUID
	Owner uuid.UUID
	Mac   string
	Rev   int
	Host  string
}

type PanelDisplay struct {
	Pixels []ColorRGB
	Width  int
}

func (d *PanelDisplay) GetColorByPos(position PanelPosition) ColorRGB {
	return d.Pixels[position.X*d.Width+position.Y]
}

type PanelPosition struct {
	X int
	Y int
}

func NewPanelPosition(x, y int) *PanelPosition {
	return &PanelPosition{
		X: x,
		Y: y,
	}
}

type PanelTask struct {
	Position PanelPosition
	Color    ColorRGB
}

type ButchReport struct {
	AllCount int
	ErrCount int
	SucCount int
}
