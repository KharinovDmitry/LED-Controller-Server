package panel

import (
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/entity"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"sync"
)

const (
	butchCount = 5
)

type Panel struct {
}

func New() *Panel {
	return &Panel{}
}

func (p *Panel) SendTaskButch(ctx context.Context, host *url.URL, tasks []entity.PanelTask) (entity.ButchReport, error) {
	errChan := make(chan error, butchCount)
	taskChan := make(chan workerTask, butchCount)
	defer close(taskChan)
	defer close(errChan)

	report := entity.ButchReport{}
	report.AllCount = len(tasks)

	for i := 0; i < butchCount; i++ {
		go p.taskWorker(ctx, taskChan, errChan)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for _, task := range tasks {
		taskChan <- workerTask{
			task: task,
			host: host,
		}
	}

	wg.Wait()

	for err := range errChan {
		if err != nil {
			report.ErrCount++
		} else {
			report.SucCount++
		}
	}

	return report, nil
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

type workerTask struct {
	task entity.PanelTask
	host *url.URL
}

func (p *Panel) taskWorker(ctx context.Context, butches chan workerTask, errs chan error) {
	for {
		select {
		case <-ctx.Done():
			slog.Debug("worker end")
			return
		case butch := <-butches:
			errs <- p.SendTask(ctx, butch.host, butch.task)
		}
	}
}
