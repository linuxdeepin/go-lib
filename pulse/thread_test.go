// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

import (
	"math/rand"
	"testing"
)

func BenchmarkTheadSafe(b *testing.B) {
	ctx := GetContext()
	sinks := ctx.GetSinkList()
	for _, s := range sinks {
		old := s.Volume
		for i := 0; i < b.N; i++ {
			v := s.Volume.SetAvg(rand.Float64())
			ctx.SetSinkVolumeByIndex(s.Index, v)
		}
		ctx.SetSinkVolumeByIndex(s.Index, old)
	}
}
