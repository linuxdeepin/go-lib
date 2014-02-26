package pseudo

import "dlib/dbus"
import "dlib/dbus/property"
import "dlib"
import "dlib/gobject-2.0"
import "dlib/gio-2.0"
import "dlib/glib-2.0"
import "dlib/logger"
import "dlib/graph"

func nothing() {
	_ = graph.ClipPNG
	_ = logger.Println
	_ = property.NewGSettingsBoolProperty
	_ = gio.DBusConnectionFlagsAuthenticationClient
	_ = glib.CanInline
	_ = gobject.NilString
	_ = dlib.StartLoop
	_ = dbus.NewConn
}
