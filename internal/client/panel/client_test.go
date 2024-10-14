package panel

import (
	"DynamicLED/internal/domain/entity"
	"context"
	"net/url"
	"testing"
	"time"
)

func TestHappyPath(t *testing.T) {
	client := New()
	path, _ := url.Parse("192.168.0.102")

	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			posX := x
			posY := y
			color := entity.ColorRGB{
				R: uint8(posX * 16),
				G: uint8(posY * 16),
				B: 0,
			}
			if y == 0 && x == 0 {
				color = entity.ColorRGB{
					R: 8,
					G: 8,
					B: 0,
				}
			}

			err := client.SendTask(context.Background(), path, entity.PanelTask{
				Position: entity.PanelPosition{
					X: posX,
					Y: posY,
				},
				Color: color,
			})
			if err != nil {
				t.Error(err.Error())
			}
			time.Sleep(time.Millisecond * 100)

		}
	}

}
