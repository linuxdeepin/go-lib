/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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

package initializer_test

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "pkg.deepin.io/lib/initializer"
)

func TestInitializer(t *testing.T) {
	Convey("test initializer with success", t, func() {
		err := NewInitializer().Init(func(interface{}) (interface{}, error) {
			return 1, nil
		}).Init(func(v interface{}) (interface{}, error) {
			So(v, ShouldEqual, 1)
			return nil, nil
		}).GetError()

		So(err, ShouldBeNil)
	})
}

func TestInitializerError(t *testing.T) {
	Convey("test initializer with error", t, func() {
		var err error
		So(func() {
			err = NewInitializer().Init(func(interface{}) (interface{}, error) {
				return 1, nil
			}).Init(func(v interface{}) (interface{}, error) {
				So(v, ShouldEqual, 1)
				return nil, nil
			}).Init(func(interface{}) (interface{}, error) {
				return nil, errors.New("initialize error")
			}).Init(func(interface{}) (interface{}, error) {
				panic("should not be executed")
			}).GetError()

		}, ShouldNotPanic)

		So(err, ShouldNotBeNil)
	})
}
