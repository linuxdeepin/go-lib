package sound_effect

import (
	"unsafe"

	paSimple "github.com/linuxdeepin/go-lib/pulse/simple"
)

type PulseAudioPlayBackend struct {
	conn paSimple.Conn
}

func newPulseAudioPlayBackend(event, device string, sampleSpec *SampleSpec) (PlayBackend, error) {
	paConn, err := paSimple.NewConn("", "com.deepin.SoundEffect",
		paSimple.StreamDirectionPlayback, device, event, sampleSpec.GetPaSampleSpec())

	if err != nil {
		return nil, err
	}

	return &PulseAudioPlayBackend{
		conn: paConn,
	}, nil
}

func (pb *PulseAudioPlayBackend) Write(data []byte) error {
	_, err := pb.conn.Write(unsafe.Pointer(&data[0]), uint(len(data)))
	return err
}

func (pb *PulseAudioPlayBackend) Drain() error {
	_, err := pb.conn.Drain()
	return err
}

func (pb *PulseAudioPlayBackend) Close() error {
	pb.conn.Free()
	return nil
}
