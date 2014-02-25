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

const (
	_PATH = "/com/deepin/api/Logger"
)

var _LOGAPI *Logapi

type Logger struct {
	name string
	id   uint64
}

func initLogApi() (err error) {
	if _LOGAPI == nil {
		_LOGAPI, err = NewLogapi(_PATH)
	}
	return
}

func DestroyLogger(obj *Logger) {
	Println("DestroyLogger")
	_LOGAPI.DeleteLogger(obj.id)
}

func New(name string) (logger *Logger, err error) {
	err = initLogApi()
	if err != nil {
		return
	}

	logger = &Logger{name: name}
	logger.id, err = _LOGAPI.NewLogger(name)
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

func Println(v ...interface{}) {
	r := fmt.Sprintln(v...)
	log.Printf(buildMsg(2, r))
}

func Printf(format string, v ...interface{}) {
	r := fmt.Sprintf(format, v...)
	log.Printf(buildMsg(2, r))
}

func Assert(exp bool, v ...interface{}) {
	if exp == false {
		panic(fmt.Sprintln(v...))
	}
}

func AssertNotReached() {
	panic("Shouldn't reached")
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	_LOGAPI.Debug(logger.id, s)
	log.Println("[DEBUG] " + s)
}

func (logger *Logger) Info(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	_LOGAPI.Info(logger.id, s)
	log.Println("[INFO] " + s)
}

func (logger *Logger) Warning(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	_LOGAPI.Warning(logger.id, s)
	log.Println("[WARNING] " + s)
}

func (logger *Logger) Error(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	_LOGAPI.Error(logger.id, s)
	log.Println("[ERROR] " + s)
}

func (logger *Logger) Panic(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	_LOGAPI.Error(logger.id, s)
	log.Println("[PANIC] " + s)
	panic(s)
}

func (logger *Logger) Fatal(format string, v ...interface{}) {
	s := buildMsg(2, format, v...)
	_LOGAPI.Fatal(logger.id, s)
	log.Println("[FATAL] " + s)
	os.Exit(1)
}
