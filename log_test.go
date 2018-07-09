package lw

import (
	"fmt"
	"runtime"
	"testing"
)

func TestInfo(t *testing.T) {
	l := LogWriter{}
	l.Enable()
	l.Log("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
	pc, f, line, ok := runtime.Caller(0)
	if ok {
		fmt.Println("pc:", pc)
		fmt.Println("f:", f)
		fmt.Println("line:", line)
	}
}
