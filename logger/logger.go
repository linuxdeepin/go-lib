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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	defaultDebugEnv  = "DDE_DEBUG"
	defaultDebugFile = "/var/cache/dde_debug"
	crashReporterExe = "/usr/bin/deepin-crash-reporter"
)

var crashReporterArgs = []string{crashReporterExe, "--remove-config", "--config"}

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

	// DebugFile if the file name that if exist the default log level
	// will be "LEVEL_DEBUG".
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

// NewLogger create a Logger object, it need a string as name to register
// Logger dbus service, if the environment variable exists which name
// stores in variable "DebugEnv", the default log level will be
// "LEVEL_DEBUG" or is "LEVEL_INFO".
func NewLogger(name string) (logger *Logger) {
	// ignore panic
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	logger = &Logger{name: name}
	if isEnvExists(DebugEnv) || isFileExists(DebugFile) {
		logger.level = LEVEL_DEBUG
	} else {
		logger.level = LEVEL_INFO
	}
	logger.config = newRestartConfig(name)

	err := initLogapi()
	if err != nil {
		log.Printf("init logger dbus api failed: %v\n", err)
		return
	}
	logger.id, err = logapi.NewLogger(name)
	if err != nil {
		log.Printf("create logger api object failed: %v\n", err)
		return
	}
	return
}

// SetLogLevel reset the log level.
func (logger *Logger) SetLogLevel(level Priority) *Logger {
	logger.level = level
	return logger
}

// SetRestartCommand reset the command and argument when restart after fatal.
func (logger *Logger) SetRestartCommand(exefile string, args ...string) {
	logger.config.RestartCommand = append([]string{exefile}, args...)
}

// AddExtArgForRestart add the command option which be used when
// process fataled and restart by Logger dbus service.
func (logger *Logger) AddExtArgForRestart(arg string) {
	if !stringInSlice(arg, logger.config.RestartCommand[1:]) {
		logger.config.RestartCommand = append(logger.config.RestartCommand, arg)
	}
}

func (logger *Logger) log(level Priority, v ...interface{}) {
	if level < logger.level {
		return
	}
	var s string
	if level >= LEVEL_ERROR {
		s = buildMsg(3, true, v...)
	} else {
		s = buildMsg(3, false, v...)
	}
	logger.doLog(level, s)
}

func (logger *Logger) logf(level Priority, format string, v ...interface{}) {
	if level < logger.level {
		return
	}
	var s string
	if level >= LEVEL_ERROR {
		s = buildFormatMsg(3, true, format, v...)
	} else {
		s = buildFormatMsg(3, false, format, v...)
	}
	logger.doLog(level, s)
}

func (logger *Logger) doLog(level Priority, s string) {
	switch level {
	case LEVEL_DEBUG:
		if logapi != nil {
			logapi.Debug(logger.id, s)
		}
		logger.printLocal("[DEBUG]", s)
	case LEVEL_INFO:
		if logapi != nil {
			logapi.Info(logger.id, s)
		}
		logger.printLocal("[INFO]", s)
	case LEVEL_WARNING:
		if logapi != nil {
			logapi.Warning(logger.id, s)
		}
		logger.printLocal("[WARNING]", s)
	case LEVEL_ERROR:
		if logapi != nil {
			logapi.Error(logger.id, s)
		}
		logger.printLocal("[ERROR]", s)
	case LEVEL_PANIC:
		if logapi != nil {
			logapi.Error(logger.id, s)
		}
		logger.printLocal("[PANIC]", s)
	case LEVEL_FATAL:
		if logapi != nil {
			logapi.Fatal(logger.id, s)
		}
		logger.printLocal("[FATAL]", s)
	}
}

func (logger *Logger) printLocal(prefix, msg string) {
	fmtmsg := prefix + " " + msg
	fmtmsg = strings.Replace(fmtmsg, "\n", "\n"+prefix+" ", -1) // format multi-lines message
	fmt.Println(fmtmsg)
}

// Debug send a log message with 'DEBUG' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Debug(v ...interface{}) {
	logger.log(LEVEL_DEBUG, v...)
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.logf(LEVEL_DEBUG, format, v...)
}

// Info send a log message with 'INFO' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Info(v ...interface{}) {
	logger.log(LEVEL_INFO, v...)
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.logf(LEVEL_INFO, format, v...)
}

// Warning send a log message with 'WARNING' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Warning(v ...interface{}) {
	logger.log(LEVEL_WARNING, v...)
}

func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.logf(LEVEL_WARNING, format, v...)
}

// Error send a log message with 'ERROR' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Error(v ...interface{}) {
	logger.log(LEVEL_ERROR, v...)
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.logf(LEVEL_ERROR, format, v...)
}

// Panic is equivalent to Error() followed by a call to panic().
func (logger *Logger) Panic(v ...interface{}) {
	logger.log(LEVEL_PANIC, v...)
	s := buildMsg(2, true, v...)
	panic(s)
}

func (logger *Logger) Panicf(format string, v ...interface{}) {
	logger.logf(LEVEL_PANIC, format, v...)
	s := buildFormatMsg(2, true, format, v...)
	panic(s)
}

// Fatal send a log message with 'FATAL' as prefix to Logger dbus service
// and print it to console, then call os.Exit(1).
func (logger *Logger) Fatal(v ...interface{}) {
	logger.log(LEVEL_FATAL, v...)
	logger.launchCrashReporter()
	os.Exit(1)
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.logf(LEVEL_FATAL, format, v...)
	logger.launchCrashReporter()
	os.Exit(1)
}

func (logger *Logger) launchCrashReporter() {
	if logapi == nil {
		return
	}
	// if deepin-crash-reporter exists, launch it
	if isFileExists(crashReporterExe) {
		// save config to a temporary json file
		logger.config.LogDetail, _ = logapi.GetLog(logger.id)
		fileContent, err := json.Marshal(logger.config)
		if err != nil {
			logger.Error("%v", err)
		}

		// create temporary json file and it will be removed by deepin-crash-reporter
		f, err := ioutil.TempFile("", "deepin_crash_reporter_config_")
		defer f.Close()
		if err != nil {
			logger.Error("%v", err)
		}
		_, err = f.Write(fileContent)
		if err != nil {
			logger.Error("%v", err)
		}

		// launch crash reporter
		logger.Info("launch deepin-crash-reporter: %s %s", crashReporterExe, append(crashReporterArgs, f.Name()))
		_, err = os.StartProcess(crashReporterExe, append(crashReporterArgs, f.Name()),
			&os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		if err != nil {
			logger.Error("launch deepin-crash-reporter failed: %v", err)
		}
	}
}
