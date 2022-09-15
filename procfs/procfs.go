// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package procfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/linuxdeepin/go-lib/encoding/kv"
	"strconv"
	"strings"
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
