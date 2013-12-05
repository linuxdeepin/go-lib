package pseudo

import "dlib/dbus"
import "dlib/dbus/property"
import "dlib"
import "dlib/gobject-2.0"
import "dlib/gio-2.0"
import "dlib/glib-2.0"

func nothing() {
	_ = property.NewGSettingsProperty
	_ = gio.DBusConnectionFlagsAuthenticationClient
	_ = glib.CanInline
	_ = gobject.NilString
	_ = dlib.StartLoop
	_ = dbus.NewConn
}
