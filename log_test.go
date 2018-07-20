package lw

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	// "fmt"
	// "runtime"
	"testing"
)

type TestSummariesTp struct {
	InfoEnabled     time.Duration
	InfoDisabled    time.Duration
	InfoWithLoc     time.Duration
	TraceEnabled    time.Duration
	TraceDisabled   time.Duration
	WarningEnabled  time.Duration
	WarningDisabled time.Duration
	WarningWithLoc  time.Duration
	PrintfStdOut    time.Duration
	PrintfNOP       time.Duration
}

var TestSummaries TestSummariesTp

func TestMain(m *testing.M) {

	// run the tests
	code := m.Run()

	// test summaries
	fmt.Println()
	fmt.Println("=======================================================================")
	fmt.Printf("InfoEnabled:\t\t\t%v\n", TestSummaries.InfoEnabled)
	fmt.Printf("InfoWithLoc:\t\t\t%v\n", TestSummaries.InfoWithLoc)
	fmt.Printf("InfoDisabled:\t\t\t%v\n", TestSummaries.InfoDisabled)
	fmt.Printf("TraceEnabled:\t\t\t%v\n", TestSummaries.TraceEnabled)
	fmt.Printf("TraceDisabled:\t\t\t%v\n", TestSummaries.TraceDisabled)
	fmt.Printf("WarningEnabled:\t\t\t%v\n", TestSummaries.WarningEnabled)
	fmt.Printf("WarningWithLoc:\t\t\t%v\n", TestSummaries.WarningWithLoc)
	fmt.Printf("WarningDisabled:\t\t%v\n", TestSummaries.WarningDisabled)
	fmt.Printf("PrintfStdOut output:\t\t%v\n", TestSummaries.PrintfStdOut)
	fmt.Printf("PrintfNOP output:\t\t%v\n", TestSummaries.PrintfNOP)
	fmt.Println("=======================================================================")
	fmt.Println()

	os.Exit(code)
}

func TestInfoEnabled(t *testing.T) {
	l := LogWriter{}
	l.InfoEnabled = true
	var d time.Duration
	var total time.Duration
	warmup := time.Now()
	log.Println(warmup)
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active info message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.InfoEnabled = total
}

func TestInfoWithLoc(t *testing.T) {
	l := LogWriter{}
	l.InfoEnabled = true
	l.LocEnabled = true
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active info message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.InfoWithLoc = total
}

func TestInfoDisabled(t *testing.T) {
	l := LogWriter{}
	l.InfoEnabled = false
	var d time.Duration
	var total time.Duration
	warmup := time.Now()
	log.Println(warmup)
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Disabled info message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.InfoDisabled = total
}

func TestTraceEnabled(t *testing.T) {
	l := LogWriter{}
	l.TraceEnabled = true
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Trace("This is an TRACE test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active trace message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.TraceEnabled = total
}

func TestTraceDisabled(t *testing.T) {
	l := LogWriter{}
	l.TraceEnabled = false
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Trace("This is an TRACE test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Disabled trace message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.TraceDisabled = total
}

func TestWarningEnabled(t *testing.T) {
	l := LogWriter{}
	l.WarningEnabled = true
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active warning message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.WarningEnabled = total
}

func TestWarningDisabled(t *testing.T) {
	l := LogWriter{}
	l.WarningEnabled = false
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("disabled warning message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.WarningDisabled = total
}

func TestWarningWithLoc(t *testing.T) {
	l := LogWriter{}
	l.WarningEnabled = true
	l.LocEnabled = true
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		l.Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active warning with location message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.WarningWithLoc = total
}

func TestLogPrintfStdout(t *testing.T) {
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		log.Printf("This is a log.Printf stdout test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("log.Printf message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.PrintfStdOut = total
}

func TestLogPrintfNOP(t *testing.T) {
	var d time.Duration
	defer log.SetOutput(os.Stdout) // just in case
	var total time.Duration
	for i := 0; i < 10; i++ {
		log.SetOutput(ioutil.Discard)
		start := time.Now()
		log.Printf("This is a log.Printf stdout test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.SetOutput(os.Stdout)
		log.Println("log.Printf message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.PrintfNOP = total
}
