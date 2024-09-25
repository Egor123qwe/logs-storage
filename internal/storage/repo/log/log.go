package log

import (
	"context"

	"github.com/Egor123qwe/logs-storage/internal/model/log"
)

type Log interface {
	Add(ctx context.Context, logs ...log.Log) error
}
