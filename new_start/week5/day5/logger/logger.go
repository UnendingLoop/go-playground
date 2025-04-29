package logger

import (
	"fmt"
	"strings"
)

// Logger - interface for controlling output
type Logger interface {
	Printf(s string, v ...any)
}

// Log - structure for methods and interface
type Log struct {
	Output []string
}

// Printf - output into string without stdout
func (log *Log) Printf(s string, v ...any) {
	output := fmt.Sprintf(s, v...)
	output = strings.TrimSpace(output)
	lines := strings.Split(output, "\n")
	log.Output = append(log.Output, lines[:]...)
}
