package event

import (
	"context"
	"fmt"

	"github.com/Egor123qwe/logs-storage/internal/handler/model/msg"
	"github.com/Egor123qwe/logs-storage/internal/model"
	logmodel "github.com/Egor123qwe/logs-storage/internal/model/log"
)

func (h handler) AddLogs(ctx context.Context, m []byte) error {
	reqMSG, err := msg.New(m).Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", model.ErrInvalidContent, err)
	}

	srvReq := make([]logmodel.Log, len(reqMSG.Content))

	for i, log := range reqMSG.Content {
		level := logmodel.ConvertLevelName(log.Level)

		if level == logmodel.Invalid {
			return fmt.Errorf("%w: %s", model.ErrInvalidLogLevel, log.Level)
		}

		srvReq[i] = logmodel.Log{
			TraceID: log.TraceID,
			Time:    log.Time,
			Module:  log.Module,
			Level:   level,
			Message: log.Message,
		}
	}

	if err := h.srv.Logs().Add(ctx, srvReq...); err != nil {
		return fmt.Errorf("%w: %w", model.ErrFailedToAddLogs, err)
	}

	return nil
}
