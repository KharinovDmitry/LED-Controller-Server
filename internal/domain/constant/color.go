package constant

import "DynamicLED/internal/domain/entity"

var (
	Red, _   = entity.NewColorRGBFromString("#FF0000")
	Green, _ = entity.NewColorRGBFromString("#00FF00")
	Blue, _  = entity.NewColorRGBFromString("#0000FF")
)
