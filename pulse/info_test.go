package pulse

import (
	"testing"
)

// Please manually enable test CFLAGS in ./pulse.go

func TestQueryInfo(t *testing.T) {
	ctx := GetContext()
	ctx = GetContextForced()
	if ctx == nil {
		t.Skip("Can't connect to pulseaudio.")
		return
	}
	_, err := ctx.GetServer()
	if err != nil {
		t.Fatal("Can't query server info", err)
	}
}
