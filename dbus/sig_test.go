/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package dbus

import (
	"testing"
)

var sigTests = []struct {
	vs  []interface{}
	sig Signature
}{
	{
		[]interface{}{new(int32)},
		Signature{"i"},
	},
	{
		[]interface{}{new(string)},
		Signature{"s"},
	},
	{
		[]interface{}{new(Signature)},
		Signature{"g"},
	},
	{
		[]interface{}{new([]int16)},
		Signature{"an"},
	},
	{
		[]interface{}{new(int16), new(uint32)},
		Signature{"nu"},
	},
	{
		[]interface{}{new(map[byte]Variant)},
		Signature{"a{yv}"},
	},
	{
		[]interface{}{new(Variant), new([]map[int32]string)},
		Signature{"vaa{is}"},
	},
}

func TestSig(t *testing.T) {
	for i, v := range sigTests {
		sig := SignatureOf(v.vs...)
		if sig != v.sig {
			t.Errorf("test %d: got %q, expected %q", i+1, sig.str, v.sig.str)
		}
	}
}

var getSigTest = []interface{}{
	[]struct {
		b byte
		i int32
		t uint64
		s string
	}{},
	map[string]Variant{},
}

func BenchmarkGetSignatureSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignatureOf("", int32(0))
	}
}

func BenchmarkGetSignatureLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignatureOf(getSigTest...)
	}
}
