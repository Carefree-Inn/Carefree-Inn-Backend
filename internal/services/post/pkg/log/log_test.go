package log

import (
	"testing"
	"time"
)

func TestFatal(t *testing.T) {
	NewLogger()
	Warn(nil, nil, "eee")
	Warn(nil, nil, "eee")
	Warn(nil, nil, "eee")
	time.Sleep(2 * time.Second)
}
