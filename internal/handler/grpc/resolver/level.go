package resolver

import (
	"context"

	srvmodel "github.com/Egor123qwe/logs-storage/internal/model"
	model "github.com/Egor123qwe/logs-storage/pkg/proto"
)

func (h Handler) GetAllowedLevels(context.Context, *model.LevelsReq) (*model.LevelsResp, error) {
	return &model.LevelsResp{Levels: srvmodel.LevelNames}, nil
}
