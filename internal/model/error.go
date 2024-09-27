package model

import (
	"errors"
)

var (
	ErrInvalidContent  = errors.New("invalid content")
	ErrInvalidLogLevel = errors.New("invalid log level")

	ErrFailedToAddLogs = errors.New("failed to add logs")
)
