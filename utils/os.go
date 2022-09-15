// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

// #include <stdlib.h>
import "C"
import "unsafe"

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

func IsEnvExists(envName string) (ok bool) {
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, envName+"=") {
			ok = true
			break
		}
	}
	return
}

func UnsetEnv(envName string) (err error) {
	doUnsetEnvC(envName) // call C.unsetenv() is necessary
	envs := os.Environ()
	newEnvsData := make(map[string]string)
	for _, e := range envs {
		a := strings.SplitN(e, "=", 2)
		var name, value string
		if len(a) == 2 {
			name = a[0]
			value = a[1]
		} else {
			name = a[0]
			value = ""
		}
		if name != envName {
			newEnvsData[name] = value
		}
	}
	os.Clearenv()
	for e, v := range newEnvsData {
		err = os.Setenv(e, v)
		if err != nil {
			return
		}
	}
	return
}

func doUnsetEnvC(envName string) {
	cname := C.CString(envName)
	defer C.free(unsafe.Pointer(cname))
	C.unsetenv(cname)
}

func GetUserName() string {
	info, err := user.Current()
	if err != nil {
		return ""
	}

	return info.Username
}

func ExecAndWait(timeout int, name string, arg ...string) (stdout, stderr string, err error) {
	cmd := exec.Command(name, arg...)
	var bufStdout, bufStderr bytes.Buffer
	cmd.Stdout = &bufStdout
	cmd.Stderr = &bufStderr
	err = cmd.Start()
	if err != nil {
		return
	}

	// wait for process finished
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		if err = cmd.Process.Kill(); err != nil {
			return
		}
		<-done
		err = fmt.Errorf("time out and process was killed")
	case err = <-done:
		stdout = bufStdout.String()
		stderr = bufStderr.String()
		if err != nil {
			return
		}
	}
	return
}
