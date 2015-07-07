package initializer_test

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	. "pkg.deepin.io/lib/initializer"
	"testing"
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
				return nil, nil
			}).GetError()

		}, ShouldNotPanic)

		So(err, ShouldNotBeNil)
	})
}
