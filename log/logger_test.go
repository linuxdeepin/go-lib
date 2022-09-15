// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var originStdout = os.Stdout
var redirectStdoutFile = "testdata/stdout"
var redirectStdout, _ = os.OpenFile(redirectStdoutFile, os.O_CREATE|os.O_RDWR, 0644)

func redirectOutput() {
	os.Stdout = redirectStdout
}
func restoreOutput() {
	os.Stdout = originStdout
}
func resetOutput() {
	_ = redirectStdout.Truncate(0)
	_, _ = redirectStdout.Seek(0, io.SeekStart)
}
func readOutput() string {
	fileContent, err := ioutil.ReadFile(redirectStdoutFile)
	if err != nil {
		std.Println("read stdout file failed:", err)
	}
	return string(fileContent)
}
func checkOutput(t *testing.T, regfmt string, preferResult bool) {
	output := readOutput()
	result, _ := regexp.MatchString(regfmt, output)
	if result != preferResult {
		t.Errorf("match output failed: `%s`, `%#v`", regfmt, output)
	}
}

func TestGeneral(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			std.Println("catch panic:", err)
		}
	}()

	redirectOutput()
	defer restoreOutput()
	defer resetOutput()

	logger := NewLogger("logger_test")
	logger.SetLogLevel(LevelDebug)

	resetOutput()
	logger.Debug("test debug")
	checkOutput(t, `^<debug> logger_test.go:\d+: test debug\n$`, true)

	resetOutput()
	logger.Info("test info")
	checkOutput(t, `^<info> logger_test.go:\d+: test info\n$`, true)

	resetOutput()
	logger.Info("test info multi-lines\n\nthe thread line and following two empty lines\n\n")
	checkOutput(t, `^<info> logger_test.go:\d+: test info multi-lines\n\nthe thread line and following two empty lines\n\n\n$`, true)

	resetOutput()
	logger.Warning("test warning:", fmt.Errorf("error message"), "append string")
	checkOutput(t, `^<warning> logger_test.go:\d+: test warning: error message append string\n$`, true)

	resetOutput()
	logger.Warning("test warning:", fmt.Errorf("error message"))
	checkOutput(t, `^<warning> logger_test.go:\d+: test warning: error message\n$`, true)

	resetOutput()
	logger.Warningf("test warningf: %v", fmt.Errorf("error message"))
	checkOutput(t, `^<warning> logger_test.go:\d+: test warningf: error message\n$`, true)

	resetOutput()
	logger.Error("test error:", fmt.Errorf("error message"))
	checkOutput(t, `^<error> logger_test.go:\d+: test error: error message\n(  ->  \w+\.\w+:\d+\n)+$`, true)

	resetOutput()
	logger.Errorf("test errorf: %v", fmt.Errorf("error message"))
	checkOutput(t, `^<error> logger_test.go:\d+: test errorf: error message\n(  ->  \w+\.\w+:\d+\n)+$`, true)

	testPanicFunc := func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Info("got panic")
			}
		}()
		logger.Panic("test panic")
	}
	resetOutput()
	testPanicFunc()
	checkOutput(t, `^<error> logger_test.go:\d+: test panic\n(  ->  \w+\.\w+:\d+\n)+<info> logger_test.go:\d+: got panic\n$`, true)
}

// TODO: remove
func TestFuncTracing(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			std.Println("catch error:", err)
		}
	}()

	logger := NewLogger("logger_test")

	logger.BeginTracing()
	defer logger.EndTracing()
	defer func() {
		logger.EndTracing()
	}()
	logger.EndTracing()

	subFunc := func() {
		logger.BeginTracing()
		logger.EndTracing()
	}
	go subFunc()

	panic("test panic")
}

func TestIsNeedLog(t *testing.T) {
	logger := &Logger{}
	logger.SetLogLevel(LevelInfo)
	assert.Equal(t, logger.isNeedLog(LevelDebug), false)
	assert.Equal(t, logger.isNeedLog(LevelInfo), true)
	assert.Equal(t, logger.isNeedLog(LevelWarning), true)
	assert.Equal(t, logger.isNeedLog(LevelError), true)
	assert.Equal(t, logger.isNeedLog(LevelPanic), true)
	assert.Equal(t, logger.isNeedLog(LevelFatal), true)
	logger.SetLogLevel(LevelDebug)
	assert.Equal(t, logger.isNeedLog(LevelDebug), true)
	assert.Equal(t, logger.isNeedLog(LevelInfo), true)
	assert.Equal(t, logger.isNeedLog(LevelWarning), true)
	assert.Equal(t, logger.isNeedLog(LevelError), true)
	assert.Equal(t, logger.isNeedLog(LevelPanic), true)
	assert.Equal(t, logger.isNeedLog(LevelFatal), true)
}

