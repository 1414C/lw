package lw

import (
	"io/ioutil"
	"log"
	"os"
	"time"
	// "fmt"
	// "runtime"
	"testing"
)

func TestInfo(t *testing.T) {
	l := LogWriter{}
	l.InfoEnabled = true
	var d time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active info message with 2 vars took:", d)
	}
}

func TestTrace(t *testing.T) {
	l := LogWriter{}
	l.TraceEnabled = true
	var d time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Trace("This is an TRACE test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active trace message with 2 vars took:", d)
	}
}

func TestWarningEnabled(t *testing.T) {
	l := LogWriter{}
	l.WarningEnabled = true
	var d time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active warning message with 2 vars took:", d)
	}
}

func TestLogPrintfStdout(t *testing.T) {
	var d time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		log.Printf("This is a log.Printf stdout test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("log.Printf message with 2 vars took:", d)
	}
}

func TestLogPrintfNOP(t *testing.T) {
	var d time.Duration
	defer log.SetOutput(os.Stdout) // just in case

	for i := 0; i < 10; i++ {
		log.SetOutput(ioutil.Discard)
		start := time.Now()
		log.Printf("This is a log.Printf stdout test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.SetOutput(os.Stdout)
		log.Println("log.Printf message with 2 vars took:", d)
	}
}

func TestWarningDisabled(t *testing.T) {
	l := LogWriter{}
	l.WarningEnabled = false
	var d time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Disabled warning message with 2 vars took:", d)
	}
}
