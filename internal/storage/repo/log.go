package repo

import (
	"context"

	"github.com/Egor123qwe/logs-storage/internal/storage/model"
)

type Log interface {
	GetLogs(ctx context.Context, req model.LogFilter) (model.LogResp, error)
	AddLogs(ctx context.Context, logs ...model.LogReq) error

	GetModules(ctx context.Context, req model.ModuleReq) ([]string, error)
	InitModule(ctx context.Context, module string) (int64, error)
}
