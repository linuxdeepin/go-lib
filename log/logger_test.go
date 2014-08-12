package log

import (
	"fmt"
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type tester struct{}

var testDebugFile = "./dde_debug"

func init() {
	DebugFile = testDebugFile
	testWrapper := &tester{}
	Suite(testWrapper)
}

func (*tester) TestFunc(c *C) {
	logger := NewLogger("logger_test")
	logger.BeginTracing()
	defer logger.EndTracing()
	logger.SetLogLevel(LEVEL_DEBUG)
	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warning("test warning: ", fmt.Errorf("error message"), "append string")
	logger.Warning("test warning: %v ", fmt.Errorf("error message"))
	logger.Warningf("test warningf: %v", fmt.Errorf("error message"))
	logger.Error("test error: ", fmt.Errorf("error message"))
	logger.Errorf("test errorf: %v", fmt.Errorf("error message"))
	logger.Panic("test panic")
}

func (*tester) TestDebugFile(c *C) {
	os.Clearenv()
	os.Create(DebugFile)
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_DEBUG)

	os.Remove(DebugFile)
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_INFO)
}

func (*tester) TestDebugEnv(c *C) {
	os.Clearenv()
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_INFO)

	os.Setenv("DDE_DEBUG", "1")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_DEBUG)
}

func (*tester) TestDebugLevelEnv(c *C) {
	os.Clearenv()
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_INFO)

	os.Setenv("DDE_DEBUG_LEVEL", "debug")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_DEBUG)

	os.Setenv("DDE_DEBUG_LEVEL", "warning")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_WARNING)
}

func (*tester) TestDebugMatchEnv(c *C) {
	os.Clearenv()
	os.Setenv("DDE_DEBUG_MATCH", "test1")
	logger1 := NewLogger("test1")
	logger2 := NewLogger("test2")
	c.Check(logger1.level, Equals, LEVEL_DEBUG)
	c.Check(logger2.level, Equals, LEVEL_DISABLE)

	os.Setenv("DDE_DEBUG_MATCH", "not match")
	logger1 = NewLogger("test1")
	logger2 = NewLogger("test2")
	c.Check(logger1.level, Equals, LEVEL_DISABLE)
	c.Check(logger2.level, Equals, LEVEL_DISABLE)
}

func (*tester) TestDebugMixEnv(c *C) {
	os.Clearenv()
	os.Setenv("DDE_DEBUG", "1")
	os.Setenv("DDE_DEBUG_LEVEL", "warning")
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_WARNING)

	os.Clearenv()
	os.Setenv("DDE_DEBUG_LEVEL", "error")
	os.Setenv("DDE_DEBUG_MATCH", "test_env")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_ERROR)

	os.Clearenv()
	os.Setenv("DDE_DEBUG_LEVEL", "error")
	os.Setenv("DDE_DEBUG_MATCH", "not match")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LEVEL_DISABLE)
}
