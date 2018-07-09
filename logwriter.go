package lw

import (
	"log"
)

// Logger is a logging interface
type Logger interface {
	Enable()
	Disable()
	Log(string, ...interface{}) (int, error)
}

// LogWriter is a logging struct implementing Logger
type LogWriter struct {
	Enabled bool
	Logger
}

// Enable the LogWriter globally
func (l *LogWriter) Enable() {
	l.Enabled = true
}

// Disable the LogWriter globally
func (l *LogWriter) Disable() {
	l.Enabled = false
}

// Log accepts a Printf-style set of arguments
func (l *LogWriter) Log(s string, i ...interface{}) (int, error) {
	if l.Enabled {
		log.Printf("%s,%v\n", s, i)
	}
	return 0, nil
}
