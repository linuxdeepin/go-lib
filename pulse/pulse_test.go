package pulse

import (
	fmtp "github.com/kr/pretty"
	"testing"
	"time"
)

func TestIntrospect(t *testing.T) {
	_ = fmtp.Print
	ctx := GetContext()
	//sink := ctx.GetSink(1)
	//sink.SetAvgVolume(1)
	////sink.SetBalance(0)
	//fmtp.Println(sink.Volume.Avg())

	//sink.SetMute(false)

	for _, si := range ctx.GetSinkInputList() {
		name := si.PropList["application.name"]
		if name == "mplayer2" {
			si.SetMute(true)
		}
	}

	<-time.After(time.Second)
}
