// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"reflect"
)

func IsElementEqual(e1, e2 interface{}) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	return reflect.DeepEqual(e1, e2)
}

func IsElementInList(e interface{}, list interface{}) bool {
	if list == nil {
		return false
	}

	v := reflect.ValueOf(list)
	if !v.IsValid() {
		return false
	}

	if v.Type().Kind() == reflect.Slice ||
		v.Type().Kind() == reflect.Array {
		l := v.Len()
		for i := 0; i < l; i++ {
			if IsElementEqual(e, v.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

func GenUuid() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic("This can failed?")
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func RandString(n int) string {
	const alphanum = "0123456789abcdef"
	var bytes = make([]byte, n)
	_, _ = rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func IsInterfaceNil(v interface{}) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)

	// The argument must be a chan, func, interface, map, pointer, or
	// slice value; if it is not, Value.IsNil panics.
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return value.IsNil()
	}

	// should be a not nil type for rest cases
	return false
}
