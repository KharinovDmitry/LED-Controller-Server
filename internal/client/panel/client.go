package panel

import (
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/entity"
	"context"
	"fmt"
	"net"
	"net/url"
)

type Panel struct {
}

func New() *Panel {
	return &Panel{}
}

func (p *Panel) SendTask(ctx context.Context, host *url.URL, task entity.PanelTask) error {
	req := []byte{
		byte(task.Position.X >> 8), byte(task.Position.X & 0xFF),
		byte(task.Position.Y >> 8), byte(task.Position.Y & 0xFF),
		task.Color.R, task.Color.G, task.Color.B,
	}

	conn, err := net.Dial(constant.ProtocolUDP, host.String()+":8888")
	if err != nil {
		return fmt.Errorf("[ Panel Client ] SendTask")
	}
	defer conn.Close()

	if _, err = conn.Write(req); err != nil {
		return fmt.Errorf("[ Panel Client ] SendTask")
	}

	return nil
}

func (p *Panel) GetDisplay(ctx context.Context, host *url.URL) (entity.PanelDisplay, error) {
	return entity.PanelDisplay{}, nil
}
