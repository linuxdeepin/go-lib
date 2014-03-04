package logger

import (
	"fmt"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

var logger *Logger

func init() {
	logger = NewLogger("logger_test")
	Suite(logger)
}

func (logger *Logger) TestFunc(c *C) {
	Println("test println")
	Printf("test printf: %s\n", "test")
	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warning("test warning: %v", fmt.Errorf("error message"))
	logger.Error("test error: ", fmt.Errorf("error message"))
	// logger.Panic("test panic")
	// logger.Fatal("test fatal")
}
