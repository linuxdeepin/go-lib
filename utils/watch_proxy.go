/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package utils

import "github.com/howeyc/fsnotify"

type WatchProxy struct {
	watcher      *fsnotify.Watcher
	eventHandler func(*fsnotify.FileEvent)
	errorHandler func(error)
	fileList     []string
	end          chan bool
	endFlag      bool
}

func NewWatchProxy() *WatchProxy {
	w := &WatchProxy{}

	var err error
	w.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil
	}

	w.end = make(chan bool, 1)
	w.endFlag = true

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
	w.endFlag = false

	if len(w.fileList) == 0 || w.watcher == nil {
		w.endFlag = true
		return
	}

	w.setFileListWatch()

	for {
		select {
		case ev, ok := <-w.watcher.Event:
			if !ok {
				w.ResetFileListWatch()
			}
			w.eventHandler(ev)
		case err, ok := <-w.watcher.Error:
			if !ok {
				w.ResetFileListWatch()
			}
			if w.errorHandler != nil {
				w.errorHandler(err)
			}
		case <-w.end:
			w.watcher.Close()
			return
		}
	}
}

func (w *WatchProxy) EndWatch() {
	if !w.endFlag {
		w.endFlag = true
		w.removeFileListWatch()
		w.end <- true
	}
}

/**
 * Todo: func (w *WatchProxy) RestartWatch() {}
 **/
