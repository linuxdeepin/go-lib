package pulse

import (
	"testing"
)

func TestQueryInfo(t *testing.T) {
	ctx := GetContext()
	if ctx == nil {
		t.Skip("Can't connect to pulseaudio.")
		return
	}
	if len(ctx.GetCardList()) == 0 {
		t.Fatal("Can't query sound card")
	}
}
