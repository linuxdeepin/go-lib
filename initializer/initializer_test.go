// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package initializer_test

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/linuxdeepin/go-lib/initializer"
)

func TestInitializer(t *testing.T) {
	Convey("test initializer with success", t, func(c C) {
		err := NewInitializer().Init(func(interface{}) (interface{}, error) {
			return 1, nil
		}).Init(func(v interface{}) (interface{}, error) {
			c.So(v, ShouldEqual, 1)
			return nil, nil
		}).GetError()

		c.So(err, ShouldBeNil)
	})
}

func TestInitializerError(t *testing.T) {
	Convey("test initializer with error", t, func(c C) {
		var err error
		c.So(func() {
			err = NewInitializer().Init(func(interface{}) (interface{}, error) {
				return 1, nil
			}).Init(func(v interface{}) (interface{}, error) {
				c.So(v, ShouldEqual, 1)
				return nil, nil
			}).Init(func(interface{}) (interface{}, error) {
				return nil, errors.New("initialize error")
			}).Init(func(interface{}) (interface{}, error) {
				panic("should not be executed")
			}).GetError()

		}, ShouldNotPanic)

		c.So(err, ShouldNotBeNil)
	})
}
