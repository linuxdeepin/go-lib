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
	"fmt"
	"os"
	"pkg.deepin.io/lib/glib-2.0"
	"sync"
)

var (
	kfLocker sync.Mutex
)

func NewKeyFileFromFile(file string) (*glib.KeyFile, error) {
	kfLocker.Lock()
	defer kfLocker.Unlock()

	var kFile = glib.NewKeyFile()
	_, err := kFile.LoadFromFile(file, glib.KeyFileFlagsKeepComments|
		glib.KeyFileFlagsKeepTranslations)
	if err != nil {
		kFile.Free()
		return nil, err
	}

	return kFile, nil
}

func ReadKeyFromKeyFile(filename, group, key string, t interface{}) (interface{}, bool) {
	if !IsFileExist(filename) {
		return nil, false
	}

	kFile, err := NewKeyFileFromFile(filename)
	if err != nil {
		return nil, false
	}
	defer kFile.Free()

	switch t.(type) {
	case bool:
		value, err := kFile.GetBoolean(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	case []bool:
		_, value, err := kFile.GetBooleanList(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	case int, int32, uint32:
		value, err := kFile.GetInteger(group, key)
		if err != nil {
			return nil, false
		}
		return int32(value), true
	case int64:
		value, err := kFile.GetInt64(group, key)
		if err != nil {
			return nil, false
		}
		return int64(value), true
	case uint64:
		value, err := kFile.GetUint64(group, key)
		if err != nil {
			return nil, false
		}
		return uint64(value), true
	case []int, []int32, []uint32, []int64, []uint64:
		_, value, err := kFile.GetIntegerList(group, key)
		if err != nil {
			return nil, false
		}
		list := []int32{}
		for _, v := range value {
			list = append(list, int32(v))
		}
		return list, true
	case float32, float64:
		value, err := kFile.GetDouble(group, key)
		if err != nil {
			return nil, false
		}
		return float64(value), true
	case []float32, []float64:
		_, value, err := kFile.GetDoubleList(group, key)
		if err != nil {
			return nil, false
		}
		list := []float64{}
		for _, v := range value {
			list = append(list, float64(v))
		}
		return list, true
	case string:
		value, err := kFile.GetString(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	case []string:
		_, value, err := kFile.GetStringList(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	default:
		value, err := kFile.GetValue(group, key)
		if err != nil {
			return nil, false
		}
		return value, true
	}

	return nil, false
}

func WriteKeyToKeyFile(filename, group, key string, value interface{}) bool {
	if len(filename) == 0 {
		return false
	}

	if !IsFileExist(filename) {
		err := CreateFile(filename)
		if err != nil {
			return false
		}
	}

	kFile, err := NewKeyFileFromFile(filename)
	if err != nil {
		return false
	}
	defer kFile.Free()

	switch value.(type) {
	case bool:
		kFile.SetBoolean(group, key, value.(bool))
	case []bool:
		kFile.SetBooleanList(group, key, value.([]bool))
	case int:
		kFile.SetInteger(group, key, int32(value.(int)))
	case int32:
		kFile.SetInteger(group, key, int32(value.(int32)))
	case uint32:
		kFile.SetInteger(group, key, int32(value.(uint32)))
	case []int:
		list := value.([]int)
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case []int32:
		kFile.SetIntegerList(group, key, value.([]int32))
	case []uint32:
		list := value.([]uint32)
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case []int64:
		list := value.([]int64)
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case []uint64:
		list := value.([]uint64)
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case int64:
		kFile.SetInt64(group, key, value.(int64))
	case uint64:
		kFile.SetUint64(group, key, value.(uint64))
	case float32:
		kFile.SetDouble(group, key, float64(value.(float32)))
	case float64:
		kFile.SetDouble(group, key, value.(float64))
	case []float32:
		list := value.([]float32)
		tmp := []float64{}
		for _, l := range list {
			tmp = append(tmp, float64(l))
		}
		kFile.SetDoubleList(group, key, tmp)
	case []float64:
		kFile.SetDoubleList(group, key, value.([]float64))
	case string:
		kFile.SetString(group, key, value.(string))
	case []string:
		kFile.SetStringList(group, key, value.([]string))
	}

	_, contents, err1 := kFile.ToData()
	if err1 != nil {
		return false
	}

	ok := WriteStringToKeyFile(filename, string(contents))
	if !ok {
		return false
	}

	return true
}

//TODO: Abandoned it
//Do't use this interface.
func WriteStringToKeyFile(filename, content string) bool {
	err := WriteStringToFile(filename, content)
	if err != nil {
		return false
	}

	return true
}

func WriteStringToFile(filename, content string) error {
	if len(filename) == 0 {
		return fmt.Errorf("Not found this file: %q", filename)
	}

	kfLocker.Lock()
	defer kfLocker.Unlock()
	var swapFile = filename + ".swap"
	fp, err := os.Create(filename + ".swap")
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.WriteString(content)
	if err != nil {
		return err
	}
	fp.Sync()
	return os.Rename(swapFile, filename)
}
