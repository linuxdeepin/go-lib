package pulse

import (
	"fmt"
	fmtp "github.com/kr/pretty"
	"testing"
)

var ctx = GetContext()

func drain() {
	ctx.GetSink(0)
}

func TestSinkInput(t *testing.T) {
	defer drain()

	for _, si := range ctx.GetSinkInputList() {
		name := si.PropList["application.name"]
		if name == "mplayer2" {
			si.SetMute(true)
		}
	}
}

func TestEvent(t *testing.T) {
	ctx.Connect(FacilitySinkInput, func(eType int, idx uint32) {
		fmt.Println("SinkInput Changed...", eType, ctx.GetSinkInput(idx).GetAvgVolume())
	})
	fmt.Println("HEHE...")
	select {}
}

func TestIntrospect(t *testing.T) {
	_ = fmtp.Print
	sink := ctx.GetSink(1)
	sink.SetAvgVolume(1)
	//sink.SetBalance(0)
	fmtp.Println(sink.Volume.Avg())

	sink.SetMute(false)

}
