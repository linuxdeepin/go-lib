/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package timer_test

import (
	. "github.com/smartystreets/goconvey/convey"
	. "pkg.deepin.io/lib/timer"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	Convey("just stop", t, func() {
		timer := NewTimer()
		timer.Stop()
		So(timer.Elapsed(), ShouldEqual, 0)
	})

	Convey("get elapse without stop", t, func() {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second*2-time.Millisecond*100, time.Second*2+time.Millisecond*100)
	})

	Convey("stop and elapse", t, func() {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		timer.Stop()
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)
	})

	Convey("stop and continue", t, func() {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		timer.Stop()
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		timer.Continue()
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		time.Sleep(time.Second)
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second*2-time.Millisecond*100, time.Second*2+time.Millisecond*100)
	})

	Convey("reset", t, func() {
		timer := NewTimer()
		timer.Start()

		time.Sleep(time.Second)
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

		timer.Reset()
		time.Sleep(time.Second)
		So(timer.Elapsed(), ShouldBeBetweenOrEqual, time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)
	})
}
