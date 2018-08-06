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
	colorEnabled   bool
	infoTxt        string
	warnTxt        string
	errorTxt       string
	traceTxt       string
	debugTxt       string
	fatalTxt       string
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
	ColorEnabled   bool
}

var logWriter LogWriter

func withColor(b bool) {
	if b {
		logWriter.infoTxt = "\x1b[32;1mINFO:\t\x1b[0m"
		logWriter.warnTxt = "\x1b[38;5;11mWARNING:  \x1b[0m"
		logWriter.errorTxt = "\x1b[38;5;9mERROR:\t\x1b[0m"
		logWriter.traceTxt = "\x1b[38;5;13mTRACE:\t\x1b[0m"
		logWriter.debugTxt = "\x1b[38;5;213mDEBUG:\t\x1b[0m"
		logWriter.fatalTxt = "\x1b[38;5;9mFATAL:\t\x1b[0m"
		return
	}
	logWriter.infoTxt = "INFO:\t"
	logWriter.warnTxt = "WARNING:  "
	logWriter.errorTxt = "ERROR:\t"
	logWriter.traceTxt = "TRACE:\t"
	logWriter.debugTxt = "DEBUG:\t"
	logWriter.fatalTxt = "FATAL:\t"
}

// Enable enables lw at the package-level.  This does not have the
// effect of generating log entries unless at least one of the logging
// message types has been activated.  By default lw writes to os.Stdout,
// but if a valid io.Writer is provided, it will be used instead.
// Passing a nil value for io.Writer w will result in os.Stdout being
// used.
func Enable(withLoc bool, withCol bool, w io.Writer) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.enabled = true
	logWriter.locEnabled = withLoc
	logWriter.colorEnabled = withCol
	withColor(logWriter.colorEnabled)
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
	logWriter.colorEnabled = s.ColorEnabled
	withColor(logWriter.colorEnabled)
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
	logWriter.colorEnabled = false
	withColor(logWriter.colorEnabled)
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
		ColorEnabled:   logWriter.colorEnabled,
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

// ColorEnable sets/unsets the coloring of the message type.
func ColorEnable(c bool) {
	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()
	logWriter.colorEnabled = c
	withColor(c)
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
				io.WriteString(logWriter.writer, logWriter.infoTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\t"+f+" line:"+strconv.Itoa(line)+"\n")
				return
			}
			io.WriteString(logWriter.writer, logWriter.infoTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\n")
			return
		}
		io.WriteString(logWriter.writer, logWriter.infoTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\n")
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
			io.WriteString(logWriter.writer, logWriter.traceTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\t"+f+" line:"+strconv.Itoa(line)+"\n")
			return
		}
		io.WriteString(logWriter.writer, logWriter.traceTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\n")
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
				io.WriteString(logWriter.writer, logWriter.warnTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\t"+f+" line:"+strconv.Itoa(line)+"\n")
				return
			}
			io.WriteString(logWriter.writer, logWriter.warnTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\n")
			return
		}
		io.WriteString(logWriter.writer, logWriter.warnTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\n")
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
			// io.WriteString(logWriter.writer, logWriter.debugTxt+time.Now().Format(time.RFC3339Nano)+"\t"+f+" line:"+strconv.Itoa(line)+"\t"+m+"\n")
			io.WriteString(logWriter.writer, logWriter.debugTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\t"+f+" line:"+strconv.Itoa(line)+"\n")
			return
		}
		io.WriteString(logWriter.writer, logWriter.debugTxt+time.Now().Format(time.RFC3339Nano)+"\t"+m+"\n")
	}
}

// Error writes an Error message based on the current lw settings.  The method accepts
// the standard golang error-type. Note that you do not need to pass the newline escape
// code ("\n").
// Usage Example:
// lw.Error(e)
func Error(e error) {
	if logWriter.errorEnabled {
		_, f, line, ok := runtime.Caller(1)
		if ok {
			// io.WriteString(logWriter.writer, logWriter.errorTxt+time.Now().Format(time.RFC3339Nano)+"\t"+f+" line:"+strconv.Itoa(line)+"\t"+e.Error()+"\n")
			io.WriteString(logWriter.writer, logWriter.errorTxt+time.Now().Format(time.RFC3339Nano)+"\t"+e.Error()+"\t"+f+" line:"+strconv.Itoa(line)+"\n")
			return
		}
		io.WriteString(logWriter.writer, logWriter.errorTxt+time.Now().Format(time.RFC3339Nano)+"\t"+e.Error()+"\n")
	}
}

// ErrorWithPrefixString writes an Error message based on the current lw settings.  The
// method accepts the standard golang error-type and a prefix string for the error message.
// Note that you do not need to pass the newline escape code ("\n").
// Usage Example:
// e error
// lw.ErrorWithPrefixString("Auth Controller Create() got:", e)
func ErrorWithPrefixString(s string, e error) {
	if logWriter.errorEnabled {
		_, f, line, ok := runtime.Caller(1)
		if ok {
			// io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t ERROR: "+s+" "+f+" line:"+strconv.Itoa(line)+" "+e.Error()+"\n")
			io.WriteString(logWriter.writer, logWriter.errorTxt+time.Now().Format(time.RFC3339Nano)+"\t"+s+" "+f+" line:"+strconv.Itoa(line)+"\t"+e.Error()+"\n")
			return
		}
		// io.WriteString(logWriter.writer, time.Now().Format(time.RFC3339Nano)+"\t ERROR: "+s+" "+e.Error()+"\n")
		io.WriteString(logWriter.writer, logWriter.errorTxt+time.Now().Format(time.RFC3339Nano)+"\t"+s+" "+e.Error()+"\n")
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
		io.WriteString(logWriter.writer, logWriter.fatalTxt+time.Now().Format(time.RFC3339Nano)+"\t"+e.Error()+"\t"+f+" line:"+strconv.Itoa(line)+"\n")
		os.Exit(1)
	}
	io.WriteString(logWriter.writer, logWriter.fatalTxt+time.Now().Format(time.RFC3339Nano)+"\t"+e.Error()+"\n")
	os.Exit(1)
}
