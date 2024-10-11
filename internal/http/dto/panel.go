package dto

import (
	"DynamicLED/internal/domain/entity"
	"github.com/google/uuid"
)

type Panel struct {
	UUID  uuid.UUID `json:"uuid"`
	Owner uuid.UUID `json:"owner"`
	Mac   string    `json:"mac"`
	Key   string    `json:"key"`
	Rev   int       `json:"rev"`
	Host  string    `json:"host"`
}

func PanelsToDTO(panels []entity.Panel) []Panel {
	res := make([]Panel, len(panels))
	for i, panel := range panels {
		res[i] = Panel(panel)
	}

	return res
}

type RegisterPanelRequest struct {
	Rev  int    `json:"rev"`
	Mac  string `json:"mac"`
	Host string `json:"host"`
}

type PanelTask struct {
	X     int      `json:"x"`
	Y     int      `json:"y"`
	Color ColorRGB `json:"color"`
}

func (p *PanelTask) ToEntity() entity.PanelTask {
	return entity.PanelTask{
		Position: entity.PanelPosition{
			X: p.X,
			Y: p.Y,
		},
		Color: entity.ColorRGB(p.Color),
	}
}

type Display struct {
	Pixels []ColorRGB `json:"pixels"`
	Width  int        `json:"width"`
}

func DisplayToDTO(display entity.PanelDisplay) Display {
	return Display{
		Pixels: PixelsToDTO(display.Pixels),
		Width:  display.Width,
	}
}

type ColorRGB struct {
	R byte `json:"r"`
	G byte `json:"g"`
	B byte `json:"b"`
}

func PixelsToDTO(pixels []entity.ColorRGB) []ColorRGB {
	res := make([]ColorRGB, len(pixels))
	for i, pixel := range pixels {
		res[i] = ColorRGB(pixel)
	}
	return res
}
