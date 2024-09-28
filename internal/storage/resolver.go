package storage

import (
	"github.com/Egor123qwe/logs-storage/internal/storage/repo"
)

func (s storage) Log() repo.Log {
	return s.psql.Log()
}
