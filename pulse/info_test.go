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
	if len(ctx.GetCardList()) == 0 {
		t.Fatal("Can't query sound card")
	}
}
