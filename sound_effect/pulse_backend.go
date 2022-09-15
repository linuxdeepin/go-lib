// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sound_effect

import (
	"fmt"
	"unsafe"

	"github.com/linuxdeepin/go-lib/pulse"
	paSimple "github.com/linuxdeepin/go-lib/pulse/simple"
)

type PulseAudioPlayBackend struct {
	conn paSimple.Conn
}

func getPulseDefaultSink() string {
	ctx := pulse.GetContext()
	if ctx == nil {
		return ""
	}
	device := ctx.GetDefaultSink()
	fmt.Printf("use '%s' instead of empty device\n", device)
	return device
}

func newPulseAudioPlayBackend(event, device string, sampleSpec *SampleSpec) (PlayBackend, error) {
	if device == "" {
		device = getPulseDefaultSink()
	}

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
