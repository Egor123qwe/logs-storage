package model

import "strings"

type Level int

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG

	Invalid = -1
)

var LevelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

func (l Level) String() string {
	return LevelNames[l]
}

func ConvertLevelName(name string) Level {
	for i, levelName := range LevelNames {
		if strings.EqualFold(levelName, name) {
			return Level(i)
		}
	}

	return Invalid
}
