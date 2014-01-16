package logger

import (
	"fmt"
	"log"
	"runtime"
)

func Println(v ...interface{}) {
	r := fmt.Sprintln(v...)
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d :%s", file, line, r)
}

func Printf(format string, v ...interface{}) {
	r := fmt.Sprintf(format, v...)
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d :%s", file, line, r)
}

func Assert(exp bool, v ...interface{}) {
	if exp == false {
		panic(fmt.Sprintln(v...))
	}
}
func AssertNotReached() {
	panic("Shouldn't reached")
}
