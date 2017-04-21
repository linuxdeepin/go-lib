// +build ignore

package main

import (
	"pkg.deepin.io/lib/notify"
)

func init() {
	notify.Init("notify-example-simple")
}

func main() {
	n := notify.NewNotification("x", "y", "player")
	n.Show()
	n.Destroy()
	notify.Destroy()
}
