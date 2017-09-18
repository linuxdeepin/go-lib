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

import "testing"

func TestSplitObjectPath(t *testing.T) {
	data := []struct {
		Path   ObjectPath
		Parent ObjectPath
		Base   string
	}{
		{"/com/deepin/Test", "/com/deepin", "Test"},
		{"/com/deepin/Test/0", "/com/deepin/Test", "0"},
		{"/com/deepin/Test/a", "/com/deepin/Test", "a"},
		{"/com/deepin/Test/abc", "/com/deepin/Test", "abc"},
		{"/", "/", ""},
		{"/abc", "/", "abc"},
		{"/com/deepin/Test/", "", ""},
		{"", "", ""},
	}

	for _, i := range data {
		p, b := splitObjectPath(i.Path)
		if p != i.Parent || b != i.Base {
			t.Errorf("splitObjectPath:(%q) get (%q,%q). It should be (%q, %q)", i.Path, p, b, i.Parent, i.Base)
		}
	}
}
