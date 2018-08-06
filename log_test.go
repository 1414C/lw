package lw

import (
	"fmt"
	// "io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

type TestSummariesTp struct {
	Console         time.Duration
	InfoEnabled     time.Duration
	InfoDisabled    time.Duration
	InfoWithLoc     time.Duration
	TraceEnabled    time.Duration
	TraceDisabled   time.Duration
	WarningEnabled  time.Duration
	WarningDisabled time.Duration
	WarningWithLoc  time.Duration
	PrintfStdOut    time.Duration
}

var TestSummaries TestSummariesTp

func TestMain(m *testing.M) {

	// run the tests
	code := m.Run()

	// test summaries
	fmt.Println()
	fmt.Println("=======================================================================")
	fmt.Printf("Console:\t\t\t%v\n", TestSummaries.Console)
	fmt.Printf("InfoEnabled:\t\t\t%v\n", TestSummaries.InfoEnabled)
	fmt.Printf("InfoWithLoc:\t\t\t%v\n", TestSummaries.InfoWithLoc)
	fmt.Printf("InfoDisabled:\t\t\t%v\n", TestSummaries.InfoDisabled)
	fmt.Printf("TraceEnabled:\t\t\t%v\n", TestSummaries.TraceEnabled)
	fmt.Printf("TraceDisabled:\t\t\t%v\n", TestSummaries.TraceDisabled)
	fmt.Printf("WarningEnabled:\t\t\t%v\n", TestSummaries.WarningEnabled)
	fmt.Printf("WarningWithLoc:\t\t\t%v\n", TestSummaries.WarningWithLoc)
	fmt.Printf("WarningDisabled:\t\t%v\n", TestSummaries.WarningDisabled)
	fmt.Printf("PrintfStdOut output:\t\t%v\n", TestSummaries.PrintfStdOut)
	fmt.Println("=======================================================================")
	fmt.Println()

	os.Exit(code)
}

func TestConsole(t *testing.T) {
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Console("This is a console test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active console message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.Console = total
}

func TestInfoEnabledColor(t *testing.T) {
	Enable(false, true, nil)
	InfoEnable(true)
	var d time.Duration
	var total time.Duration
	warmup := time.Now()
	log.Println(warmup)
	for i := 0; i < 10; i++ {
		start := time.Now()
		Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active info message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.InfoEnabled = total
	InfoEnable(false)
}

func TestInfoEnabled(t *testing.T) {
	Enable(false, false, nil)
	InfoEnable(true)
	var d time.Duration
	var total time.Duration
	warmup := time.Now()
	log.Println(warmup)
	for i := 0; i < 10; i++ {
		start := time.Now()
		Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active info message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.InfoEnabled = total
	InfoEnable(false)
}

func TestInfoEnabledNoTimintg(t *testing.T) {
	Enable(false, true, nil)
	InfoEnable(true)
	for i := 0; i < 10; i++ {
		Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
	}
	InfoEnable(false)
}

func TestInfoWithLoc(t *testing.T) {
	Enable(true, true, nil)
	InfoEnable(true)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active info message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.InfoWithLoc = total
	InfoEnable(false)
}

func TestInfoDisabled(t *testing.T) {
	Enable(false, true, nil)
	InfoEnable(false)
	var d time.Duration
	var total time.Duration
	warmup := time.Now()
	log.Println(warmup)
	for i := 0; i < 10; i++ {
		start := time.Now()
		Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Disabled info message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.InfoDisabled = total
	InfoEnable(false)
}

func TestTraceEnabled(t *testing.T) {
	Enable(false, false, nil)
	TraceEnable(true)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Trace("This is a TRACE test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active trace message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.TraceEnabled = total
	TraceEnable(false)
}

func TestTraceEnabledColor(t *testing.T) {
	Enable(false, true, nil)
	TraceEnable(true)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Trace("This is a TRACE test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active trace message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.TraceEnabled = total
	TraceEnable(false)
}

func TestTraceDisabled(t *testing.T) {
	Enable(false, true, nil)
	TraceEnable(false)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Trace("This is a TRACE test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Disabled trace message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.TraceDisabled = total
	TraceEnable(false)
}

func TestWarningEnabledColor(t *testing.T) {
	Enable(false, true, nil)
	WarningEnable(true)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active warning message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.WarningEnabled = total
	WarningEnable(false)
}

func TestWarningEnabled(t *testing.T) {
	Enable(false, false, nil)
	WarningEnable(true)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active warning message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.WarningEnabled = total
	WarningEnable(false)
}

func TestWarningDisabled(t *testing.T) {
	Enable(false, true, nil)
	WarningEnable(false)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("disabled warning message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.WarningDisabled = total
	WarningEnable(false)
}

func TestWarningWithLoc(t *testing.T) {
	Enable(true, true, nil)
	WarningEnable(true)
	var d time.Duration
	var total time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		Warning("This is a WARNING test with 2 vars. one: %v, two: %v", "var_1", 2)
		d = time.Since(start)
		log.Println("Active warning with location message with 2 vars took:", d)
		total += d
	}
	log.Println("Total log writing time:", total)
	TestSummaries.WarningWithLoc = total
	WarningEnable(false)
}

func TestDebugColor(t *testing.T) {
	Enable(true, true, nil)
	DebugEnable(true)
	for i := 0; i < 10; i++ {
		Debug("This is a DEBUG test with 2 vars. one: %v, two: %v", "var_1", 2)
	}
	DebugEnable(false)
}

func TestDebug(t *testing.T) {
	Enable(true, false, nil)
	DebugEnable(true)
	for i := 0; i < 10; i++ {
		Debug("This is a DEBUG test with 2 vars. one: %v, two: %v", "var_1", 2)
	}
	DebugEnable(false)
}

func TestErrorColor(t *testing.T) {
	Enable(true, true, nil)
	ErrorEnable(true)
	for i := 0; i < 10; i++ {
		Error(fmt.Errorf("This is an error test with 2 vars. one: %v, two: %v", "var_1", 2))
	}
	ErrorEnable(false)
}

func TestError(t *testing.T) {
	Enable(true, false, nil)
	ErrorEnable(true)
	for i := 0; i < 10; i++ {
		Error(fmt.Errorf("This is an error test with 2 vars. one: %v, two: %v", "var_1", 2))
	}
	ErrorEnable(false)
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

// func TestFatalColor(t *testing.T) {
// 	Enable(true, true, nil)
// 	for i := 0; i < 10; i++ {
// 		Fatal(fmt.Errorf("This is a fatal test with 2 vars. one: %v, two: %v", "var_1", 2))
// 	}
// }
