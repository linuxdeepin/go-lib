// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package initializer_test

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/linuxdeepin/go-lib/initializer/v2"
)

func e1() error {
	return nil
}

func e2() error {
	var err error
	return err
}

func TestInitializer(t *testing.T) {
	Convey("test initializer with success", t, func(c C) {
		err := initializer.Do(func() error {
			return nil
		}).Do(func() error {
			return e1()
		}).Do(func() error {
			return e2()
		}).GetError()

		c.So(err, ShouldBeNil)
	})
}

func TestInitializerError(t *testing.T) {
	Convey("test initializer with error", t, func(c C) {
		var err error
		c.So(func() {
			err = initializer.Do(func() error {
				return e1()
			}).Do(func() error {
				return e2()
			}).Do(func() error {
				return errors.New("initialize error")
			}).Do(func() error {
				panic("should not be executed")
			}).GetError()

		}, ShouldNotPanic)

		c.So(err, ShouldNotBeNil)
		c.So(err.Error(), ShouldEqual, "initialize error")
	})
}
