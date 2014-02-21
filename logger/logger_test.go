package logger

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

var logger *Logger

func init() {
	// logger = &Logger{}
	var err error
	logger, err = New("logger_test")
	if err == nil {
		// run test only new logger success
		Suite(logger)
	}
}

func (logger *Logger) TestFunc(c *C) {
	Println("test println")
	Printf("test printf: %s", "test")
	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warning("test warning")
	logger.Error("test error")
	// logger.Panic("test panic")
	// logger.Fatal("test fatal")
}
