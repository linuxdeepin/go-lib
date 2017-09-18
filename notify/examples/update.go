// +build ignore

/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
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
	"pkg.deepin.io/lib/notify"
	"time"
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
		log.Fatal(err)
	}
	log.Println("server caps:", caps)

	n := notify.NewNotification("summary", "body", "deepin-music")
	n.Show()

	time.Sleep(time.Second * 2)
	n.Update("xxxx", "yyyy", "deepin-terminal")
	n.Show()
}
