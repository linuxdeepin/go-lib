// +build ignore

/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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

package main

import (
	"log"
	"github.com/linuxdeepin/go-lib/notify"
	"time"
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
