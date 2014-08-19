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
var logger = NewLogger("logger_test")

func init() {
	DebugFile = testDebugFile
	testWrapper := &tester{}
	Suite(testWrapper)
}

func (*tester) BenchmarkSyslog(c *C) {
	b := newBackendSyslog("benchSyslog")
	for i := 0; i < c.N; i++ {
		b.log(LevelInfo, "test")
	}
}

func (*tester) BenchmarkDeepinlog(c *C) {
	b := newBackendDeepinlog("benchDeepinlog")
	for i := 0; i < c.N; i++ {
		b.log(LevelInfo, "test")
	}
}

func (*tester) TestGeneral(c *C) {
	defer func() {
		if err := recover(); err != nil {
			logger.Info("catch error:", err)
		}
	}()
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

func (*tester) TestFuncTracing(c *C) {
	defer func() {
		if err := recover(); err != nil {
			logger.Info("catch error:", err)
		}
	}()
	logger.BeginTracing()
	defer logger.EndTracing()
	defer func() {
		logger.EndTracing()
	}()
	logger.EndTracing()
	go doTestFuncTracing()
	panic("test panic")
}
func doTestFuncTracing() {
	logger.BeginTracing()
	logger.EndTracing()
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

func (*tester) TestAddRemoveBackend(c *C) {
	logger := &Logger{}

	var backendNull Backend
	logger.AddBackend(backendNull)
	c.Check(len(logger.backends), Equals, 0)
	var backendConsoleNull *backendConsole
	logger.AddBackend(backendConsoleNull)
	c.Check(len(logger.backends), Equals, 0)
	logger.ResetBackends()

	logger.AddBackendConsole()
	c.Check(len(logger.backends), Equals, 1)
	logger.AddBackendConsole()
	c.Check(len(logger.backends), Equals, 2)
	logger.RemoveBackendConsole()
	c.Check(len(logger.backends), Equals, 0)
}

func (*tester) TestDebugFile(c *C) {
	os.Clearenv()
	defer os.Clearenv()

	os.Create(DebugFile)
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelDebug)

	os.Remove(DebugFile)
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelInfo)
}

func (*tester) TestDebugEnv(c *C) {
	os.Clearenv()
	defer os.Clearenv()

	c.Check(getDefaultLogLevel("test_env"), Equals, LevelInfo)

	os.Clearenv()
	os.Setenv("DDE_DEBUG", "")
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelDebug)

	os.Clearenv()
	os.Setenv("DDE_DEBUG", "1")
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelDebug)
}

func (*tester) TestDebugLevelEnv(c *C) {
	os.Clearenv()
	defer os.Clearenv()

	c.Check(getDefaultLogLevel("test_env"), Equals, LevelInfo)

	os.Setenv("DDE_DEBUG_LEVEL", "debug")
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelDebug)

	os.Setenv("DDE_DEBUG_LEVEL", "warning")
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelWarning)
}

func (*tester) TestDebugMatchEnv(c *C) {
	os.Clearenv()
	defer os.Clearenv()

	os.Setenv("DDE_DEBUG_MATCH", "test1")
	c.Check(getDefaultLogLevel("test1"), Equals, LevelDebug)
	c.Check(getDefaultLogLevel("test2"), Equals, LevelDisable)

	os.Setenv("DDE_DEBUG_MATCH", "not match")
	c.Check(getDefaultLogLevel("test1"), Equals, LevelDisable)
	c.Check(getDefaultLogLevel("test2"), Equals, LevelDisable)
}

func (*tester) TestDebugMixEnv(c *C) {
	os.Clearenv()
	defer os.Clearenv()

	os.Setenv("DDE_DEBUG", "1")
	os.Setenv("DDE_DEBUG_LEVEL", "warning")
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelWarning)

	os.Clearenv()
	os.Setenv("DDE_DEBUG_LEVEL", "error")
	os.Setenv("DDE_DEBUG_MATCH", "test_env")
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelError)

	os.Clearenv()
	os.Setenv("DDE_DEBUG_LEVEL", "error")
	os.Setenv("DDE_DEBUG_MATCH", "not match")
	c.Check(getDefaultLogLevel("test_env"), Equals, LevelDisable)
}

func (*tester) TestDebugConsoleEnv(c *C) {
	os.Clearenv()
	defer os.Clearenv()

	os.Setenv("DDE_DEBUG_CONSOLE", "1")
	console := newBackendConsole("test-console")
	c.Check(console.syslogMode, Equals, true)
	console.log(LevelInfo, "this line shows as syslog format in console")
}

func (*tester) TestFmtSprint(c *C) {
	c.Check(fmtSprint(""), Equals, "")
	c.Check(fmtSprint("a", "b", "c"), Equals, "a b c")
	c.Check(fmtSprint("a\n", "b\n", "c\n"), Equals, "a\n b\n c\n")
}
