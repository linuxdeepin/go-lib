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

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/linuxdeepin/go-lib/asound"
	paSimple "github.com/linuxdeepin/go-lib/pulse/simple"
	"github.com/linuxdeepin/go-lib/sound_effect/theme"
)

type Player struct {
	finder *theme.Finder

	cache       map[string]*CacheItem
	cacheMu     sync.Mutex
	UseCache    bool
	backendType PlayBackendType
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
	filename := player.finder.Find(theme, "stereo", event)
	if filename == "" {
		return errors.New("not found file")
	}

	return player.play(filename, event, device)
}

func cacheItemOk(cacheItem *CacheItem, fileInfo os.FileInfo) bool {
	return cacheItem.modTime == fileInfo.ModTime()
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
