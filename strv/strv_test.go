/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
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

package strv

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContains(t *testing.T) {
	Convey("Contains", t, func() {
		vector := Strv([]string{"a", "b", "c"})
		So(vector.Contains("b"), ShouldBeTrue)
		So(vector.Contains("d"), ShouldBeFalse)
	})
}

func TestEqual(t *testing.T) {
	Convey("Equal", t, func() {
		v1 := Strv([]string{"a", "b", "c"})
		v2 := Strv([]string{"a", "b", "c", "d"})
		v3 := Strv(v1[:])
		So(v1.Equal(v2), ShouldBeFalse)
		So(v1.Equal(v3), ShouldBeTrue)
	})
}

func TestUniq(t *testing.T) {
	Convey("Uniq", t, func() {
		vector := Strv([]string{"a", "b", "c", "c", "b", "a", "c"})
		vector = vector.Uniq()
		So(vector, ShouldResemble, Strv([]string{"a", "b", "c"}))
	})
}

func TestFilterFunc(t *testing.T) {
	Convey("FilterFunc", t, func() {
		vector := Strv([]string{"hello", "", "world", "", "!"})
		vector = vector.FilterFunc(func(str string) bool {
			return len(str) == 0
		})
		So(vector, ShouldResemble, Strv([]string{"hello", "world", "!"}))
	})
}

func TestFilterEmpty(t *testing.T) {
	Convey("FilterEmpty", t, func() {
		vector := Strv([]string{"hello", "", "world", "", "!"})
		vector = vector.FilterEmpty()
		So(vector, ShouldResemble, Strv([]string{"hello", "world", "!"}))
	})
}

func TestAdd(t *testing.T) {
	Convey("Add", t, func() {
		vector := Strv([]string{"a", "b", "c"})

		vector0, b0 := vector.Add("d")
		So(vector, ShouldResemble, Strv([]string{"a", "b", "c"}))
		So(vector0, ShouldResemble, Strv([]string{"a", "b", "c", "d"}))
		So(b0, ShouldBeTrue)

		vector1, b1 := vector.Add("c")
		So(vector, ShouldResemble, Strv([]string{"a", "b", "c"}))
		So(vector1, ShouldResemble, Strv([]string{"a", "b", "c"}))
		So(b1, ShouldBeFalse)
	})
}

func TestDelete(t *testing.T) {
	Convey("Delete", t, func() {
		vector := Strv([]string{"a", "b", "c"})
		vector0, b0 := vector.Delete("d")
		So(vector, ShouldResemble, Strv([]string{"a", "b", "c"}))
		So(vector0, ShouldResemble, Strv([]string{"a", "b", "c"}))
		So(b0, ShouldBeFalse)

		vector1, b1 := vector.Delete("c")
		So(vector, ShouldResemble, Strv([]string{"a", "b", "c"}))
		So(vector1, ShouldResemble, Strv([]string{"a", "b"}))
		So(b1, ShouldBeTrue)
	})
}
