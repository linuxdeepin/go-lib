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

package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const defaultDebugEnv = "DDE_DEBUG"

type Priority int

// Definition of log levels, the larger of the value, the higher of
// the priority.
const (
	LEVEL_DEBUG Priority = iota
	LEVEL_INFO
	LEVEL_WARNING
	LEVEL_ERROR
	LEVEL_PANIC
	LEVEL_FATAL
)

var (
	logapi *Logapi

	// DebugEnv is the name of environment variable to control the
	// default log level, if exists the default log level will be
	// "LEVEL_DEBUG".
	DebugEnv = defaultDebugEnv
)

func initLogapi() (err error) {
	if logapi == nil {
		logapi, err = NewLogapi("/com/deepin/api/Logger")
	}
	return
}

// ProcessInfo store the process information which will be
// used to restart if application fataled.
type ProcessInfo struct {
	uid     int32    // Real user ID
	dir     string   // Working directory
	environ []string // Environment variables
	exefile string   // Program file
	args    []string // Command-line arguments
}

func getProcessInfo() *ProcessInfo {
	processInfo := &ProcessInfo{}
	processInfo.uid = int32(os.Getuid())
	processInfo.dir, _ = os.Getwd()
	processInfo.exefile, _ = filepath.Abs(os.Args[0])
	processInfo.args = os.Args[1:]
	processInfo.environ = os.Environ()
	return processInfo
}

func buildMsg(calldepth int, format string, v ...interface{}) string {
	s := fmt.Sprintf(format, v...)
	_, file, line, _ := runtime.Caller(calldepth)
	return fmt.Sprintf("%s:%d: %s", file, line, s)
}

// Println print message to console directly.
func Println(v ...interface{}) {
	r := fmt.Sprintln(v...)
	fmt.Printf(buildMsg(2, r))
}

// Printf print message to console directly.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	r := fmt.Sprintf(format, v...)
	fmt.Printf(buildMsg(2, r))
}

// Assert will check if a expression is true, or will call panic().
// Arguments are handled in the manner of fmt.Sprintln.
func Assert(exp bool, v ...interface{}) {
	if exp == false {
		panic(fmt.Sprintln(v...))
	}
}

// AssertNotReached is a helper function which just call panic().
func AssertNotReached() {
	panic("Shouldn't reached")
}

// Logger is a wrapper object to access Logger dbus service.
type Logger struct {
	name        string
	id          uint64
	level       Priority
	processInfo *ProcessInfo
}

// NewLogger create a Logger object, it need a string as name to register
// Logger dbus service, if the environment variable exists which name
// stores in variable "DebugEnv", the default log level will be
// "LEVEL_DEBUG" or is "LEVEL_INFO".
func NewLogger(name string) (logger *Logger, err error) {
	logger = &Logger{name: name}
	if isEnvExists(DebugEnv) {
		logger.level = LEVEL_DEBUG
	} else {
		logger.level = LEVEL_INFO
	}
	logger.processInfo = getProcessInfo()

	err = initLogapi()
	if err != nil {
		return
	}
	logger.id, err = logapi.NewLogger(name)
	if err != nil {
		return
	}
	return
}

// SetLogLevel reset the log level.
func (logger *Logger) SetLogLevel(level Priority) {
	logger.level = level
}

// AddExtArgForRestart set the command option which be used when
// process fataled and restart by Logger dbus service.
func (logger *Logger) AddExtArgForRestart(arg string) {
	if !stringInSlice(arg, logger.processInfo.args) {
		logger.processInfo.args = append(logger.processInfo.args, arg)
	}
}

func (logger *Logger) doLog(level Priority, format string, v ...interface{}) {
	if level < logger.level {
		return
	}

	s := buildMsg(3, format, v...)
	switch level {
	case LEVEL_DEBUG:
		if logapi != nil {
			logapi.Debug(logger.id, s)
		}
		fmt.Println("[DEBUG] " + s)
	case LEVEL_INFO:
		if logapi != nil {
			logapi.Info(logger.id, s)
		}
		fmt.Println("[INFO] " + s)
	case LEVEL_WARNING:
		if logapi != nil {
			logapi.Warning(logger.id, s)
		}
		fmt.Println("[WARNING] " + s)
	case LEVEL_ERROR:
		if logapi != nil {
			logapi.Error(logger.id, s)
		}
		fmt.Println("[ERROR] " + s)
	case LEVEL_PANIC:
		if logapi != nil {
			logapi.Error(logger.id, s)
		}
		fmt.Println("[PANIC] " + s)
	case LEVEL_FATAL:
		if logapi != nil {
			logapi.Fatal(logger.id, s)
		}
		fmt.Println("[FATAL] " + s)
	}
}

// Debug send a log message with 'DEBUG' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Debug(format string, v ...interface{}) {
	logger.doLog(LEVEL_DEBUG, format, v...)
}

// Info send a log message with 'INFO' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Info(format string, v ...interface{}) {
	logger.doLog(LEVEL_INFO, format, v...)
}

// Warning send a log message with 'WARNING' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Warning(format string, v ...interface{}) {
	logger.doLog(LEVEL_WARNING, format, v...)
}

// Error send a log message with 'ERROR' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Error(format string, v ...interface{}) {
	logger.doLog(LEVEL_ERROR, format, v...)
}

// Panic is equivalent to Error() followed by a call to panic().
func (logger *Logger) Panic(format string, v ...interface{}) {
	logger.doLog(LEVEL_PANIC, format, v...)
	s := buildMsg(2, format, v...)
	panic(s)
}

// Fatal send a log message with 'FATAL' as prefix to Logger dbus service
// and print it to console, then call os.Exit(1).
func (logger *Logger) Fatal(format string, v ...interface{}) {
	logger.doLog(LEVEL_FATAL, format, v...)
	logapi.NotifyRestart(logger.id, logger.processInfo.uid, logger.processInfo.dir,
		logger.processInfo.environ, logger.processInfo.exefile, logger.processInfo.args)
	os.Exit(1)
}
