package display

import (
	"DynamicLED/internal/domain/entity"
	"encoding/json"
)

type Display struct {
	Width  int       `json:"width"`
	Pixels ColorsRGB `json:"pixels"`
}

func (d *Display) ToEntity() entity.PanelDisplay {
	return entity.PanelDisplay{
		Width:  d.Width,
		Pixels: d.Pixels.ToEntities(),
	}
}

func (d *Display) FromEntity(e entity.PanelDisplay) {
	d.Width = e.Width
	d.Pixels = ColorsRGB{}
	d.Pixels.FromEntities(e.Pixels)
}

func (d *Display) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *Display) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, d)
}

type ColorRGB struct {
	R byte `json:"r"`
	G byte `json:"g"`
	B byte `json:"b"`
}

func (c *ColorRGB) ToEntity() entity.ColorRGB {
	return entity.ColorRGB{
		R: c.R,
		G: c.G,
		B: c.B,
	}
}

type ColorsRGB []ColorRGB

func (colors ColorsRGB) ToEntities() []entity.ColorRGB {
	res := make([]entity.ColorRGB, len(colors))
	for i, color := range colors {
		res[i] = color.ToEntity()
	}

	return res
}

func (colors *ColorsRGB) FromEntities(es []entity.ColorRGB) {
	res := make([]ColorRGB, len(es))
	for i, e := range es {
		res[i] = ColorRGB(e)
	}

	*colors = res
}
