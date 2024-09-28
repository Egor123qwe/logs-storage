package resolver

import (
	"context"

	module "github.com/Egor123qwe/logs-storage/internal/handler/grpc/generate"
	srvmodel "github.com/Egor123qwe/logs-storage/internal/model"
)

func (h Handler) GetAllowedLevels(context.Context, *module.LevelsReq) (*module.LevelsResp, error) {
	return &module.LevelsResp{Levels: srvmodel.LevelNames}, nil
}
