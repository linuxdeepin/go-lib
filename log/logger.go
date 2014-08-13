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
	"pkg.linuxdeepin.com/lib/utils"
	"runtime"
	"strings"
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

// Backend defines interface of logger's back-ends.
type Backend interface {
	log(name string, level Priority, msg string) error
}

// Logger is a wrapper object to access Logger dbus service.
type Logger struct {
	name     string
	level    Priority
	backends []Backend
	config   *restartConfig
}

// NewLogger create a Logger object, which need a string as name to
// register Logger dbus service, if the environment variable exists
// which name stores in variable "DebugEnv", the default log level
// will be "LevelDebug" or is "LevelInfo".
func NewLogger(name string) (l *Logger) {
	// ignore panic
	defer func() {
		if err := recover(); err != nil {
			golog.Println("<info> dbus unavailable,", err)
		}
	}()
	golog.SetFlags(golog.Llongfile)

	l = &Logger{name: name}
	l.config = newRestartConfig(name)
	l.level = getDefaultLogLevel(name)
	l.AppendBackend(GetBackendConsole())
	l.AppendBackend(GetBackendDeepinlog(name))

	// notify new logger created
	for _, b := range l.backends {
		b.log(name, LevelInfo, "new logger: "+name)
	}

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
		if strings.Contains(strings.ToLower(name),
			strings.ToLower(os.Getenv(DebugMatchEnv))) {
			if !utils.IsEnvExists(DebugLevelEnv) {
				level = LevelDebug
			}
		} else {
			level = LevelDisable
		}
	}
	return
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

// AppendBackend append a log backend.
func (l *Logger) AppendBackend(b Backend) {
	l.backends = append(l.backends, b)
}

// ResetBackends clear all backends.
func (l *Logger) ResetBackends() {
	l.backends = nil
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

// TODO
func (l *Logger) BeginTracing() {
	l.Infof("%s begin", l.name)
}
func (l *Logger) EndTracing() {
	if err := recover(); err != nil {
		// TODO how to launch crash reporter
		l.Error(err)
		l.logEndFailed()
	} else {
		l.logEndSuccess()
	}
}
func (l *Logger) logEndSuccess() {
	l.Infof("%s end", l.name)
}
func (l *Logger) logEndFailed() {
	l.Infof("%s interruption", l.name)
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
		b.log(l.name, level, msg)
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
	if !loop {
		_, file, line, _ := runtime.Caller(calldepth)
		msg = fmt.Sprintf("%s:%d: %s", file, line, s)
	} else {
		_, file, line, ok := runtime.Caller(calldepth)
		msg = fmt.Sprintf("%s:%d: %s", file, line, s)
		for ok {
			calldepth++
			_, file, line, ok = runtime.Caller(calldepth)
			if ok {
				msg = fmt.Sprintf("%s\n  -> %s:%d", msg, file, line)
			}
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
	// TODO use lib/logquery to get log messages.
	if logapi == nil {
		return
	}
	var logId string
	for _, b := range l.backends {
		switch b.(type) {
		case *deepinlog:
			d, _ := b.(*deepinlog)
			logId = d.id
			break
		}
	}
	if len(logId) == 0 {
		return
	}
	// if deepin-crash-reporter exists, launch it
	if utils.IsFileExist(crashReporterExe) {
		// save config to a temporary json file
		l.config.LogDetail, _ = logapi.GetLog(logId)
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
