package log

import (
	"testing"
	"time"
)

func TestFatal(t *testing.T) {
	NewLogger()
	//Warn(nil, nil, "eee")
	//Trace(log.WithField("TEST", "ttt"))
	//Info(nil, "hello")
	Info(logger.WithField("TTT", "ttt"))
	time.Sleep(2 * time.Second)
}
