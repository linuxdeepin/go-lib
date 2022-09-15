// +build ignore

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"log"
	"time"

	"github.com/linuxdeepin/go-lib/notify"
)

func init() {
	notify.Init("notify-example-action")
}

func show() {
	n := notify.NewNotification("summary", "body", "icon")
	n.Timeout = notify.ExpiresSecond * 5
	n.AddAction("x", "XXX", func(_n *notify.Notification, action string) {
		log.Println("action", action, "invoked")
		_n.Summary = n.Summary + "!"
		_n.Show()
	})

	n.AddAction("close", "Close", func(_n *notify.Notification, action string) {
		log.Println("close it")
	})

	n.Closed().On(func(_n *notify.Notification, reason notify.ClosedReason) {
		log.Printf("reason: %d %s\n", reason, reason)
	})
	n.Show()
}

func main() {
	go show()
	time.Sleep(time.Second * 100)
}
