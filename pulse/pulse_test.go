package pulse

import (
	"fmt"
	fmtp "github.com/kr/pretty"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var ctx = GetContext()

func drain() {
	ctx.GetSink(0)
}

func TestDefault(t *testing.T) {
	return
	fmt.Println(ctx.GetServer())
}

func TestRepeat(t *testing.T) {
	fmt.Println("Begin...")
	for i := 0; i < 100000; i++ {
		s := ctx.GetSink(1)
		v := s.Volume.SetAvg(float64(rand.Int31n(100)) / 100.0)
		<-time.After(time.Microsecond * 5)
		//fmt.Println(v)
		s.SetVolume(v)
	}
	fmt.Println("PASSS")
	runtime.Gosched()
}

func TestSinkInput(t *testing.T) {
	return
	defer drain()

	for _, si := range ctx.GetSinkInputList() {
		name := si.PropList["application.name"]
		if name == "mplayer2" {
			//si.SetMute(true)
		}
	}
	ctx.GetSinkInput(0)
}

func TestEvent(t *testing.T) {
	return
	ctx.Connect(FacilitySinkInput, func(eType int, idx uint32) {
		fmt.Println("SinkInput Changed...", eType, ctx.GetSinkInput(idx).Volume.Avg())
	})
}

func TestIntrospect(t *testing.T) {
	return
	_ = fmtp.Print
	sink := ctx.GetSink(1)
	sink.SetVolume(sink.Volume.SetAvg(1))
	sink.SetVolume(sink.Volume.SetBalance(sink.ChannelMap, -0.8))
	fmt.Println(sink.Description, "Volume:", sink.Volume.Avg())
	fmtp.Println(sink.Volume.Avg())

	sink.SetMute(false)

}
