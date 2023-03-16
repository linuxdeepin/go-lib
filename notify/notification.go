// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package notify

import (
	"errors"
	"runtime"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/event"
)

type ActionCallback func(n *Notification, action string)
type _HintMap map[string]dbus.Variant

type Notification struct {
	id           uint32
	AppName      string
	Summary      string
	Body         string
	Icon         string
	Timeout      int32
	closedReason ClosedReason
	mu           sync.Mutex

	actions   []string
	actionMap map[string]ActionCallback
	hints     _HintMap
	// event closed (n *Notification, reason ClosedReason)
	closed *event.Event

	actionInvokedConnectRet      func()
	notificationClosedConnectRet func()
}

// Create a new Notification.
// summary text is required, but all other paramters are optional.
func NewNotification(summary, body, icon string) *Notification {
	n := &Notification{
		Summary:      summary,
		Body:         body,
		Icon:         icon,
		Timeout:      ExpiresDefault,
		hints:        make(_HintMap),
		closedReason: ClosedReasonInvalid,
		actionMap:    make(map[string]ActionCallback),
		closed:       event.New(func(_n *Notification, reason ClosedReason) {}),
	}

	runtime.SetFinalizer(n, func(_n *Notification) { _n.Destroy() })
	return n
}

func (n *Notification) Closed() *event.Event {
	return n.closed
}

func (n *Notification) Destroy() {
	if n.actionInvokedConnectRet != nil {
		n.actionInvokedConnectRet()
		n.actionInvokedConnectRet = nil
	}

	if n.notificationClosedConnectRet != nil {
		n.notificationClosedConnectRet()
		n.notificationClosedConnectRet = nil
	}

	runtime.SetFinalizer(n, nil)
}

// Updates the notification text and icon.
// This won't send the update out and display it on the screen.
// For that, you will need to call Show()
func (n *Notification) Update(summary, body, icon string) {
	n.Summary = summary
	n.Body = body
	n.Icon = icon
}

func (n *Notification) getAppName() string {
	if n.AppName == "" {
		return defaultAppName
	}
	return n.AppName
}

// Tells the notification server to display the notification on the screen.
func (n *Notification) Show() error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if n.actionInvokedConnectRet == nil {
		n.actionInvokedConnectRet = notifier.ConnectActionInvoked(func(id uint32, action string) {
			if n.id != id {
				return
			}

			callback := n.actionMap[action]
			if callback != nil {
				callback(n, action)
			}
		})
	}

	if n.notificationClosedConnectRet == nil {
		n.notificationClosedConnectRet = notifier.ConnectNotificationClosed(func(id, reason uint32) {
			if n.id != id {
				return
			}
			n.closedReason = ClosedReason(reason)
			_ = n.closed.Trigger(n, ClosedReason(reason))
		})
	}

	id, err := notifier.Notify(n.getAppName(), n.id, n.Icon, n.Summary, n.Body, n.actions, n.hints, n.Timeout)
	if err != nil {
		return err
	}
	n.id = id
	return nil
}

func (n *Notification) SetHint(key string, value interface{}) {
	n.hints[key] = dbus.MakeVariant(value)
}

func (n *Notification) ClearHints() {
	n.hints = make(_HintMap)
}

func (n *Notification) AddAction(action, label string, callback ActionCallback) {
	if action == "" || label == "" || callback == nil {
		return
	}

	n.actions = append(n.actions, action, label)
	n.actionMap[action] = callback
}

func (n *Notification) ClearActions() {
	n.actions = nil
	n.actionMap = make(map[string]ActionCallback)
}

// Sets the category of this notification.
// This can be used by the notification server to filter or display the data in a certain way.
func (n *Notification) SetCategory(category string) {
	n.SetHint(HintCategory, category)
}

// Sets the urgency level of this notification.
func (n *Notification) SetUrgency(urgency byte) {
	n.SetHint(HintUrgency, urgency)
}

func (n *Notification) Close() error {
	if n.id == 0 {
		return errors.New("notification id is 0")
	}
	return notifier.CloseNotification(n.id)
}

func (n *Notification) GetClosedReason() ClosedReason {
	return n.closedReason
}
