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
	"os"
	"pkg.linuxdeepin.com/lib/glib-2.0"
	"sync"
)

func ReadKeyFromKeyFile(filename, group, key string, t interface{}) (interface{}, bool) {
	if len(filename) <= 0 || !IsFileExist(filename) {
		return nil, false
	}
	rwMutex := new(sync.RWMutex)
	rwMutex.Lock()
	defer rwMutex.Unlock()

	keyFile := glib.NewKeyFile()
	defer keyFile.Free()
	_, err := keyFile.LoadFromFile(filename,
		glib.KeyFileFlagsKeepComments)
	if err != nil {
		return nil, false
	}

	switch t.(type) {
	case bool:
		value, err := keyFile.GetBoolean(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	case []bool:
		_, value, err := keyFile.GetBooleanList(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	case int, int32, uint32:
		value, err := keyFile.GetInteger(group, key)
		if err != nil {
			return nil, false
		}
		return int32(value), true
	case int64:
		value, err := keyFile.GetInt64(group, key)
		if err != nil {
			return nil, false
		}
		return int64(value), true
	case uint64:
		value, err := keyFile.GetUint64(group, key)
		if err != nil {
			return nil, false
		}
		return uint64(value), true
	case []int, []int32, []uint32, []int64, []uint64:
		_, value, err := keyFile.GetIntegerList(group, key)
		if err != nil {
			return nil, false
		}
		list := []int32{}
		for _, v := range value {
			list = append(list, int32(v))
		}
		return list, true
	case float32, float64:
		value, err := keyFile.GetDouble(group, key)
		if err != nil {
			return nil, false
		}
		return float64(value), true
	case []float32, []float64:
		_, value, err := keyFile.GetDoubleList(group, key)
		if err != nil {
			return nil, false
		}
		list := []float64{}
		for _, v := range value {
			list = append(list, float64(v))
		}
		return list, true
	case string:
		value, err := keyFile.GetString(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	case []string:
		_, value, err := keyFile.GetStringList(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	default:
		value, err := keyFile.GetValue(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	}

	return nil, false
}

func WriteKeyToKeyFile(filename, group, key string, value interface{}) bool {
	if len(filename) <= 0 {
		return false
	}
	rwMutex := new(sync.RWMutex)
	rwMutex.Lock()
	defer rwMutex.Unlock()

	if !IsFileExist(filename) {
		f, err := os.Create(filename)
		if err != nil {
			return false
		}
		f.Close()
	}

	keyFile := glib.NewKeyFile()
	defer keyFile.Free()
	_, err := keyFile.LoadFromFile(filename,
		glib.KeyFileFlagsKeepComments)
	if err != nil {
		return false
	}

	switch value.(type) {
	case bool:
		keyFile.SetBoolean(group, key, value.(bool))
	case []bool:
		keyFile.SetBooleanList(group, key, value.([]bool))
	case int:
		keyFile.SetInteger(group, key, value.(int))
	case int32:
		keyFile.SetInteger(group, key, int(value.(int32)))
	case uint32:
		keyFile.SetInteger(group, key, int(value.(uint32)))
	case []int:
		keyFile.SetIntegerList(group, key, value.([]int))
	case []int32:
		list := value.([]int32)
		tmp := []int{}
		for _, l := range list {
			tmp = append(tmp, int(l))
		}
		keyFile.SetIntegerList(group, key, tmp)
	case []uint32:
		list := value.([]uint32)
		tmp := []int{}
		for _, l := range list {
			tmp = append(tmp, int(l))
		}
		keyFile.SetIntegerList(group, key, tmp)
	case []int64:
		list := value.([]int64)
		tmp := []int{}
		for _, l := range list {
			tmp = append(tmp, int(l))
		}
		keyFile.SetIntegerList(group, key, tmp)
	case []uint64:
		list := value.([]uint64)
		tmp := []int{}
		for _, l := range list {
			tmp = append(tmp, int(l))
		}
		keyFile.SetIntegerList(group, key, tmp)
	case int64:
		keyFile.SetInt64(group, key, value.(int64))
	case uint64:
		keyFile.SetUint64(group, key, value.(uint64))
	case float32:
		keyFile.SetDouble(group, key, float64(value.(float32)))
	case float64:
		keyFile.SetDouble(group, key, value.(float64))
	case []float32:
		list := value.([]float32)
		tmp := []float64{}
		for _, l := range list {
			tmp = append(tmp, float64(l))
		}
		keyFile.SetDoubleList(group, key, tmp)
	case []float64:
		keyFile.SetDoubleList(group, key, value.([]float64))
	case string:
		keyFile.SetString(group, key, value.(string))
	case []string:
		keyFile.SetStringList(group, key, value.([]string))
	}

	_, contents, err1 := keyFile.ToData()
	if err1 != nil {
		return false
	}

	ok := WriteStringToKeyFile(filename, string(contents))
	if !ok {
		return false
	}

	return true
}

func WriteStringToKeyFile(filename, contents string) bool {
	if len(filename) <= 0 {
		return false
	}

	f, err := os.Create(filename + "~")
	if err != nil {
		return false
	}
	defer f.Close()

	if _, err = f.WriteString(contents); err != nil {
		return false
	}
	f.Sync()
	os.Rename(filename+"~", filename)

	return true
}
