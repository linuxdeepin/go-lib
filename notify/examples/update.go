// +build ignore

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
