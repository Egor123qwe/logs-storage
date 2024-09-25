package model

import (
	"errors"
)

var (
	ErrInvalidContent = errors.New("invalid content")

	ErrFailedToAddLogs = errors.New("failed to add logs")
)
