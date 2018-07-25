package lw

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// LogWriter is a logging struct implementing Logger
type LogWriter struct {
	mu             sync.Mutex
	writer         io.Writer
	enabled        bool
	locEnabled     bool
	traceEnabled   bool
	infoEnabled    bool
	warningEnabled bool
	debugEnabled   bool
	errorEnabled   bool
}

// LogWriterState is used to return the current status/state
// of lw's config.
type LogWriterState struct {
	Enabled        bool
	LocEnabled     bool
	TraceEnabled   bool
	InfoEnabled    bool
	WarningEnabled bool
	DebugEnabled   bool
	ErrorEnabled   bool
}

var logWriter LogWriter

// Enable enables lw at the package-level.  This does not have the
// effect of generating log entries unless at least one of the logging
// message types has been activated.  By default lw writes to os.Stdout,
// but if a valid io.Writer is provided, it will be used instead.
// Passing a nil value for io.Writer w will result in os.Stdout being
// used.
func Enable(withLoc bool, w io.Writer) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.enabled = true
	logWriter.locEnabled = withLoc
	if w != nil {
		logWriter.writer = w
		return
	}
	logWriter.writer = os.Stdout
}

// InitWithSettings configures lw as per the supplied parameters.
func InitWithSettings(s LogWriterState, w io.Writer) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.enabled = s.Enabled
	logWriter.locEnabled = s.LocEnabled
	logWriter.infoEnabled = s.InfoEnabled
	logWriter.warningEnabled = s.WarningEnabled
	logWriter.traceEnabled = s.TraceEnabled
	logWriter.debugEnabled = s.DebugEnabled
	logWriter.errorEnabled = s.ErrorEnabled
	if w != nil {
		logWriter.writer = w
		return
	}
	logWriter.writer = os.Stdout
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
	logWriter.writer = os.Stdout
	logWriter.infoEnabled = false
	logWriter.warningEnabled = false
	logWriter.traceEnabled = false
	logWriter.debugEnabled = false
	logWriter.errorEnabled = false
}

// SetWriter uses the supplied writer to set the output of the
// underlying log.  Be careful using this, as the Enable and
// Disable* functions will override this setting.
func SetWriter(w io.Writer) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	if w != nil {
		logWriter.writer = w
	}
	logWriter.writer = os.Stdout
}

// GetState returns the current state of the lw settings.  Note
// that this provides a snap-shot in time, as the settings may
// be changed in another goroutine immediately following the
// release of the mutex.
func GetState() LogWriterState {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	s := LogWriterState{
		Enabled:        logWriter.enabled,
		LocEnabled:     logWriter.locEnabled,
		TraceEnabled:   logWriter.traceEnabled,
		InfoEnabled:    logWriter.infoEnabled,
		WarningEnabled: logWriter.warningEnabled,
		DebugEnabled:   logWriter.debugEnabled,
		ErrorEnabled:   logWriter.errorEnabled,
	}
	return s
}

// InfoEnable enables the creation and output of Info messages.  Messages
// will be output based on the state of the logWriter.Enabled flag and the
// current value of the writer assigned to log.
func InfoEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.infoEnabled = a
}

// WarningEnable enables the creation and output of Warning messages.  Messages
// will be output based on the state of the logWriter.Enabled flag and the current
// value of the writer assigned to log.
func WarningEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.warningEnabled = a
}

// TraceEnable enables the creation and output of Trace messages.  Messages
// will be output based on the state of the logWriter.Enabled flag and the
// current value of the writer assigned to log.
func TraceEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.traceEnabled = a
}

// DebugEnable enables the creation and output of Debug messages.  Messages
// will be output based on the state of the logWriter.Enabled flag and the
// current value of the writer assigned to log.
func DebugEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.debugEnabled = a
}

// ErrorEnable enables the creation and output of Error messages.  Messages
// will be output based on the state of the logWriter.Enabled flag and the
// current value of the writer assigned to log.
func ErrorEnable(a bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.errorEnabled = a
}

// Console always writes to os.Stdout regardless of the lw.Enabled setting.
func Console(s string, i ...interface{}) {
	m := fmt.Sprintf(s, i...)
	io.WriteString(os.Stdout, m+"\n")
	return
}

