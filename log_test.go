package lw

import (
	// "fmt"
	// "runtime"
	"testing"
)

func TestInfo(t *testing.T) {
	l := LogWriter{}
	l.InfoEnabled = true
	l.Info("This is an INFO test with 2 vars. one: %v, two: %v", "var_1", 2)
}
