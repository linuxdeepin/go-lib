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
