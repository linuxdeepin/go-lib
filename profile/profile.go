/*
 * Copyright (C) 2015 ~ 2017 Deepin Technology Co., Ltd.
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

// profiling of your Go application.
package profile

import (
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/pprof"
	"sync"

	"pkg.deepin.io/lib/initializer/v2"
)

type _Profile struct {
	writer io.Writer
	file   string
}

func (prof *_Profile) File() string {
	return prof.file
}

type _CPUProfile struct {
	_Profile
}

func newCPUProfile(file string) *_CPUProfile {
	return &_CPUProfile{_Profile{file: file}}
}

func (prof *_CPUProfile) Start(writer io.Writer) {
	pprof.StartCPUProfile(writer)
}

func (prof *_CPUProfile) Stop() {
	pprof.StopCPUProfile()
}

type _MemProfile struct {
	_Profile
}

func newMemPrifle(file string) *_MemProfile {
	return &_MemProfile{_Profile{file: file}}
}

func (prof *_MemProfile) Start(writer io.Writer) {
	prof.writer = writer
}

func (prof *_MemProfile) Stop() {
	pprof.Lookup("heap").WriteTo(prof.writer, 0)
}

type _BlockProfile struct {
	_Profile
}

func newBlockProfile(file string) *_BlockProfile {
	return &_BlockProfile{_Profile{file: file}}
}

func (prof *_BlockProfile) Start(writer io.Writer) {
	prof.writer = writer
}

func (prof *_BlockProfile) Stop() {
	pprof.Lookup("block").WriteTo(prof.writer, 0)
}

// Config controls the operation of the profile package.
type Config struct {
	// CPUProfile is the name of cpu profile which controls if cpu profiling will be enabled.
	// It defaults to false.
	CPUProfile string

	// MemProfile is the name of memory profile which controls if cpu profiling will be enabled.
	// It defaults to false.
	MemProfile string

	// MemProfile is the name of memory profile which controls if cpu profiling will be enabled.
	// It defaults to false.
	BlockProfile string

	// NoShutdownHook controls whether the profiling package should
	// hook SIGINT to write profiles cleanly.
	// Programs with more sophisticated signal handling should set
	// this to true and ensure the Stop() function returned from Start()
	// is called during shutdown.
	NoShutdownHook bool

	closers   []func()
	closeOnce sync.Once
}

func (cfg *Config) enableProfile(prof interface {
	File() string
	Start(io.Writer)
	Stop()
}) error {
	file := prof.File()

	// if name is empty, do not enable profile.
	if file == "" {
		return nil
	}

	var err error

	err = os.MkdirAll(filepath.Dir(file), 0777)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		cfg.stop()
		return err
	}

	prof.Start(f)
	cfg.closers = append(cfg.closers, func() {
		prof.Stop()
		f.Close()
	})

	return nil
}

// Start starts a new profiling session configured using *Config.
// The caller should call the Stop method to cleanly stop profiling.
func (cfg *Config) Start() error {
	if err := initializer.Do(func() error {
		return cfg.enableProfile(newCPUProfile(cfg.CPUProfile))
	}).Do(func() error {
		return cfg.enableProfile(newMemPrifle(cfg.MemProfile))
	}).Do(func() error {
		return cfg.enableProfile(newBlockProfile(cfg.BlockProfile))
	}).GetError(); err != nil {
		return err
	}

	if !cfg.NoShutdownHook {
		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			<-c

			cfg.stop()

			os.Exit(0)
		}()
	}
	return nil
}

func (cfg *Config) stop() {
	cfg.closeOnce.Do(func() {
		for _, c := range cfg.closers {
			c()
		}
	})
}

// Stop stops all profile.
func (cfg *Config) Stop() {
	if !cfg.NoShutdownHook {
		return
	}
	cfg.stop()

}
