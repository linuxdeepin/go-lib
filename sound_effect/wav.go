// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sound_effect

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/linuxdeepin/go-lib/asound"
	paSimple "github.com/linuxdeepin/go-lib/pulse/simple"
	wav "github.com/youpy/go-wav"
)

type WavDecoder struct {
	reader        *wav.Reader
	f             *os.File
	sampleSpec    *SampleSpec
	bufSize       int
	bytesPerFrame int
	duration      time.Duration
}

func newWavDecoder(filename string) (Decoder, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	wavReader := wav.NewReader(f)
	format, err := wavReader.Format()
	if err != nil {
		return nil, err
	}
	paFormat, pcmFormat, err := getWavFormat(format)
	if err != nil {
		return nil, err
	}

	sampleSpec := &SampleSpec{
		paFormat:  paFormat,
		pcmFormat: pcmFormat,
		rate:      int(format.SampleRate),
		channels:  int(format.NumChannels),
	}

	bytesPerFrame := int(format.BlockAlign)
	bufSize := int(format.SampleRate/8) * bytesPerFrame
	duration, err := wavReader.Duration()
	if err != nil {
		return nil, err
	}

	return &WavDecoder{
		f:             f,
		reader:        wavReader,
		sampleSpec:    sampleSpec,
		bufSize:       bufSize,
		bytesPerFrame: bytesPerFrame,
		duration:      duration,
	}, nil
}

func getWavFormat(wavFormat *wav.WavFormat) (paSimple.SampleFormat, asound.PCMFormat, error) {
	if wavFormat.AudioFormat != wav.AudioFormatPCM {
		// 非 PCM 格式，暂时不支持
		return 0, 0, fmt.Errorf("wav audio format %v unsupported", wavFormat.AudioFormat)
	}
	switch wavFormat.BitsPerSample {
	case 8:
		return paSimple.SampleFormatU8, asound.PCMFormatU8, nil
	case 16:
		return paSimple.SampleFormatS16LE, asound.PCMFormatS16LE, nil
	case 24:
		return paSimple.SampleFormatS24LE, asound.PCMFormatS24_3LE, nil
	case 32:
		return paSimple.SampleFormatS32LE, asound.PCMFormatS32LE, nil
	default:
		return 0, 0, errors.New("unsupported format")
	}
}

func (d *WavDecoder) GetSampleSpec() *SampleSpec {
	return d.sampleSpec
}

func (d *WavDecoder) Decode() ([]byte, error) {
	buf := make([]byte, d.bufSize)
	n, err := d.read(buf)
	return buf[:n], err
}

func (d *WavDecoder) GetDuration() time.Duration {
	return d.duration
}

func (d *WavDecoder) read(buf []byte) (int, error) {
	return d.reader.Read(buf)
}

func (d *WavDecoder) Close() error {
	return d.f.Close()
}
