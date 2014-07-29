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
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func IsInterfaceNil(v interface{}) bool {
	defer func() { recover() }()
	return v == nil || reflect.ValueOf(v).IsNil()
}
