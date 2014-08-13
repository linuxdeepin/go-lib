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

func (*tester) TestGeneral(c *C) {
	logger := NewLogger("logger_test")
	logger.BeginTracing()
	defer logger.EndTracing()
	logger.SetLogLevel(LevelDebug)
	logger.Debug("test debug")
	logger.Info("test info")
	logger.Info("test info multi-lines\n\nthe thread line and following two empty lines\n\n")
	logger.Warning("test warning: ", fmt.Errorf("error message"), "append string")
	logger.Warning("test warning: %v ", fmt.Errorf("error message"))
	logger.Warningf("test warningf: %v", fmt.Errorf("error message"))
	logger.Error("test error: ", fmt.Errorf("error message"))
	logger.Errorf("test errorf: %v", fmt.Errorf("error message"))
	logger.Panic("test panic")
}

func (*tester) TestIsNeedLog(c *C) {
	logger := &Logger{}
	logger.SetLogLevel(LevelInfo)
	c.Check(logger.isNeedLog(LevelDebug), Equals, false)
	c.Check(logger.isNeedLog(LevelInfo), Equals, true)
	c.Check(logger.isNeedLog(LevelWarning), Equals, true)
	c.Check(logger.isNeedLog(LevelError), Equals, true)
	c.Check(logger.isNeedLog(LevelPanic), Equals, true)
	c.Check(logger.isNeedLog(LevelFatal), Equals, true)
	logger.SetLogLevel(LevelDebug)
	c.Check(logger.isNeedLog(LevelDebug), Equals, true)
	c.Check(logger.isNeedLog(LevelInfo), Equals, true)
	c.Check(logger.isNeedLog(LevelWarning), Equals, true)
	c.Check(logger.isNeedLog(LevelError), Equals, true)
	c.Check(logger.isNeedLog(LevelPanic), Equals, true)
	c.Check(logger.isNeedLog(LevelFatal), Equals, true)
}

func (*tester) TestIsNeedTraceMore(c *C) {
	logger := &Logger{}
	logger.SetLogLevel(LevelInfo)
	c.Check(logger.isNeedTraceMore(LevelDebug), Equals, false)
	c.Check(logger.isNeedTraceMore(LevelInfo), Equals, false)
	c.Check(logger.isNeedTraceMore(LevelWarning), Equals, false)
	c.Check(logger.isNeedTraceMore(LevelError), Equals, true)
	c.Check(logger.isNeedTraceMore(LevelPanic), Equals, true)
	c.Check(logger.isNeedTraceMore(LevelFatal), Equals, true)
}

func (*tester) TestDebugFile(c *C) {
	os.Clearenv()
	os.Create(DebugFile)
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LevelDebug)

	os.Remove(DebugFile)
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LevelInfo)
}

func (*tester) TestDebugEnv(c *C) {
	os.Clearenv()
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LevelInfo)

	os.Setenv("DDE_DEBUG", "1")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LevelDebug)
}

func (*tester) TestDebugLevelEnv(c *C) {
	os.Clearenv()
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LevelInfo)

	os.Setenv("DDE_DEBUG_LEVEL", "debug")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LevelDebug)

	os.Setenv("DDE_DEBUG_LEVEL", "warning")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LevelWarning)
}

func (*tester) TestDebugMatchEnv(c *C) {
	os.Clearenv()
	os.Setenv("DDE_DEBUG_MATCH", "test1")
	logger1 := NewLogger("test1")
	logger2 := NewLogger("test2")
	c.Check(logger1.level, Equals, LevelDebug)
	c.Check(logger2.level, Equals, LevelDisable)

	os.Setenv("DDE_DEBUG_MATCH", "not match")
	logger1 = NewLogger("test1")
	logger2 = NewLogger("test2")
	c.Check(logger1.level, Equals, LevelDisable)
	c.Check(logger2.level, Equals, LevelDisable)
}

func (*tester) TestDebugMixEnv(c *C) {
	os.Clearenv()
	os.Setenv("DDE_DEBUG", "1")
	os.Setenv("DDE_DEBUG_LEVEL", "warning")
	logger := NewLogger("test_env")
	c.Check(logger.level, Equals, LevelWarning)

	os.Clearenv()
	os.Setenv("DDE_DEBUG_LEVEL", "error")
	os.Setenv("DDE_DEBUG_MATCH", "test_env")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LevelError)

	os.Clearenv()
	os.Setenv("DDE_DEBUG_LEVEL", "error")
	os.Setenv("DDE_DEBUG_MATCH", "not match")
	logger = NewLogger("test_env")
	c.Check(logger.level, Equals, LevelDisable)
}

func (*tester) TestFmtSprint(c *C) {
	c.Check(fmtSprint(""), Equals, "")
	c.Check(fmtSprint("a", "b", "c"), Equals, "a b c")
	c.Check(fmtSprint("a\n", "b\n", "c\n"), Equals, "a\n b\n c\n")
}
