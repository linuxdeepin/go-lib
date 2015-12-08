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
