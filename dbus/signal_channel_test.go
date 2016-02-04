/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package dbus

import "testing"

//Note: run "go test" with timeout flags
func TestSignalChannle(t *testing.T) {
	ch := newSignalChannel()
	loops := []int{0, 1, 10, 100, 1000, 10000, 1000000}
	for _, count := range loops {
		for i := 0; i < count; i++ {
			ch.In() <- &Signal{}
		}
		for i := 0; i < count; i++ {
			<-ch.Out()
		}
		if len(ch.In()) != 0 || len(ch.Out()) != 0 || len(ch.caches) > 1 {
			t.Fatal("Count: %d  %d %d %d\n", count, len(ch.In()), len(ch.Out()), len(ch.caches))
		}
	}
}
