// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package procfs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/linuxdeepin/go-lib/encoding/kv"
)

type Process uint

func (p Process) getFile(name string) string {
	return fmt.Sprintf("/proc/%d/%s", uint(p), name)
}

func (p Process) Exist() bool {
	procDir := fmt.Sprintf("/proc/%d", uint(p))
	_, err := os.Stat(procDir)
	return err == nil
}

func (p Process) Cmdline() ([]string, error) {
	cmdlineFile := p.getFile("cmdline")
	bytes, err := ioutil.ReadFile(cmdlineFile)
	if err != nil {
		return nil, err
	}
	content := string(bytes)
	parts := strings.Split(content, "\x00")
	length := len(parts)
	if length >= 2 && parts[length-1] == "" {
		return parts[:length-1], nil
	}
	return parts, nil
}

func (p Process) Cwd() (string, error) {
	cwdFile := p.getFile("cwd")
	cwd, err := os.Readlink(cwdFile)
	return cwd, err
}

func (p Process) Exe() (string, error) {
	exeFile := p.getFile("exe")
	exe, err := filepath.EvalSymlinks(exeFile)
	return exe, err
}

func (p Process) TrustedExe() (string, error) {
	exeFile := p.getFile("exe")
	exe, err := filepath.EvalSymlinks(exeFile)
	if err != nil {
		return "", err
	}
	if os.Getuid() == 0 {
		if !checkSenderNsMntValid(uint32(p)) {
			return "", errors.New("due to the difference between the current process's ns mnt and the init process's ns mnt, the exe field is not reliable")
		}
	}
	return exe, nil
}

type EnvVars []string

func (p Process) Environ() (EnvVars, error) {
	envFile := p.getFile("environ")
	contents, err := ioutil.ReadFile(envFile)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents), "\x00"), nil
}

func (vars EnvVars) Lookup(key string) (string, bool) {
	prefix := key + "="
	for _, aVar := range vars {
		if strings.HasPrefix(aVar, prefix) {
			return aVar[len(prefix):], true
		}
	}
	return "", false
}

func (vars EnvVars) Get(key string) string {
	ret, _ := vars.Lookup(key)
	return ret
}

type Status []*kv.Pair

func (p Process) Status() (Status, error) {
	statusFile := p.getFile("status")
	f, err := os.Open(statusFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := kv.NewReader(f)
	r.Delim = ':'
	r.TrimSpace = kv.TrimDelimRightSpace | kv.TrimTailingSpace

	pairs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return Status(pairs), nil
}

func (st Status) lookup(key string) (string, error) {
	for _, pair := range st {
		if pair.Key == key {
			return pair.Value, nil
		}
	}
	return "", StatusFieldNotFoundErr{key}
}

type StatusFieldNotFoundErr struct {
	Field string
}

func (err StatusFieldNotFoundErr) Error() string {
	return fmt.Sprintf("field %s is not found in proc status file", err.Field)
}

func (st Status) Uids() ([]uint, error) {
	uids := make([]uint, 0, 4)
	value, err := st.lookup("Uid")
	if err != nil {
		return nil, err
	}
	for _, i := range strings.Split(value, "\t") {
		v, err := strconv.ParseUint(i, 10, 32)
		if err != nil {
			return nil, err
		}
		uids = append(uids, uint(v))
	}
	return uids, nil
}

func (st Status) PPid() (uint, error) {
	value, err := st.lookup("PPid")
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

var _initProcNsMnt string
var _once sync.Once

// 通过判断/proc/pid/ns/mnt 和 /proc/1/ns/mnt是否相同，如果不相同，则进程exe字段不可信
func checkSenderNsMntValid(pid uint32) bool {
	_once.Do(func() {
		out, err := os.Readlink("/proc/1/ns/mnt")
		if err != nil {
			fmt.Println(err)
			return
		}
		_initProcNsMnt = strings.TrimSpace(out)
	})
	c, err := os.Readlink(fmt.Sprintf("/proc/%v/ns/mnt", pid))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer func() {
		fmt.Printf("pid 1 mnt ns is %v,pid %v mnt ns is %v\n", _initProcNsMnt, pid, strings.TrimSpace(c))
	}()
	return strings.TrimSpace(c) == _initProcNsMnt
}
