package storage

import "github.com/Egor123qwe/logs-storage/internal/storage/repo/log"

func (s storage) Log() log.Log {
	return s.psql.Log()
}
