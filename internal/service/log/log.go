package log

import (
	"context"

	"github.com/Egor123qwe/logs-storage/internal/model"
	"github.com/Egor123qwe/logs-storage/internal/storage"
	dbModel "github.com/Egor123qwe/logs-storage/internal/storage/model"
	"github.com/Egor123qwe/logs-storage/internal/util"
)

type Service interface {
	GetLogs(ctx context.Context, req model.LogFilter) (model.LogResp, error)
	AddLogs(ctx context.Context, logs ...model.LogReq) error

	GetModules(ctx context.Context, req model.ModuleReq) ([]string, error)
	InitModule(ctx context.Context, module string) (int64, error)
}

type service struct {
	storage storage.Storage
}

func New(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s service) AddLogs(ctx context.Context, logs ...model.LogReq) error {
	req := make([]dbModel.LogReq, len(logs))

	for i, log := range logs {
		req[i] = dbModel.LogReq{
			TraceID:  log.TraceID,
			ModuleID: log.ModuleID,
			Time:     log.Time,
			Level:    dbModel.Level(log.Level),
			Message:  log.Message,
		}
	}

	return s.storage.Log().AddLogs(ctx, req...)
}

func (s service) GetLogs(ctx context.Context, req model.LogFilter) (model.LogResp, error) {
	dbReq := dbModel.LogFilter{
		TraceID:  req.TraceID,
		ModuleID: req.ModuleID,

		Message: req.Message,

		StartTime: req.StartTime,
		EndTime:   req.EndTime,

		CountOnPage: req.CountOnPage,
		Page:        req.Page,
	}

	if req.Level != nil {
		dbReq.Level = util.Ptr(dbModel.Level(*req.Level))
	}

	dbResp, err := s.storage.Log().GetLogs(ctx, dbReq)
	if err != nil {
		return model.LogResp{}, err
	}

	result := model.LogResp{
		Total: dbResp.Total,
	}

	for _, l := range dbResp.Logs {
		log := model.Log{
			ID:      l.ID,
			TraceID: l.TraceID,
			Module:  l.Module,

			Time:  l.Time,
			Level: l.Level,

			Message: l.Message,
		}

		result.Logs = append(result.Logs, log)
	}

	return result, nil
}
