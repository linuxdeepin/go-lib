// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sound_effect

import (
	"io"
	"time"

	"github.com/linuxdeepin/go-lib/asound"
	paSimple "github.com/linuxdeepin/go-lib/pulse/simple"
	"github.com/linuxdeepin/go-lib/stb_vorbis"
)

func newOggDecoder(file string) (Decoder, error) {
	decoder, err := stb_vorbis.OpenFile(file)
	if err != nil {
		return nil, err
	}
	info := decoder.GetInfo()
	channels := info.Channels
	sampleRate := info.SampleRate
	sampleSpec := &SampleSpec{
		channels:  channels,
		rate:      int(sampleRate),
		paFormat:  paSimple.SampleFormatS16LE,
		pcmFormat: asound.PCMFormatS16LE,
	}
	bufSize := int(sampleRate) / 8 * channels * 2
	return &OggDecoder{
		core:       decoder,
		sampleSpec: sampleSpec,
		bufSize:    bufSize,
	}, nil
}

type OggDecoder struct {
	core       stb_vorbis.Decoder
	sampleSpec *SampleSpec
	bufSize    int
}

func (d *OggDecoder) GetDuration() time.Duration {
	// TODO
	return 0
}

func (d *OggDecoder) Decode() ([]byte, error) {
	channels := d.sampleSpec.channels
	buf := make([]byte, d.bufSize)
	samples := d.core.GetSamplesShortInterleaved(channels, buf)
	if samples <= 0 {
		err := d.core.GetError()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	}

	bytes := uint(2 * samples * channels)
	return buf[:bytes], nil
}

func (d *OggDecoder) GetSampleSpec() *SampleSpec {
	return d.sampleSpec
}

func (d *OggDecoder) Close() error {
	d.core.Close()
	return nil
}
