/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
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

package kv

import (
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	Convey("Test Reader", t, func(c C) {
		f, err := os.Open("./testdata/a")
		c.So(err, ShouldBeNil)
		c.So(f, ShouldNotBeNil)
		defer f.Close()

		r := NewReader(f)
		r.Delim = ':'
		r.TrimSpace = TrimDelimRightSpace | TrimTailingSpace

		c.Convey("Get Pid", func(c C) {
			var resultPair *Pair
			for {
				pair, err := r.Read()
				if err != nil {
					break
				}
				if pair.Key == "Pid" {
					resultPair = pair
					break
				}
			}
			c.So(resultPair, ShouldNotBeNil)
			c.So(resultPair.Value, ShouldEqual, "21722")
		})

		c.Convey("ReadAll", func(c C) {
			pairs, err := r.ReadAll()
			c.So(err, ShouldBeNil)
			c.So(len(pairs), ShouldEqual, 48)
		})

	})

	Convey("Test ReadAll error", t, func(c C) {
		f, err := os.Open("./testdata/b")
		c.So(err, ShouldBeNil)
		c.So(f, ShouldNotBeNil)
		defer f.Close()

		r := NewReader(f)
		pairs, err := r.ReadAll()
		c.So(err, ShouldEqual, ErrBadLine)
		c.So(pairs, ShouldBeNil)
	})

	Convey("Test Read shell vars", t, func(c C) {
		f, err := os.Open("./testdata/c")
		c.So(err, ShouldBeNil)
		c.So(f, ShouldNotBeNil)
		defer f.Close()

		r := NewReader(f)
		r.TrimSpace = TrimLeadingTailingSpace
		r.Comment = '#'

		pair, err := r.Read()
		c.So(err, ShouldBeNil)
		c.So(pair, ShouldResemble, &Pair{"LANG", "zh_CN.UTF-8"})

		pair, err = r.Read()
		c.So(err, ShouldBeNil)
		c.So(pair, ShouldResemble, &Pair{"LANGUAGE", "zh_CN"})

		pair, err = r.Read()
		c.So(pair, ShouldBeNil)
		c.So(err, ShouldEqual, io.EOF)
	})
}
