/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package pulse

import "testing"

import "math/rand"

func BenchmarkTheadSafe(b *testing.B) {
	ctx := GetContext()
	sinks := ctx.GetSinkList()
	for _, s := range sinks {
		old := s.Volume
		for i := 0; i < b.N; i++ {
			v := s.Volume.SetAvg(rand.Float64())
			s.SetVolume(v)
		}
		s.SetVolume(old)
	}
}
