/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package sound_effect

// #include "wav.h"
//#cgo LDFLAGS: -lm
import "C"

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"pkg.deepin.io/lib/asound"
	paSimple "pkg.deepin.io/lib/pulse/simple"
	"pkg.deepin.io/lib/sound_effect/theme"
)

type Player struct {
	finder *theme.Finder

	cache       map[string]*CacheItem
	cacheMu     sync.Mutex
	UseCache    bool
	backendType PlayBackendType
	Volume      float32
}

type CacheItem struct {
	modTime    time.Time
	event      string
	sampleSpec *SampleSpec
	data       [][]byte
}

type Decoder interface {
	GetSampleSpec() *SampleSpec
	Decode() ([]byte, error)
	Close() error
	GetDuration() time.Duration
}

type SampleSpec struct {
	channels  int
	rate      int
	paFormat  paSimple.SampleFormat
	pcmFormat asound.PCMFormat
}

func (ss *SampleSpec) GetPaSampleSpec() *paSimple.SampleSpec {
	return &paSimple.SampleSpec{
		Channels: uint8(ss.channels),
		Rate:     uint32(ss.rate),
		Format:   ss.paFormat,
	}
}

func NewPlayer(useCache bool, backendType PlayBackendType) *Player {
	player := &Player{
		finder:      theme.NewFinder(),
		UseCache:    useCache,
		backendType: backendType,
	}
	if useCache {
		player.cache = make(map[string]*CacheItem)
	}
	return player
}

func (player *Player) Finder() *theme.Finder {
	return player.finder
}

func (player *Player) GetDuration(theme, event string) (time.Duration, error) {
	filename := player.finder.Find(theme, "stereo", event)
	if filename == "" {
		return 0, errors.New("not found file")
	}
	return player.getDuration(filename)
}

func (player *Player) Play(theme, event, device string) error {
	speakerSwitch := false
	filename := player.finder.Find(theme, "stereo", event)
	if filename == "" {
		return errors.New("not found file")
	}
	if event == "desktop-login" && player.backendType == PlayBackendALSA {
		speakerSwitch = true
		os.Remove("/tmp/desktop-login.wav")
		C.wav_convert(C.CString(filename), C.CString("/tmp/desktop-login.wav"), C.float(player.Volume/100.0))
		filename = "/tmp/desktop-login.wav"
	} else if event == "system-shutdown" && player.backendType == PlayBackendALSA {
		speakerSwitch = true
		os.Remove("/tmp/system-shutdown.wav")
		C.wav_convert(C.CString(filename), C.CString("/tmp/system-shutdown.wav"), C.float(player.Volume/100.0))
		filename = "/tmp/system-shutdown.wav"
	} else if event == "desktop-logout" && player.backendType == PlayBackendALSA {
		speakerSwitch = true
		os.Remove("/tmp/desktop-logout.wav")
		C.wav_convert(C.CString(filename), C.CString("/tmp/desktop-logout.wav"), C.float(player.Volume/100.0))
		filename = "/tmp/desktop-logout.wav"
	}

	var err error
	if speakerSwitch {
		hwSpeakerEnable(true)
		err = player.play(filename, event, device)
		hwSpeakerEnable(false)
	} else {
		err = player.play(filename, event, device)
	}

	return err
}

func cacheItemOk(cacheItem *CacheItem, fileInfo os.FileInfo) bool {
	if cacheItem.modTime != fileInfo.ModTime() {
		return false
	}
	return true
}

func getDecoder(file string, fileInfo os.FileInfo) (Decoder, error) {
	ext := filepath.Ext(file)
	switch ext {
	case ".ogg", ".oga":
		return newOggDecoder(file)
	case ".wav":
		return newWavDecoder(file, fileInfo)
	default:
		return nil, fmt.Errorf("unsupported ext %q", ext)
	}
}

type PlayBackendType uint

const (
	PlayBackendPulseAudio = iota
	PlayBackendALSA
)

type PlayBackend interface {
	Write(data []byte) error
	Drain() error
	Close() error
}

func getPlayBackend(type0 PlayBackendType, event, device string, sampleSpec *SampleSpec) (PlayBackend, error) {
	switch type0 {
	case PlayBackendPulseAudio:
		return newPulseAudioPlayBackend(event, device, sampleSpec)
	case PlayBackendALSA:
		return newALSAPlayBackend(device, sampleSpec)
	default:
		return nil, fmt.Errorf("unknown play backend type %d", type0)
	}
}

func (Player *Player) getDuration(file string) (time.Duration, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	decoder, err := getDecoder(file, fileInfo)
	if err != nil {
		return 0, err
	}
	defer decoder.Close()
	return decoder.GetDuration(), nil
}

func (player *Player) play(file, event, device string) error {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return err
	}

	var doCache = player.UseCache
	if fileInfo.Size() > 30*1024 {
		doCache = false
	}

	if doCache {
		player.cacheMu.Lock()
		cacheItem, ok := player.cache[file]
		player.cacheMu.Unlock()

		if ok {
			cacheOk := cacheItemOk(cacheItem, fileInfo)
			if cacheOk {
				return player.playCacheItem(cacheItem, device)
			} else {
				player.cacheMu.Lock()
				delete(player.cache, file)
				player.cacheMu.Unlock()
			}
		}
	}

	decoder, err := getDecoder(file, fileInfo)
	if err != nil {
		return err
	}
	defer decoder.Close()

	sampleSpec := decoder.GetSampleSpec()
	backend, err := getPlayBackend(player.backendType, event, device, sampleSpec)
	if err != nil {
		return err
	}
	defer backend.Close()

	var cacheData [][]byte
	for {
		data, err := decoder.Decode()
		if len(data) > 0 {
			if doCache {
				cacheData = append(cacheData, data)
			}

			err := backend.Write(data)
			if err != nil {
				return err
			}
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	if doCache {
		cacheItem := &CacheItem{
			event:      event,
			modTime:    fileInfo.ModTime(),
			sampleSpec: sampleSpec,
			data:       cacheData,
		}

		player.cacheMu.Lock()
		player.cache[file] = cacheItem
		player.cacheMu.Unlock()
	}

	return backend.Drain()
}

func (player *Player) playCacheItem(cacheItem *CacheItem, device string) error {
	backend, err := getPlayBackend(player.backendType, cacheItem.event, device, cacheItem.sampleSpec)
	if err != nil {
		return err
	}
	defer backend.Close()
	for _, data := range cacheItem.data {
		err := backend.Write(data)
		if err != nil {
			return err
		}
	}

	return backend.Drain()
}

func hwSpeakerEnable(s bool) error {
	param := ""
	if s {
		param = "-o"
	} else {
		param = "-c"
	}
	//如果可执行文件执行失败，开关音效会失效联系华为解决
	binPath := []string{
		"/usr/share/hw-audio/hwaudioservice", //TODO 82239 华为为了规避内核版本不同，需要我们上层做出规避,执行不同目录下的hwaudioservice 文件
		"/usr/bin/hwaudioservice",
	}

	for _, binFile := range binPath {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		err := syscall.Access(binFile, syscall.F_OK)
		if !os.IsNotExist(err) {
			cmd := exec.CommandContext(ctx, binFile, param)
			cmd.CombinedOutput()
		}
	}

	return nil
}
