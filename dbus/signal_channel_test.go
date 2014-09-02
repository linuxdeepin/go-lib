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
