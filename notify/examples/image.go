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
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"pkg.deepin.io/lib/notify"
)

func init() {
	notify.Init("notify-example-image")
}

func loadImage(filename string) (image.Image, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func show(filename string) {
	n := notify.NewNotification("image test", "file "+filename, "")
	img, err := loadImage(filename)
	if err != nil {
		log.Fatal(err)
	}
	n.SetImage(img)
	n.Show()
	n.Destroy()
}

func main() {
	// opaque
	show("./desktop-icon.jpeg")
	// not opaque
	show("./libreoffice-base.png")

	notify.Destroy()
}
