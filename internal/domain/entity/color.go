package entity

import (
	"encoding/hex"
	"errors"
	"strings"
)

var (
	Red, _   = NewColorRGBFromString("#FF0000")
	Green, _ = NewColorRGBFromString("#00FF00")
	Blue, _  = NewColorRGBFromString("#0000FF")
)

type ColorRGB struct {
	R byte
	G byte
	B byte
}

func NewColorRGB(r, g, b byte) ColorRGB {
	return ColorRGB{
		R: r,
		G: g,
		B: b,
	}
}

func NewColorRGBFromString(color string) (ColorRGB, error) {
	colorRGB, _ := strings.CutPrefix(color, "#")
	bytes, err := hex.DecodeString(colorRGB)
	if err != nil {
		return ColorRGB{}, err
	}

	if len(bytes) < 3 {
		return ColorRGB{}, errors.New("invalid color")
	}

	return ColorRGB{
		B: bytes[2],
		G: bytes[1],
		R: bytes[0],
	}, nil
}
