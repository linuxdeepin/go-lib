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
	"log"
	"os"
	"runtime"
)

var logapi *Logapi

// Logger is a wrapper object to access Logger dbus service.
type Logger struct {
	name string
	id   uint64
}

func initLogapi() (err error) {
	if logapi == nil {
		logapi, err = NewLogapi("/com/deepin/api/Logger")
	}
	return
}

// DestroyLogger TODO[remove]
func DestroyLogger(obj *Logger) {
	Println("DestroyLogger")
	logapi.DeleteLogger(obj.id)
}

// New create a new Logger object, it need a string as name to
// register Logger dbus service.
func New(name string) (logger *Logger, err error) {
	err = initLogapi()
	if err != nil {
		return
	}

	logger = &Logger{name: name}
	logger.id, err = logapi.NewLogger(name)
	if err != nil {
		return
	}

	// TODO fix
	// runtime.SetFinalizer(logger, DestroyLogger)
	runtime.SetFinalizer(logger, func(obj *Logger) { DestroyLogger(obj) })

	return
}

func buildMsg(calldepth int, format string, v ...interface{}) string {
	s := fmt.Sprintf(format, v...)
	_, file, line, _ := runtime.Caller(calldepth)
	return fmt.Sprintf("%s:%d : %s", file, line, s)
}

// Println print message to console directly.
func Println(v ...interface{}) {
	r := fmt.Sprintln(v...)
	log.Printf(buildMsg(2, r))
}

// Printf print message to console directly.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	r := fmt.Sprintf(format, v...)
	log.Printf(buildMsg(2, r))
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

// Debug send a log message with 'DEBUG' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Debug(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	logapi.Debug(logger.id, s)
	log.Println("[DEBUG] " + s)
}

// Info send a log message with 'INFO' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Info(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	logapi.Info(logger.id, s)
	log.Println("[INFO] " + s)
}

// Warning send a log message with 'WARNING' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Warning(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	logapi.Warning(logger.id, s)
	log.Println("[WARNING] " + s)
}

// Error send a log message with 'ERROR' as prefix to Logger dbus service and print it to console, too.
func (logger *Logger) Error(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	logapi.Error(logger.id, s)
	log.Println("[ERROR] " + s)
}

// Panic is equivalent to Error() followed by a call to panic().
func (logger *Logger) Panic(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	logapi.Error(logger.id, s)
	log.Println("[PANIC] " + s)
	panic(s)
}

// Fatal send a log message with 'FATAL' as prefix to Logger dbus service
// and print it to console, then call os.Exit(1).
func (logger *Logger) Fatal(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	logapi.Fatal(logger.id, s)
	log.Println("[FATAL] " + s)
	os.Exit(1)
}
