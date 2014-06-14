/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
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
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package utils

import (
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

func deleteStartSpace(str string) string {
	if len(str) <= 0 {
		return ""
	}

	tmp := strings.TrimLeft(str, " ")

	return tmp
}

func (op *Manager) CopyFile(src, dest string) bool {
	if ok := op.IsFileExist(src); !ok && len(dest) <= 0 {
		return false
	}

	contents, err := ioutil.ReadFile(src)
	if err != nil {
		logger.Infof("ReadFile '%s' failed: %v", src, err)
		return false
	}

	f, err1 := os.Create(dest + "~")
	if err1 != nil {
		logger.Infof("Create '%s' failed: %v", dest+"~", err1)
		return false
	}
	defer f.Close()

	if _, err := f.Write(contents); err != nil {
		logger.Infof("Write '%s' failed: %v", dest+"~", err)
		return false
	}
	f.Sync()

	if err := os.Rename(dest+"~", dest); err != nil {
		logger.Infof("Rename '%s' failed: %v", dest+"~", err)
		return false
	}

	return true
}

func (op *Manager) GetBaseName(path string) (string, bool) {
	if len(path) <= 0 {
		return "", false
	}

	as := strings.Split(path, "/")
	if l := len(as); l > 1 {
		return as[l-1], true
	}

	return "", false
}

func (op *Manager) IsContainFromStart(str, substr string) bool {
	l1 := len(substr)
	l2 := len(str)

	l := 0
	if l1 > l2 {
		l = l2
	} else {
		l = l1
	}

	for i := 0; i < l; i++ {
		if str[i] != substr[i] {
			return false
		}
	}

	return true
}

func (op *Manager) IsFileExist(filename string) bool {
	if len(filename) <= 0 {
		return false
	}

	path, ok := op.URIToPath(filename)
	if !ok {
		return false
	}
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func (op *Manager) IsElementExist(e interface{}, l interface{}) bool {
	if e == nil || l == nil {
		panic("args is nil in IsElementExist")
	}

	switch e.(type) {
	case int32:
		element := e.(int32)
		list := l.([]int32)
		for _, v := range list {
			if element == v {
				return true
			}
		}
	case uint32:
		element := e.(uint32)
		list := l.([]uint32)
		for _, v := range list {
			if element == v {
				return true
			}
		}
	case int64:
		element := e.(int64)
		list := l.([]int64)
		for _, v := range list {
			if element == v {
				return true
			}
		}
	case uint64:
		element := e.(uint64)
		list := l.([]uint64)
		for _, v := range list {
			if element == v {
				return true
			}
		}
	case string:
		element := e.(string)
		list := l.([]string)
		for _, v := range list {
			if element == v {
				return true
			}
		}
	case bool:
		element := e.(bool)
		list := l.([]bool)
		for _, v := range list {
			if element == v {
				return true
			}
		}
	case byte:
		element := e.(byte)
		list := l.([]byte)
		for _, v := range list {
			if element == v {
				return true
			}
		}
	}

	return false
}

func (op *Manager) IsListEqual(l1, l2 interface{}) bool {
	if l1 == nil || l2 == nil {
		panic("Args is nil in IsListEqual")
	}

	switch l1.(type) {
	case []int32:
		list1 := l1.([]int32)
		list2 := l2.([]int32)

		len1 := len(list1)
		len2 := len(list2)

		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	case []uint32:
		list1 := l1.([]uint32)
		list2 := l2.([]uint32)

		len1 := len(list1)
		len2 := len(list2)

		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	case []int64:
		list1 := l1.([]int64)
		list2 := l2.([]int64)

		len1 := len(list1)
		len2 := len(list2)

		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	case []uint64:
		list1 := l1.([]uint64)
		list2 := l2.([]uint64)

		len1 := len(list1)
		len2 := len(list2)

		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	case []string:
		list1 := l1.([]string)
		list2 := l2.([]string)

		len1 := len(list1)
		len2 := len(list2)

		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	case []bool:
		list1 := l1.([]bool)
		list2 := l2.([]bool)

		len1 := len(list1)
		len2 := len(list2)

		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	case []byte:
		list1 := l1.([]byte)
		list2 := l2.([]byte)

		len1 := len(list1)
		len2 := len(list2)

		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	}

	return true
}

func (op *Manager) GetHomeDir() (string, bool) {
	info, err := user.Current()
	if err != nil {
		logger.Warning("Get User Info Failed: ", err)
		return "", false
	}

	return info.HomeDir, true
}

func (op *Manager) GetConfigDir() (string, bool) {
	if homeDir, ok := op.GetHomeDir(); ok {
		return homeDir + "/.config", true
	}
	return "", false
}

func (op *Manager) GetCacheDir() (string, bool) {
	if homeDir, ok := op.GetHomeDir(); ok {
		return homeDir + "/.cache", true
	}
	return "", false
}

func (op *Manager) UnsetEnv(envName string) {
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
		err := os.Setenv(e, v)
		if err != nil {
			logger.Info(e, v, err)
		}
	}
}
