package lw

import (
	"fmt"
	"log"
)

// Logger is a logging interface
type Logger interface {
	Enable()
	Disable()
	Info(string, ...interface{})
	Trace(string, ...interface{})
	Warning(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
	Log(string, ...interface{}) (int, error)
}

// LogWriter is a logging struct implementing Logger
type LogWriter struct {
	Enabled        bool
	InfoEnabled    bool
	TraceEnabled   bool
	WarningEnabled bool
	DebugEnabled   bool
	ErrorEnabled   bool
	FatalEnabled   bool
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
		m := fmt.Sprintf(s, i...)
		log.Println(m)
	}
	return 0, nil
}

// Info writes and Info log-entry
func (l *LogWriter) Info(s string, i ...interface{}) {
	if l.InfoEnabled {
		m := fmt.Sprintf(s, i...)
		// _, f, line, ok := runtime.Caller(1)
		// if ok {
		// 	log.Println(f + " " + strconv.Itoa(line) + " " + m)
		// 	return
		// }
		log.Println(m)
	}
}

// Trace writes a Trace log-entry
func (l *LogWriter) Trace(s string, i ...interface{}) {
	if l.TraceEnabled {
		m := fmt.Sprintf(s, i...)
		log.Println(m)
	}
}

// Warning writes a Warning log-entry
func (l *LogWriter) Warning(s string, i ...interface{}) {
	if l.WarningEnabled {
		m := fmt.Sprintf(s, i...)
		log.Println(m)
	}
}

// Debug writes a Debug log-entry
func (l *LogWriter) Debug(s string, i ...interface{}) {
	if l.DebugEnabled {
		m := fmt.Sprintf(s, i...)
		log.Println(m)
	}
}

// Error writes an Error log-entry
func (l *LogWriter) Error(s string, i ...interface{}) {
	if l.ErrorEnabled {
		m := fmt.Sprintf(s, i...)
		log.Println(m)
	}
}

// Fatal writes a Fatal log-entry
func (l *LogWriter) Fatal(s string, i ...interface{}) {
	if l.FatalEnabled {
		m := fmt.Sprintf(s, i...)
		log.Println(m)
	}
}
