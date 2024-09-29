package model

import (
	"errors"
)

var (
	EmptyModuleName = errors.New("empty module name")

	ErrInvalidContent  = errors.New("invalid content")
	ErrInvalidLogLevel = errors.New("invalid log level")

	ErrFailedToAddLogs = errors.New("failed to add logs")
)
