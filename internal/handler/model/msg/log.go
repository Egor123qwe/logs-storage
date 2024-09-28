package msg

import (
	"time"
)

type LogReq struct {
	ID       int64  `json:"id,omitempty"`
	TraceID  string `json:"trace_id"`
	ModuleID int64  `json:"module_id"`

	Time  time.Time `json:"time"`
	Level string    `json:"level"`

	Message string `json:"message"`
}
