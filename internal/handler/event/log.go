package event

import (
	"context"
	"fmt"

	"github.com/Egor123qwe/logs-storage/internal/handler/model/msg"
	"github.com/Egor123qwe/logs-storage/internal/model"
)

func (h handler) AddLogs(ctx context.Context, m []byte) error {
	reqMSG, err := msg.New(m).Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", model.ErrInvalidContent, err)
	}

	srvReq := make([]model.LogReq, len(reqMSG.Content))

	for i, log := range reqMSG.Content {
		level := model.ConvertLevelName(log.Level)
		if level == model.Invalid {
			return fmt.Errorf("%w: %s", model.ErrInvalidLogLevel, log.Level)
		}

		srvReq[i] = model.LogReq{
			TraceID:  log.TraceID,
			Time:     log.Time,
			ModuleID: log.ModuleID,
			Level:    level,
			Message:  log.Message,
		}
	}

	if err := h.srv.Logs().AddLogs(ctx, srvReq...); err != nil {
		return fmt.Errorf("%w: %w", model.ErrFailedToAddLogs, err)
	}

	return nil
}
