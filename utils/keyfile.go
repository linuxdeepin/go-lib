// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"fmt"
	"os"
	"sync"

	"github.com/linuxdeepin/go-gir/glib-2.0"
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

	switch vType := value.(type) {
	case bool:
		kFile.SetBoolean(group, key, vType)
	case []bool:
		kFile.SetBooleanList(group, key, vType)
	case int:
		kFile.SetInteger(group, key, int32(vType))
	case int32:
		kFile.SetInteger(group, key, int32(vType))
	case uint32:
		kFile.SetInteger(group, key, int32(vType))
	case []int:
		list := vType
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case []int32:
		kFile.SetIntegerList(group, key, vType)
	case []uint32:
		list := vType
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case []int64:
		list := vType
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case []uint64:
		list := vType
		tmp := []int32{}
		for _, l := range list {
			tmp = append(tmp, int32(l))
		}
		kFile.SetIntegerList(group, key, tmp)
	case int64:
		kFile.SetInt64(group, key, vType)
	case uint64:
		kFile.SetUint64(group, key, vType)
	case float32:
		kFile.SetDouble(group, key, float64(vType))
	case float64:
		kFile.SetDouble(group, key, vType)
	case []float32:
		list := vType
		tmp := []float64{}
		for _, l := range list {
			tmp = append(tmp, float64(l))
		}
		kFile.SetDoubleList(group, key, tmp)
	case []float64:
		kFile.SetDoubleList(group, key, vType)
	case string:
		kFile.SetString(group, key, vType)
	case []string:
		kFile.SetStringList(group, key, vType)
	}

	_, contents, err1 := kFile.ToData()
	if err1 != nil {
		return false
	}

	ok := WriteStringToKeyFile(filename, string(contents))
	return ok
}

//TODO: Abandoned it
//Do't use this interface.
func WriteStringToKeyFile(filename, content string) bool {
	err := WriteStringToFile(filename, content)
	return err == nil
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
	defer func() {
		_ = fp.Close()
	}()

	_, err = fp.WriteString(content)
	if err != nil {
		return err
	}
	_ = fp.Sync()
	return os.Rename(swapFile, filename)
}
