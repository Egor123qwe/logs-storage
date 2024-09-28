package resolver

import (
	"context"
	"fmt"

	srvmodel "github.com/Egor123qwe/logs-storage/internal/model"
	"github.com/Egor123qwe/logs-storage/internal/util"
	model "github.com/Egor123qwe/logs-storage/pkg/proto"
)

func (h Handler) GetLogs(ctx context.Context, req *model.LogFilter) (*model.LogResp, error) {
	srvreq := srvmodel.LogFilter{
		TraceID:  req.TraceID,
		ModuleID: req.ModuleID,

		Message: req.Message,

		CountOnPage: req.CountOnPage,
		Page:        req.Page,
	}

	if req.StartTime != nil {
		srvreq.StartTime = util.Ptr(req.StartTime.AsTime())
	}

	if req.EndTime != nil {
		srvreq.EndTime = util.Ptr(req.EndTime.AsTime())
	}

	if req.Level != nil {
		level := srvmodel.ConvertLevelName(*req.Level)
		if level == srvmodel.Invalid {
			return nil, fmt.Errorf("invalid level: %s", *req.Level)
		}

		srvreq.Level = &level
	}

	resp, err := h.srv.Logs().GetLogs(ctx, srvreq)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	result := &model.LogResp{
		PagesCount: resp.Total,
	}

	for _, log := range resp.Logs {
		result.Logs = append(result.Logs, &model.Log{
			Id:      log.ID,
			TraceID: log.TraceID,
			Module:  log.Module,

			Time:  log.Time.String(),
			Level: log.Level,

			Message: log.Message,
		})
	}

	return result, nil
}
