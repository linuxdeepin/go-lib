// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gsettings

import (
	"path"
	"strings"
	"sync"

	"github.com/godbus/dbus/v5"
)

func addMatch(bus *dbus.Conn, rule string) error {
	err := bus.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, rule).Err
	return err
}

func initBus() (*dbus.Conn, error) {
	bus, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	err = addMatch(bus, "type='signal',sender='ca.desrt.dconf',path='/ca/desrt/dconf/Writer/user',interface='ca.desrt.dconf.Writer',member='Notify'")
	if err != nil {
		return nil, err
	}
	return bus, nil
}

type changedCallback struct {
	all bool
	fns []ChangedCallbackFunc
}

type ChangedCallbackFunc func(key string)

func toPath(schemaOrPath string) string {
	if schemaOrPath[0] == '/' {
		// is path
		if schemaOrPath[len(schemaOrPath)-1] == '/' {
			schemaOrPath = schemaOrPath[:len(schemaOrPath)-1]
		}
		return schemaOrPath
	}
	// is schema id
	return "/" + strings.Replace(schemaOrPath, ".", "/", -1)
}

var changedCallbackMap map[string]*changedCallback
var changedCallbackMapMu sync.RWMutex

func ConnectChanged(schemaOrPath, key string, fn ChangedCallbackFunc) {
	if schemaOrPath == "" || key == "" {
		return
	}
	var all bool
	if key == "*" {
		all = true
	}

	keyPath := toPath(schemaOrPath)
	if !all {
		keyPath += "/" + key
	}

	//log.Println("ConnectChanged", keyPath)

	changedCallbackMapMu.Lock()
	if changedCallbackMap == nil {
		changedCallbackMap = make(map[string]*changedCallback)
	}

	callback, ok := changedCallbackMap[keyPath]
	if ok {
		if callback.all == all {
			callback.fns = append(callback.fns, fn)
		}
	} else {
		changedCallbackMap[keyPath] = &changedCallback{
			all: all,
			fns: []ChangedCallbackFunc{fn},
		}
	}
	changedCallbackMapMu.Unlock()
}

var started bool
var startedMu sync.Mutex

func StartMonitor() error {
	startedMu.Lock()
	defer startedMu.Unlock()
	if started {
		return nil
	}

	bus, err := initBus()
	if err != nil {
		return err
	}
	signalCh := make(chan *dbus.Signal, 10)
	bus.Signal(signalCh)

	go func() {
		for signal := range signalCh {
			if signal.Name == "ca.desrt.dconf.Writer.Notify" &&
				signal.Path == "/ca/desrt/dconf/Writer/user" {
				if len(signal.Body) == 3 {
					keyPath, ok := signal.Body[0].(string)
					if !ok {
						continue
					}
					subPathList, ok := signal.Body[1].([]string)
					if !ok || len(subPathList) == 0 {
						handleSignal(keyPath)
					} else {
						for _, subPath := range subPathList {
							handleSignal(keyPath + subPath)
						}
					}
				}
			}
		}

	}()

	started = true
	return nil
}

func handleSignal(keyPath string) {
	parent, key := path.Split(keyPath)
	if parent == "" || key == "" {
		return
	}
	if parent[len(parent)-1] == '/' {
		parent = parent[:len(parent)-1]
	}

	changedCallbackMapMu.RLock()

	callback := changedCallbackMap[parent]
	if callback != nil && callback.all {
		for _, fn := range callback.fns {
			go fn(key)
		}
	}

	callback = changedCallbackMap[keyPath]
	if callback != nil && !callback.all {
		for _, fn := range callback.fns {
			go fn(key)
		}
	}

	changedCallbackMapMu.RUnlock()
}
