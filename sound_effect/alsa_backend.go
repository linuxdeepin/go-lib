// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sound_effect

import (
	"unsafe"

	"github.com/linuxdeepin/go-lib/asound"
)

type ALSAPlayBackend struct {
	pcm       asound.PCM
	frameSize int
}

func newALSAPlayBackend(device string, sampleSpec *SampleSpec) (pb PlayBackend, err error) {
	if device == "" {
		device = "default"
	}

	pcm, err := asound.OpenPCM(device, asound.PCMStreamPlayback, 0)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			//println("call pcm.Close()")
			pcm.Close()
		}
	}()

	params, err := asound.NewPCMHwParams()
	if err != nil {
		return nil, err
	}
	defer params.Free()

	// fill it in with default values
	pcm.HwParamsAny(params)

	channels := sampleSpec.channels
	format := sampleSpec.pcmFormat
	err = pcm.HwParamsSetAccess(params, asound.PCMAccessRWInterleaved)
	if err != nil {
		return
	}

	err = pcm.HwParamsSetFormat(params, format)
	if err != nil {
		return
	}

	err = pcm.HwParamsSetChannels(params, uint(channels))
	if err != nil {
		return
	}

	sampleRate := uint(sampleSpec.rate)
	_, err = pcm.HwParamsSetRateNear(params, &sampleRate)
	if err != nil {
		return
	}

	bufferTime, _, err := params.GetBufferTimeMax()
	if err != nil {
		return
	}

	if bufferTime > 500000 {
		bufferTime = 500000
	}

	periodTime := uint(0)
	if bufferTime > 0 {
		periodTime = bufferTime / 4
	}

	// set period time
	if periodTime > 0 {
		_, err = pcm.HwParamsSetPeriodTimeNear(params, &periodTime)
		if err != nil {
			return
		}
	}

	// set buffer time
	if bufferTime > 0 {
		_, err = pcm.HwParamsSetBufferTimeNear(params, &bufferTime)
		if err != nil {
			return
		}
	}

	err = pcm.HwParams(params)
	if err != nil {
		return
	}

	err = pcm.Prepare()
	if err != nil {
		return
	}

	frameSize := channels * sampleSpec.pcmFormat.Size(1)
	return &ALSAPlayBackend{
		pcm:       pcm,
		frameSize: frameSize,
	}, nil
}

func (pb *ALSAPlayBackend) Write(data []byte) error {
	frames := len(data) / pb.frameSize
	_, err := pb.pcm.Writei(unsafe.Pointer(&data[0]), asound.PCMUFrames(frames))
	if err == asound.ErrUnderrun {
		return pb.pcm.Prepare()
	}
	return err
}

func (pb *ALSAPlayBackend) Drain() error {
	return pb.pcm.Drain()
}

func (pb *ALSAPlayBackend) Close() error {
	return pb.pcm.Close()
}
