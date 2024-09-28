package log

import (
	"context"

	"github.com/Egor123qwe/logs-storage/internal/model"
	dbModel "github.com/Egor123qwe/logs-storage/internal/storage/model"
)

func (s service) InitModule(ctx context.Context, module string) (int64, error) {
	return s.storage.Log().InitModule(ctx, module)
}

func (s service) GetModules(ctx context.Context, req model.ModuleReq) ([]string, error) {
	return s.storage.Log().GetModules(ctx, dbModel.ModuleReq{NameFilter: req.NameFilter})
}
