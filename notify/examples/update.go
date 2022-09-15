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
	notify.Init("notify-example-update")
}

func main() {
	serverInfo, err := notify.GetServerInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("serverInfo: %#v\n", serverInfo)
	caps, err := notify.GetServerCaps()
	if err != nil {
		log.Println(err)
	}
	log.Println("server caps:", caps)

	n := notify.NewNotification("s1", "b1", "deepin-music")
	n.Show()

	time.Sleep(time.Second * 2)
	n.Update("s2", "b2", "deepin-terminal")
	n.Show()

	time.Sleep(time.Second * 2)
	n.Update("s3", "b3", "deepin-terminal")
	n.Show()
}
