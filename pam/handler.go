// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pam

import "sync"

// ConversationHandler is an interface for objects that can be used as
// conversation callbacks during PAM authentication.
type ConversationHandler interface {
	// RespondPAM receives a message style and a message string. If the
	// message Style is PromptEchoOff or PromptEchoOn then the function
	// should return a response string.
	RespondPAM(Style, string) (string, error)
}

// ConversationFunc is an adapter to allow the use of ordinary functions as
// conversation callbacks.
type ConversationFunc func(Style, string) (string, error)

func (f ConversationFunc) RespondPAM(s Style, msg string) (string, error) {
	return f(s, msg)
}

var globalHandlers struct {
	sync.Mutex
	m      map[handlerId]ConversationHandler
	nextId handlerId
}

type handlerId int

func init() {
	globalHandlers.m = make(map[handlerId]ConversationHandler)
	globalHandlers.nextId = 1
}

func addHandler(handler ConversationHandler) handlerId {
	globalHandlers.Lock()
	defer globalHandlers.Unlock()
	id := globalHandlers.nextId
	globalHandlers.nextId++
	globalHandlers.m[id] = handler
	return id
}

func getHandler(id handlerId) ConversationHandler {
	globalHandlers.Lock()
	defer globalHandlers.Unlock()
	v := globalHandlers.m[id]
	if v != nil {
		return v
	}
	panic("handler not found")
}

func deleteHandler(id handlerId) {
	globalHandlers.Lock()
	defer globalHandlers.Unlock()
	if _, ok := globalHandlers.m[id]; !ok {
		panic("handler not found")
	}
	delete(globalHandlers.m, id)
}
