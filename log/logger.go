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

// TODO
var crashReporterArgs = []string{crashReporterExe, "--remove-config", "--config"}

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
	logapi *Logapi

	// DebugEnv is the name of environment variable to enable debug
	// mode , if exists the default log level will be "LevelDebug".
	DebugEnv = defaultDebugEnv

	// DebugLevelEnv is the name of environment variable to control
	// the log level, could be "debug", "info", "warning", "error" and
	// "fatal".
	DebugLevelEnv = defaultDebugLelveEnv

	// DebugMatchEnv is the name of environment variable to enable
	// debug mode for target logger object.
	DebugMatchEnv = defaultDebugMatchEnv

	// DebugFile if the file name that if exist the default log level
	// will be "LevelDebug".
	DebugFile = defaultDebugFile
)

func initLogapi() (err error) {
	if logapi == nil {
		logapi, err = newLogapi("/com/deepin/api/Logger")
	}
	return
}

// restartConfig stores data to be used by deepin-crash-reporter
type restartConfig struct {
	AppName          string
	RestartCommand   []string
	RestartEnv       map[string]string
	RestartDirectory string
	LogDetail        string
}

func newRestartConfig(logname string) *restartConfig {
	config := &restartConfig{}
	config.AppName = logname
	config.RestartCommand = os.Args
	config.RestartCommand[0], _ = filepath.Abs(os.Args[0])
	config.RestartDirectory, _ = os.Getwd()

	// setup envrionment variables
	config.RestartEnv = make(map[string]string)
	environs := os.Environ()
	for _, env := range environs {
		values := strings.SplitN(env, "=", 2)
		// values[0] is environment variable name, values[1] is the value
		if len(values) == 2 {
			config.RestartEnv[values[0]] = values[1]
		}
	}
	return config
}

func buildMsg(calldepth int, loop bool, v ...interface{}) (msg string) {
	s := fmt.Sprintln(v...)
	s = strings.TrimSuffix(s, "\n")
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
				msg = fmt.Sprintf("%s\n-> %s:%d", msg, file, line)
			}
		}
	}
	return
}

// Logger is a wrapper object to access Logger dbus service.
type Logger struct {
	name   string
	id     string
	level  Priority
	config *restartConfig
}

// NewLogger create a Logger object, which need a string as name to
// register Logger dbus service, if the environment variable exists
// which name stores in variable "DebugEnv", the default log level
// will be "LevelDebug" or is "LevelInfo".
func NewLogger(name string) (l *Logger) {
	golog.SetFlags(golog.Llongfile)

	// ignore panic
	defer func() {
		if err := recover(); err != nil {
			golog.Println("[INFO] dbus unavailable,", err)
		}
	}()

	l = &Logger{name: name}
	l.config = newRestartConfig(name)

	// dispatch environment variables to set default log level
	level := LevelInfo
	var customLevel Priority
	// level := defaultLevel
	if utils.IsEnvExists(DebugLevelEnv) {
		switch os.Getenv(DebugLevelEnv) {
		case "debug":
			customLevel = LevelDebug
		case "info":
			customLevel = LevelInfo
		case "warning":
			customLevel = LevelWarning
		case "error":
			customLevel = LevelError
		case "fatal":
			customLevel = LevelFatal
		}
		level = customLevel
	}
	if utils.IsEnvExists(DebugEnv) || isFileExists(DebugFile) {
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
	l.level = level

	err := initLogapi()
	if err != nil {
		golog.Printf("init logger dbus api failed: %v\n", err)
		return
	}
	l.id, err = logapi.NewLogger(name)
	if err != nil {
		golog.Printf("create logger api object failed: %v\n", err)
		return
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

// SetRestartCommand reset the command and argument when restart after fatal.
func (l *Logger) SetRestartCommand(exefile string, args ...string) {
	l.config.RestartCommand = append([]string{exefile}, args...)
}

// AddExtArgForRestart add the command option which be used when
// process fataled and restart by Logger dbus service.
func (l *Logger) AddExtArgForRestart(arg string) {
	if !stringInSlice(arg, l.config.RestartCommand[1:]) {
		l.config.RestartCommand = append(l.config.RestartCommand, arg)
	}
}

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

func (l *Logger) doLog(level Priority, s string) {
	switch level {
	case LevelDebug:
		if logapi != nil {
			logapi.Debug(l.id, l.name, s)
		}
		// TODO
		l.printLocal("[DEBUG]", s)
	case LevelInfo:
		if logapi != nil {
			logapi.Info(l.id, l.name, s)
		}
		l.printLocal("[INFO]", s)
	case LevelWarning:
		if logapi != nil {
			logapi.Warning(l.id, l.name, s)
		}
		l.printLocal("[WARNING]", s)
	case LevelError:
		if logapi != nil {
			logapi.Error(l.id, l.name, s)
		}
		l.printLocal("[ERROR]", s)
	case LevelPanic:
		if logapi != nil {
			logapi.Error(l.id, l.name, s)
		}
		l.printLocal("[PANIC]", s)
	case LevelFatal:
		if logapi != nil {
			logapi.Fatal(l.id, l.name, s)
		}
		l.printLocal("[FATAL]", s)
	}
}

func (l *Logger) printLocal(prefix, msg string) {
	fmtmsg := prefix + " " + msg
	fmtmsg = strings.Replace(fmtmsg, "\n", "\n"+prefix+" ", -1) // format multi-lines message
	fmt.Println(fmtmsg)
}

// Debug send a log message with 'DEBUG' as prefix to Logger dbus service and print it to console, too.
func (l *Logger) Debug(v ...interface{}) {
	l.log(LevelDebug, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logf(LevelDebug, format, v...)
}

// Info send a log message with 'INFO' as prefix to Logger dbus service and print it to console, too.
func (l *Logger) Info(v ...interface{}) {
	l.log(LevelInfo, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logf(LevelInfo, format, v...)
}

// Warning send a log message with 'WARNING' as prefix to Logger dbus service and print it to console, too.
func (l *Logger) Warning(v ...interface{}) {
	l.log(LevelWarning, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.logf(LevelWarning, format, v...)
}

// Error send a log message with 'ERROR' as prefix to Logger dbus service and print it to console, too.
func (l *Logger) Error(v ...interface{}) {
	l.log(LevelError, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logf(LevelError, format, v...)
}

// Panic is equivalent to Error() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	l.log(LevelPanic, v...)
	s := buildMsg(2, true, v...)
	panic(s)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logf(LevelPanic, format, v...)
	s := buildFormatMsg(2, true, format, v...)
	panic(s)
}

// Fatal send a log message with 'FATAL' as prefix to Logger dbus service
// and print it to console, then call os.Exit(1).
func (l *Logger) Fatal(v ...interface{}) {
	l.log(LevelFatal, v...)
	l.launchCrashReporter() // TODO
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logf(LevelFatal, format, v...)
	l.launchCrashReporter()
	os.Exit(1)
}

func (l *Logger) launchCrashReporter() {
	if logapi == nil {
		return
	}
	// if deepin-crash-reporter exists, launch it
	if isFileExists(crashReporterExe) {
		// save config to a temporary json file
		l.config.LogDetail, _ = logapi.GetLog(l.id)
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
