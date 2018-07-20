package lw

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
)

// LogWriter is a logging struct implementing Logger
type LogWriter struct {
	mu             sync.Mutex
	enabled        bool
	locEnabled     bool
	traceEnabled   bool
	infoEnabled    bool
	warningEnabled bool
	debugEnabled   bool
	errorEnabled   bool
	fatalEnabled   bool
}

// LogWriterState is used to return the current status/state
// of lw's config.
type LogWriterState struct {
	enabled        bool
	locEnabled     bool
	traceEnabled   bool
	infoEnabled    bool
	warningEnabled bool
	debugEnabled   bool
	errorEnabled   bool
	fatalEnabled   bool
}

var logWriter LogWriter

// Enable enables lw at the package-level.  This does not have the
// effect of generating log entries unless at least one of the logging
// message types has been activated.  By default lw writes to os.Stdout,
// but if a valid io.Writer is provided, it will be used instead.
// Passing a nil value for io.Writer w will result in os.Stdout being
// used.
func Enable(withLoc bool, w *io.Writer) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.enabled = true
	logWriter.locEnabled = withLoc
	if w != nil {
		log.SetOutput(*w)
		return
	}
	log.SetOutput(os.Stdout)
}

// InitWithSettings configures lw as per the supplied parameters.
func InitWithSettings(s LogWriterState, w *io.Writer) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.enabled = s.enabled
	logWriter.locEnabled = s.locEnabled
	logWriter.infoEnabled = s.infoEnabled
	logWriter.warningEnabled = s.warningEnabled
	logWriter.traceEnabled = s.traceEnabled
	logWriter.debugEnabled = s.debugEnabled
	logWriter.errorEnabled = s.errorEnabled
	logWriter.fatalEnabled = s.fatalEnabled
	if w != nil {
		log.SetOutput(*w)
		return
	}
	log.SetOutput(os.Stdout)
}

// Disable disables lw at the package-level, but leaves all current
// lw activation and output settings intact.
func Disable() {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.enabled = false
}

// DisableAndReset disables lw at the package-level and resets all lw
// activations to their initial state (no logging of any message-type).
// lw output will be reset to os.Stdout.
func DisableAndReset() {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.enabled = false
	logWriter.locEnabled = false
	log.SetOutput(os.Stdout)
	logWriter.infoEnabled = false
	logWriter.warningEnabled = false
	logWriter.traceEnabled = false
	logWriter.debugEnabled = false
	logWriter.errorEnabled = false
	logWriter.fatalEnabled = false
}

// SetWriter uses the supplied writer to set the output of the
// underlying log.  Be careful using this, as the Enable and
// Disable* functions will override this setting.
func SetWriter(w *io.Writer) {
	// don't bother with mu here, as log does it
	if w != nil {
		log.SetOutput(*w)
	}
	log.SetOutput(os.Stdout)
}

// GetState returns the current state of the lw settings.  Note
// that this provides a snap-shot in time, as the settings may
// be changed in another goroutine immediately following the
// release of mutex.
func GetState() LogWriterState {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	s := LogWriterState{
		enabled:        logWriter.enabled,
		locEnabled:     logWriter.locEnabled,
		traceEnabled:   logWriter.traceEnabled,
		infoEnabled:    logWriter.infoEnabled,
		warningEnabled: logWriter.warningEnabled,
		debugEnabled:   logWriter.debugEnabled,
		errorEnabled:   logWriter.errorEnabled,
		fatalEnabled:   logWriter.fatalEnabled,
	}
	return s
}

// InfoEnable enables the creation and output of Info messages.  Messages
// will be output based on the state of the logWriter.Enabled and the current
// value of the writer assigned to log.
func InfoEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.infoEnabled = a
}

// WarningEnable enables the creation and output of Warning messages.  Messages
// will be output based on the state of the logWriter.Enabled and the current
// value of the writer assigned to log.
func WarningEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.warningEnabled = a
}

// TraceEnable enables the creation and output of Trace messages.  Messages
// will be output based on the state of the logWriter.Enabled and the current
// value of the writer assigned to log.
func TraceEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.traceEnabled = a
}

// DebugEnable enables the creation and output of Debug messages.  Messages
// will be output based on the state of the logWriter.Enabled and the current
// value of the writer assigned to log.
func DebugEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.debugEnabled = a
}

// ErrorEnable enables the creation and output of Error messages.  Messages
// will be output based on the state of the logWriter.Enabled and the current
// value of the writer assigned to log.
func ErrorEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.errorEnabled = a
}

// FatalEnable enables the creation and output of Fatal messages.  Messages
// will be output based on the state of the logWriter.Enabled and the current
// value of the writer assigned to log.
func FatalEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.fatalEnabled = a
}

// Info writes an Info message based on the current lw settings.
func Info(s string, i ...interface{}) {
	if logWriter.infoEnabled {
		m := fmt.Sprintf(s, i...)
		if logWriter.locEnabled {
			_, f, line, ok := runtime.Caller(1)
			if ok {
				log.Println("INFO: " + f + " line:" + strconv.Itoa(line) + " " + m)
				return
			}
		}
		log.Println("INFO:", m)
	}
}

// Trace writes a Trace log-entry
func Trace(s string, i ...interface{}) {
	if logWriter.traceEnabled {
		m := fmt.Sprintf(s, i...)
		_, f, line, ok := runtime.Caller(1)
		if ok {
			log.Println("TRACE: " + f + " line:" + strconv.Itoa(line) + " " + m)
			return
		}
		log.Println("TRACE: ", m)
	}
}

// Warning writes a Warning log-entry
func Warning(s string, i ...interface{}) {
	if logWriter.warningEnabled {
		m := fmt.Sprintf(s, i...)
		if logWriter.locEnabled {
			_, f, line, ok := runtime.Caller(1)
			if ok {
				log.Println("WARNING: " + f + " line:" + strconv.Itoa(line) + " " + m)
				return
			}
		}
		log.Println("WARNING:", m)
	}
}

// Debug writes a Debug log-entry
func Debug(s string, i ...interface{}) {
	if logWriter.debugEnabled {
		m := fmt.Sprintf(s, i...)
		_, f, line, ok := runtime.Caller(1)
		if ok {
			log.Println("DEBUG: " + f + " line:" + strconv.Itoa(line) + " " + m)
			return
		}
		log.Println("DEBUG:", m)
	}
}

// Error writes an Error log-entry
func Error(s string, i ...interface{}) {
	if logWriter.errorEnabled {
		m := fmt.Sprintf(s, i...)
		_, f, line, ok := runtime.Caller(1)
		if ok {
			log.Println("ERROR: " + f + " line:" + strconv.Itoa(line) + " " + m)
			return
		}
		log.Println("ERROR:", m)
	}
}

// Fatal writes a Fatal log-entry
func Fatal(s string, i ...interface{}) {
	if logWriter.fatalEnabled {
		m := fmt.Sprintf(s, i...)
		_, f, line, ok := runtime.Caller(1)
		if ok {
			log.Println("FATAL: " + f + " line:" + strconv.Itoa(line) + " " + m)
			return
		}
		log.Println("FATAL:", m)
	}
}
