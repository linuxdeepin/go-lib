package pseudo

import "dlib/dbus"
import "dlib/dbus/property"
import "dlib"
import "dlib/gobject-2.0"
import "dlib/gio-2.0"
import "dlib/glib-2.0"
import "dlib/logger"
import "dlib/graphic"
import "dlib/pulse"
import "dlib/utils"

func nothing() {
	_ = graphic.GetDominantColorOfImage
	_ = logger.NewLogger
	_ = property.NewGSettingsBoolProperty
	_ = gio.DBusConnectionFlagsAuthenticationClient
	_ = glib.CanInline
	_ = gobject.NilString
	_ = dlib.SessionBus
	_ = dbus.NewConn
	_ = utils.NewUtils
	_ = pulse.GetContext
}
