package resolver

import (
	"context"
	"fmt"

	srvmodel "github.com/Egor123qwe/logs-storage/internal/model"
	model "github.com/Egor123qwe/logs-storage/pkg/proto"
)

func (h Handler) GetModules(ctx context.Context, req *model.ModuleReq) (*model.ModuleResp, error) {
	srvreq := srvmodel.ModuleReq{
		NameFilter: req.NameFilter,
	}

	resp, err := h.srv.Logs().GetModules(ctx, srvreq)
	if err != nil {
		return nil, fmt.Errorf("failed to get modules: %w", err)
	}

	return &model.ModuleResp{Modules: resp}, nil
}

func (h Handler) InitModule(ctx context.Context, req *model.InitModuleReq) (*model.InitModuleResp, error) {
	resp, err := h.srv.Logs().InitModule(ctx, req.Module)
	if err != nil {
		return nil, fmt.Errorf("failed to init module: %w", err)
	}

	return &model.InitModuleResp{ModuleId: resp}, nil
}
