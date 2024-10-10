package constant

import "math"

var (
	PanelKeyLength     = 8
	PanelKeyLengthMask = int(math.Pow(float64(PanelKeyLength), 10))
)
