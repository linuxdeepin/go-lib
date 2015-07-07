/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 **/

package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	golog "log"
	"os"
	"path/filepath"
	"pkg.deepin.io/lib/utils"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

const (
	defaultDebugEnv      = "DDE_DEBUG"
	defaultDebugLelveEnv = "DDE_DEBUG_LEVEL"
	defaultDebugMatchEnv = "DDE_DEBUG_MATCH"
	defaultDebugFile     = "/var/cache/dde_debug"
	crashReporterExe     = "/usr/bin/deepin-crash-reporter" // TODO
)

// Priority is the data type of log level.
type Priority int

// Definitions of log level, the larger of the value, the higher of
// the priority.
const (
	LevelDisable Priority = iota
	LevelFatal
	LevelPanic
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

var (
	// DebugEnv is the name of environment variable that used to
	// enable debug mode , if exists the default log level will be
	// "LevelDebug".
	DebugEnv = defaultDebugEnv

	// DebugLevelEnv is the name of environment variable that used to
	// control the log level, could be "debug", "info", "warning",
	// "error", "fatal" and "disable".
	DebugLevelEnv = defaultDebugLelveEnv

	// DebugMatchEnv is the name of environment variable that used to
	// enable debug mode for target logger object.
	DebugMatchEnv = defaultDebugMatchEnv

	// DebugFile if the file name that if exist the default log level
	// will be "LevelDebug".
	DebugFile = defaultDebugFile
)

var (
	errUnknownLogLevel = fmt.Errorf("unknown log level")
	std                = golog.New(os.Stderr, "", golog.Lshortfile)
)

// Backend defines interface of logger's back-ends.
type Backend interface {
	log(level Priority, msg string) error
	close() error
}

// Logger is a wrapper object to access Logger dbus service.
type Logger struct {
	name         string
	level        Priority
	backends     []Backend
	backendsLock sync.Mutex
	config       *restartConfig
}

// NewLogger create a Logger object, which need a string as name to
// register Logger dbus service, if the environment variable exists
// which name stores in variable "DebugEnv", the default log level
// will be "LevelDebug" or is "LevelInfo".
func NewLogger(name string) (l *Logger) {
	// ignore panic
	defer func() {
		if err := recover(); err != nil {
			std.Println("<info> dbus unavailable,", err)
		}
	}()

	l = &Logger{name: name}
	l.config = newRestartConfig(name)
	l.level = getDefaultLogLevel(name)
	l.AddBackendConsole()
	l.AddBackendSyslog()
	return
}

// parse environment variables to get default log level
func getDefaultLogLevel(name string) (level Priority) {
	level = LevelInfo
	if utils.IsEnvExists(DebugLevelEnv) {
		switch os.Getenv(DebugLevelEnv) {
		case "debug":
			level = LevelDebug
		case "info":
			level = LevelInfo
		case "warning":
			level = LevelWarning
		case "error":
			level = LevelError
		case "fatal":
			level = LevelFatal
		case "disable":
			level = LevelDisable
		}
	}
	if utils.IsEnvExists(DebugEnv) || utils.IsFileExist(DebugFile) {
		if !utils.IsEnvExists(DebugLevelEnv) {
			level = LevelDebug
		}
	}
	if utils.IsEnvExists(DebugMatchEnv) {
		if matchLogger(name, os.Getenv(DebugMatchEnv)) {
			if !utils.IsEnvExists(DebugLevelEnv) {
				level = LevelDebug
			}
		} else {
			level = LevelDisable
		}
	}
	return
}
func matchLogger(name, expr string) bool {
	reg, err := regexp.Compile(expr)
	if err != nil {
		std.Printf("match variable $%s failed %v\n", DebugMatchEnv, err)
		return false
	}
	return reg.Match([]byte(name))
}

// SetLogLevel reset the log level.
func (l *Logger) SetLogLevel(level Priority) *Logger {
	l.level = level
	return l
}

// GetLogLevel return the log level.
func (l *Logger) GetLogLevel() Priority {
	return l.level
}

// ResetBackends clear all backends.
func (l *Logger) ResetBackends() {
	for _, b := range l.backends {
		b.close()
	}
	l.backends = nil
}

// AddBackend append a log back-end.
func (l *Logger) AddBackend(b Backend) bool {
	l.backendsLock.Lock()
	defer l.backendsLock.Unlock()
	if utils.IsInterfaceNil(b) {
		return false
	}
	l.backends = append(l.backends, b)
	return true
}

// RemoveBackend remove all back-end with target type.
func (l *Logger) RemoveBackend(b Backend) {
	len := len(l.backends)
	targetType := reflect.TypeOf(b)
	for i := len - 1; i >= 0; i-- {
		itemType := reflect.TypeOf(l.backends[i])
		if itemType == targetType {
			l.doRemoveBackend(i)
		}
	}
}
func (l *Logger) doRemoveBackend(i int) {
	l.backendsLock.Lock()
	defer l.backendsLock.Unlock()
	l.backends[i].close()
	l.backends[i] = nil
	newLen := len(l.backends) - 1
	copy(l.backends[i:], l.backends[i+1:])
	l.backends[newLen] = nil
	l.backends = l.backends[:newLen]
}

// AddBackendConsole append a console back-end.
func (l *Logger) AddBackendConsole() bool {
	return l.AddBackend(newBackendConsole(l.name))
}

// RemoveBackendConsole remove all console back-end.
func (l *Logger) RemoveBackendConsole() {
	l.RemoveBackend(&backendConsole{})
}

// AddBackendSyslog append a syslog back-end.
func (l *Logger) AddBackendSyslog() bool {
	return l.AddBackend(newBackendSyslog(l.name))
}

// RemoveBackendSyslog remove all console back-end.
func (l *Logger) RemoveBackendSyslog() {
	l.RemoveBackend(&backendSyslog{})
}

// SetRestartCommand reset the command and argument when restart after fatal.
func (l *Logger) SetRestartCommand(exefile string, args ...string) {
	l.config.RestartCommand = append([]string{exefile}, args...)
}

// AddExtArgForRestart add the command option which be used when
// process fataled and restart by Logger dbus service.
func (l *Logger) AddExtArgForRestart(arg string) {
	if !isStringInArray(arg, l.config.RestartCommand[1:]) {
		l.config.RestartCommand = append(l.config.RestartCommand, arg)
	}
}

// BeginTracing log function information when entering it.
func (l *Logger) BeginTracing() {
	// TODO
	// funcName, file, line, ok := getCallerFuncInfo(2)
	// if !ok {
	// 	return
	// }
	// msg := fmt.Sprintf("%s:%d %s begin", filepath.Base(file), line, funcName)
	// l.doLog(LevelInfo, msg)
}

// EndTracing log function information when leaving it.
func (l *Logger) EndTracing() {
	// TODO
	// funcName, file, line, ok := getCallerFuncInfo(2)
	// if !ok {
	// 	return
	// }
	// msg := fmt.Sprintf("%s:%d %s end", filepath.Base(file), line, funcName)
	// l.doLog(LevelInfo, msg)
}

func getCallerFuncInfo(skip int) (funcName string, file string, line int, ok bool) {
	_, file, line, ok = runtime.Caller(skip)
	if !ok {
		return
	}
	for {
		pc, _, _, ok2 := runtime.Caller(skip)
		if !ok2 {
			break
		}
		if funcInfo := runtime.FuncForPC(pc); funcInfo != nil {
			// get the short function name
			fullFuncName := funcInfo.Name()
			funcName = filepath.Base(fullFuncName)
			if funcName == "runtime.panic" {
				// if is panic function, skip it and get file/line again
				_, file, line, ok = runtime.Caller(skip + 1)
				skip++
				continue
			} else if strings.Contains(funcName, "func·") {
				// if is an anonymous function, just skip it
				skip++
				continue
			} else {
				break
			}
		} else {
			break
		}
	}
	// fix function name
	a := strings.Split(funcName, ".")
	// strings.Split() will return at least one element, so we could
	// get the last element safely
	funcName = a[len(a)-1]
	funcName = funcName + "()"
	return
}

func (l *Logger) isNeedLog(level Priority) bool {
	if level <= l.level {
		return true
	}
	return false
}

func (l *Logger) isNeedTraceMore(level Priority) bool {
	if level <= LevelError {
		return true
	}
	return false
}

func (l *Logger) log(level Priority, v ...interface{}) {
	if !l.isNeedLog(level) {
		return
	}
	s := buildMsg(3, l.isNeedTraceMore(level), v...)
	l.doLog(level, s)
}
func (l *Logger) logf(level Priority, format string, v ...interface{}) {
	if !l.isNeedLog(level) {
		return
	}
	s := buildFormatMsg(3, l.isNeedTraceMore(level), format, v...)
	l.doLog(level, s)
}
func (l *Logger) doLog(level Priority, msg string) {
	for _, b := range l.backends {
		b.log(level, msg)
	}
}

func buildMsg(calldepth int, loop bool, v ...interface{}) (msg string) {
	s := fmtSprint(v...)
	msg = doBuildMsg(calldepth+1, loop, s)
	return
}
func buildFormatMsg(calldepth int, loop bool, format string, v ...interface{}) (msg string) {
	s := fmt.Sprintf(format, v...)
	msg = doBuildMsg(calldepth+1, loop, s)
	return
}
func doBuildMsg(calldepth int, loop bool, s string) (msg string) {
	var file, lastFile string
	var line, lastLine int
	var ok bool
	_, file, line, ok = runtime.Caller(calldepth)
	lastFile, lastLine = file, line
	msg = fmt.Sprintf("%s:%d: %s", filepath.Base(file), line, s)
	if loop && ok {
		for {
			calldepth++
			_, file, line, ok = runtime.Caller(calldepth)
			if file == lastFile && line == lastLine {
				// prevent infinite loop for that some platforms not
				// works well, e.g. mips
				break
			}
			if ok {
				msg = fmt.Sprintf("%s\n  ->  %s:%d", msg, filepath.Base(file), line)
			}
			lastFile, lastLine = file, line
		}
	}
	return
}

// Debug log a message in "debug" level.
func (l *Logger) Debug(v ...interface{}) {
	l.log(LevelDebug, v...)
}

// Debugf formats message according to a format specifier and log it in "debug" level.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logf(LevelDebug, format, v...)
}

