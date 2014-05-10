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

func TestDefault(t *testing.T) {
	fmt.Println(ctx.GetServer())
}

func TestSinkInput(t *testing.T) {
	defer drain()

	for _, si := range ctx.GetSinkInputList() {
		name := si.PropList["application.name"]
		if name == "mplayer2" {
			si.SetMute(true)
		}
	}
	ctx.GetSinkInput(0)
}

func TestEvent(t *testing.T) {
	ctx.Connect(FacilitySinkInput, func(eType int, idx uint32) {
		fmt.Println("SinkInput Changed...", eType, ctx.GetSinkInput(idx).GetAvgVolume())
	})
	fmt.Println("HEHE...")
}

func TestIntrospect(t *testing.T) {
	_ = fmtp.Print
	sink := ctx.GetSink(1)
	fmt.Println(sink.Description, "Volume:", sink.Volume.Avg())
	sink.SetAvgVolume(0.1)
	//sink.SetBalance(-1)
	fmtp.Println(sink.Volume.Avg())

	sink.SetMute(false)

}
