package initializer_test

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"pkg.deepin.io/lib/initializer/v2"
	"testing"
)

func e1() error {
	return nil
}

func e2() error {
	var err error
	return err
}

func TestInitializer(t *testing.T) {
	Convey("test initializer with success", t, func() {
		err := initializer.Do(func() error {
			return nil
		}).Do(func() error {
			return e1()
		}).Do(func() error {
			return e2()
		}).GetError()

		So(err, ShouldBeNil)
	})
}

func TestInitializerError(t *testing.T) {
	Convey("test initializer with error", t, func() {
		var err error
		So(func() {
			err = initializer.Do(func() error {
				return e1()
			}).Do(func() error {
				return e2()
			}).Do(func() error {
				return errors.New("initialize error")
			}).Do(func() error {
				panic("should not be executed")
				return nil
			}).GetError()

		}, ShouldNotPanic)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "initialize error")
	})
}
