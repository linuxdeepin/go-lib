// +build ignore

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/linuxdeepin/go-lib/notify"
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
