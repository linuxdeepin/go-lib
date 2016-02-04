/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package utils

import (
	"github.com/howeyc/fsnotify"
	"sync"
)

type WatchProxy struct {
	watcher      *fsnotify.Watcher
	eventHandler func(*fsnotify.FileEvent)
	errorHandler func(error)
	fileList     []string
	end          chan bool
	endFlag      bool
	lock         sync.Mutex
}

func NewWatchProxy() *WatchProxy {
	w := &WatchProxy{}

	var err error
	w.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil
	}

	w.end = make(chan bool, 1)
	w.setEndFlag(true)

	return w
}

func (w *WatchProxy) setFileListWatch() {
	if w.watcher == nil {
		return
	}

	for _, filename := range w.fileList {
		w.watcher.Watch(filename)
	}
}

func (w *WatchProxy) removeFileListWatch() {
	if w.watcher == nil {
		return
	}

	for _, filename := range w.fileList {
		w.watcher.RemoveWatch(filename)
	}
}

func (w *WatchProxy) SetFileList(fileList []string) {
	w.fileList = fileList
}

func (w *WatchProxy) SetEventHandler(f func(*fsnotify.FileEvent)) {
	w.eventHandler = f
}

func (w *WatchProxy) SetErrorHandler(f func(error)) {
	w.errorHandler = f
}

func (w *WatchProxy) ResetFileListWatch() {
	w.removeFileListWatch()
	w.setFileListWatch()
}

func (w *WatchProxy) StartWatch() {
	if !w.endFlag || w.eventHandler == nil {
		return
	}
	w.setEndFlag(false)

	if len(w.fileList) == 0 || w.watcher == nil {
		w.setEndFlag(true)
		return
	}

	w.setFileListWatch()

	for {
		select {
		case ev, ok := <-w.watcher.Event:
			if !ok {
				w.ResetFileListWatch()
				break
			}
			if w.eventHandler == nil {
				break
			}
			w.eventHandler(ev)
		case err, ok := <-w.watcher.Error:
			if !ok {
				w.ResetFileListWatch()
				break
			}
			if w.errorHandler == nil {
				break
			}
			w.errorHandler(err)
		case <-w.end:
			if w.watcher == nil {
				return
			}
			w.watcher.Close()
			return
		}
	}
}

func (w *WatchProxy) EndWatch() {
	if !w.endFlag {
		w.setEndFlag(true)
		w.removeFileListWatch()
		w.end <- true
	}
}

func (w *WatchProxy) setEndFlag(value bool) {
	w.lock.Lock()
	w.endFlag = value
	w.lock.Unlock()
}

/**
 * Todo: func (w *WatchProxy) RestartWatch() {}
 **/
