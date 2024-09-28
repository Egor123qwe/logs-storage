package model

import "time"

type Log struct {
	ID      int64
	TraceID string
	Module  string

	Time  time.Time
	Level string

	Message string
}

type LogReq struct {
	ID       int64
	TraceID  string
	ModuleID int64

	Time  time.Time
	Level Level

	Message string
}

type LogFilter struct {
	TraceID  *string
	ModuleID *int64

	Level   *Level
	Message string

	StartTime *time.Time
	EndTime   *time.Time

	CountOnPage int64
	Page        int64
}

type LogResp struct {
	Logs []Log

	Total int64
}