// Info log a message in "info" level.
func (l *Logger) Info(v ...interface{}) {
	l.log(LevelInfo, v...)
}

// Infof formats message according to a format specifier and log it in "info" level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.logf(LevelInfo, format, v...)
}

// Warning log a message in "warning" level.
func (l *Logger) Warning(v ...interface{}) {
	l.log(LevelWarning, v...)
}

// Warningf formats message according to a format specifier and log it in "warning" level.
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.logf(LevelWarning, format, v...)
}

// Error log a message in "error" level.
func (l *Logger) Error(v ...interface{}) {
	l.log(LevelError, v...)
}

// Errorf formats message according to a format specifier and log it in "error" level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logf(LevelError, format, v...)
}

// Panic is equivalent to Error() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	l.log(LevelPanic, v...)
	s := fmtSprint(v...)
	panic(s)
}

// Panicf is equivalent to Errorf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logf(LevelPanic, format, v...)
	s := fmt.Sprintf(format, v...)
	panic(s)
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1).
func (l *Logger) Fatal(v ...interface{}) {
	l.log(LevelFatal, v...)
	l.launchCrashReporter()
	os.Exit(1)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logf(LevelFatal, format, v...)
	l.launchCrashReporter()
	os.Exit(1)
}

func (l *Logger) launchCrashReporter() {
	// if deepin-crash-reporter exists, launch it
	if utils.IsFileExist(crashReporterExe) {
		// save config to a temporary json file
		l.config.LogDetail = "not ready" // TODO use lib/logquery to get log messages.
		fileContent, err := json.Marshal(l.config)
		if err != nil {
			l.Error(err)
		}

		// create temporary json file and it will be removed by deepin-crash-reporter
		f, err := ioutil.TempFile("", "deepin_crash_reporter_config_")
		defer f.Close()
		if err != nil {
			l.Error(err)
		}
		_, err = f.Write(fileContent)
		if err != nil {
			l.Error(err)
		}

		// launch crash reporter
		l.Info("launch deepin-crash-reporter: %s %s", crashReporterExe, append(crashReporterArgs, f.Name()))
		_, err = os.StartProcess(crashReporterExe, append(crashReporterArgs, f.Name()),
			&os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		if err != nil {
			l.Error("launch deepin-crash-reporter failed:", err)
		}
	}
}
