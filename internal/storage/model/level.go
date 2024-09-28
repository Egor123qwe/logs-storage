package model

type Level int

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG

	INVALID = -1
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

func (l Level) String() string {
	return levelNames[l]
}
