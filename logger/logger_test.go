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
	logger.SetLogLevel(LEVEL_DEBUG)
	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warning("test warning: ", fmt.Errorf("error message"), "append string")
	logger.Warning("test warning: %v ", fmt.Errorf("error message"))
	logger.Warningf("test warningf: %v", fmt.Errorf("error message"))
	logger.Error("test error: ", fmt.Errorf("error message"))
	logger.Errorf("test errorf: %v", fmt.Errorf("error message"))
	// logger.Panic("test panic")
	// logger.Fatal("test fatal")
}
