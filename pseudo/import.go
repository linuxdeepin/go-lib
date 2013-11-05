package pseudo

import "dlib/dbus"
import "dlib"

func nothing() {
	_ = dlib.NewSettings
	_ = dbus.NewConn
}