// Info writes an Info message based on the current lw settings.  The method accepts a
// Printf-type formatted string and a list of operands to use in the verb-replacement.
// Note that you do not need to pass the newline escape code ("\n").
// Usage Example:
// lw.Info("This is a test %s with the number %d", "MESSAGE", 42)
func Info(s string, i ...interface{}) {
	if logWriter.infoEnabled {
		m := fmt.Sprintf(s, i...)
		if logWriter.locEnabled {
			_, f, line, ok := runtime.Caller(1)
			if ok {
				io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t INFO: "+f+" line:"+strconv.Itoa(line)+" "+m+"\n")
				return
			}
			io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t INFO: "+m+"\n")
			return
		}
		io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t INFO: "+m+"\n")
	}
}

// Trace writes a Trace message based on the current lw settings.  The method accepts a
// Printf-type formatted string and a list of operands to use in the verb-replacement.
// Note that you do not need to pass the newline escape code ("\n").
// Usage Example:
// lw.Trace("This is a test %s with the number %d", "MESSAGE", 42)
func Trace(s string, i ...interface{}) {
	if logWriter.traceEnabled {
		m := fmt.Sprintf(s, i...)
		_, f, line, ok := runtime.Caller(1)
		if ok {
			io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t TRACE: "+f+" line:"+strconv.Itoa(line)+" "+m+"\n")
			return
		}
		io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t TRACE: "+m+"\n")
	}
}

// Warning writes a Warning mesage based on the current lw settings.  The method accepts a
// Printf-type formatted string and a list of operands to use in the verb-replacement.
// Note that you do not need to pass the newline escape code ("\n").
// Usage Example:
// lw.Warning("This is a test %s with the number %d", "MESSAGE", 42)
func Warning(s string, i ...interface{}) {
	if logWriter.warningEnabled {
		m := fmt.Sprintf(s, i...)
		if logWriter.locEnabled {
			_, f, line, ok := runtime.Caller(1)
			if ok {
				io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t WARNING: "+f+" line:"+strconv.Itoa(line)+" "+m+"\n")
				return
			}
			io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t WARNING: "+m+"\n")
			return
		}
		io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t WARNING: "+m+"\n")
	}
}

// Debug writes a Debug message based on the current lw settings.  The method accepts a
// Printf-type formatted string and a list of operands to use in the verb-replacement.
// Note that you do not need to pass the newline escape code ("\n").
// Usage Example:
// lw.Debug("This is a test %s with the number %d", "MESSAGE", 42)
func Debug(s string, i ...interface{}) {
	if logWriter.debugEnabled {
		m := fmt.Sprintf(s, i...)
		_, f, line, ok := runtime.Caller(1)
		if ok {
			io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t DEBUG: "+f+" line:"+strconv.Itoa(line)+" "+m+"\n")
			return
		}
		io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t DEBUG: "+m+"\n")
	}
}

// Error writes an Error message based on the current lw settings.  The method accepts
// the standard golang error-type. Note that you do not need to pass the newline escape
// code ("\n").
// Usage Example:
// lw.Error("This is a test %s with the number %d", "MESSAGE", 42)
func Error(e error) {
	if logWriter.errorEnabled {
		_, f, line, ok := runtime.Caller(1)
		if ok {
			io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t ERROR: "+f+" line:"+strconv.Itoa(line)+" "+e.Error()+"\n")
			return
		}
		io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t ERROR: "+e.Error()+"\n")
	}
}

// Fatal writes a Fatal log-entry based on the current lw settings and the terminates
// the application via os.Exit(1).  The method accepts a printf-type formatted string
// and a list of operands to use in the verb-replacement.
// The Fatal message-type is always active irrespective of lw-settings.
// Note that you do not need to pass the newline escape code ("\n").
// Usage Example:
// lw.Fatal("This is a test %s with the number %d", "MESSAGE", 42)
func Fatal(e error) {
	_, f, line, ok := runtime.Caller(1)
	if ok {
		io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t FATAL: "+f+" line:"+strconv.Itoa(line)+" "+e.Error()+"\n")
		os.Exit(1)
	}
	io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t FATAL: "+e.Error()+"\n")
	os.Exit(1)
}