func TestIsNeedTraceMore(t *testing.T) {
	logger := &Logger{}
	logger.SetLogLevel(LevelInfo)
	assert.Equal(t, logger.isNeedTraceMore(LevelDebug), false)
	assert.Equal(t, logger.isNeedTraceMore(LevelInfo), false)
	assert.Equal(t, logger.isNeedTraceMore(LevelWarning), false)
	assert.Equal(t, logger.isNeedTraceMore(LevelError), true)
	assert.Equal(t, logger.isNeedTraceMore(LevelPanic), true)
	assert.Equal(t, logger.isNeedTraceMore(LevelFatal), true)
}

func TestAddRemoveBackend(t *testing.T) {
	logger := &Logger{}

	var backendNull Backend
	logger.AddBackend(backendNull)
	assert.Equal(t, len(logger.backends), 0)
	var backendConsoleNull *backendConsole
	logger.AddBackend(backendConsoleNull)
	assert.Equal(t, len(logger.backends), 0)
	logger.ResetBackends()

	logger.AddBackendConsole()
	assert.Equal(t, len(logger.backends), 1)
	logger.AddBackendConsole()
	assert.Equal(t, len(logger.backends), 2)
	logger.RemoveBackendConsole()
	assert.Equal(t, len(logger.backends), 0)
}

func TestDebugFile(t *testing.T) {
	oldDebugFile := DebugFile
	DebugFile = "testdata/dde_debug"
	defer func() { DebugFile = oldDebugFile }()

	os.Clearenv()
	defer os.Clearenv()

	_ = os.Remove(DebugFile)
	assert.Equal(t, getDefaultLogLevel("test_debug_file"), LevelInfo)

	_, _ = os.Create(DebugFile)
	assert.Equal(t, getDefaultLogLevel("test_debug_file"), LevelDebug)
}

func TestDebugEnv(t *testing.T) {
	os.Clearenv()
	defer os.Clearenv()

	assert.Equal(t, getDefaultLogLevel("test_env"), LevelInfo)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG", "")
	assert.Equal(t, getDefaultLogLevel("test_env"), LevelDebug)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG", "1")
	assert.Equal(t, getDefaultLogLevel("test_env"), LevelDebug)
}

func TestDebugLevelEnv(t *testing.T) {
	os.Clearenv()
	defer os.Clearenv()

	assert.Equal(t, getDefaultLogLevel("test_env"), LevelInfo)

	_ = os.Setenv("DDE_DEBUG_LEVEL", "debug")
	assert.Equal(t, getDefaultLogLevel("test_env"), LevelDebug)

	_ = os.Setenv("DDE_DEBUG_LEVEL", "warning")
	assert.Equal(t, getDefaultLogLevel("test_env"), LevelWarning)
}

func TestDebugMatchEnv(t *testing.T) {
	os.Clearenv()
	defer os.Clearenv()

	_ = os.Setenv("DDE_DEBUG_MATCH", "test1")
	assert.Equal(t, getDefaultLogLevel("test1"), LevelDebug)
	assert.Equal(t, getDefaultLogLevel("test2"), LevelDisable)

	_ = os.Setenv("DDE_DEBUG_MATCH", "test1|test2")
	assert.Equal(t, getDefaultLogLevel("test1"), LevelDebug)
	assert.Equal(t, getDefaultLogLevel("test2"), LevelDebug)

	_ = os.Setenv("DDE_DEBUG_MATCH", "not match")
	assert.Equal(t, getDefaultLogLevel("test1"), LevelDisable)
	assert.Equal(t, getDefaultLogLevel("test2"), LevelDisable)
}

func TestDebugMixEnv(t *testing.T) {
	os.Clearenv()
	defer os.Clearenv()

	_ = os.Setenv("DDE_DEBUG", "1")
	_ = os.Setenv("DDE_DEBUG_LEVEL", "warning")
	assert.Equal(t, getDefaultLogLevel("test_env"), LevelWarning)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG_LEVEL", "error")
	_ = os.Setenv("DDE_DEBUG_MATCH", "test_env")
	assert.Equal(t, getDefaultLogLevel("test_env"), LevelError)

	os.Clearenv()
	_ = os.Setenv("DDE_DEBUG_LEVEL", "error")
	_ = os.Setenv("DDE_DEBUG_MATCH", "not match")
	assert.Equal(t, getDefaultLogLevel("test_env"), LevelDisable)
}

func TestDebugConsoleEnv(t *testing.T) {
	os.Clearenv()
	defer os.Clearenv()

	_ = os.Setenv("DDE_DEBUG_CONSOLE", "1")
	console := newBackendConsole("test-console")
	assert.Equal(t, console.syslogMode, true)

	redirectOutput()
	defer restoreOutput()
	resetOutput()
	_ = console.log(LevelInfo, "this line shows as syslog format in console")
	checkOutput(t, `\w+ \d+ \d{2}:\d{2}:\d{2} .* test-console\[\d+\]: <info> this line shows as syslog format in console\n$`, true)
}

func TestFmtSprint(t *testing.T) {
	assert.Equal(t, fmtSprint(""), "")
	assert.Equal(t, fmtSprint("a", "b", "c"), "a b c")
	assert.Equal(t, fmtSprint("a\n", "b\n", "c\n"), "a\n b\n c\n")
}
