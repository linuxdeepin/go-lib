package pulse

/*
#include "dde-pulse.h"
*/
import "C"
import "unsafe"

const (
	PA_PROP_MEDIA_NAME                     = "media.name"
	PA_PROP_MEDIA_TITLE                    = "media.title"
	PA_PROP_MEDIA_ARTIST                   = "media.artist"
	PA_PROP_MEDIA_COPYRIGHT                = "media.copyright"
	PA_PROP_MEDIA_SOFTWARE                 = "media.software"
	PA_PROP_MEDIA_LANGUAGE                 = "media.language"
	PA_PROP_MEDIA_FILENAME                 = "media.filename"
	PA_PROP_MEDIA_ICON_NAME                = "media.icon_name"
	PA_PROP_MEDIA_ROLE                     = "media.role"
	PA_PROP_FILTER_WANT                    = "filter.want"
	PA_PROP_FILTER_APPLY                   = "filter.apply"
	PA_PROP_FILTER_SUPPRESS                = "filter.suppress"
	PA_PROP_EVENT_ID                       = "event.id"
	PA_PROP_EVENT_DESCRIPTION              = "event.description"
	PA_PROP_EVENT_MOUSE_X                  = "event.mouse.x"
	PA_PROP_EVENT_MOUSE_Y                  = "event.mouse.y"
	PA_PROP_EVENT_MOUSE_HPOS               = "event.mouse.hpos"
	PA_PROP_EVENT_MOUSE_VPOS               = "event.mouse.vpos"
	PA_PROP_EVENT_MOUSE_BUTTON             = "event.mouse.button"
	PA_PROP_WINDOW_NAME                    = "window.name"
	PA_PROP_WINDOW_ID                      = "window.id"
	PA_PROP_WINDOW_ICON_NAME               = "window.icon_name"
	PA_PROP_WINDOW_X                       = "window.x"
	PA_PROP_WINDOW_Y                       = "window.y"
	PA_PROP_WINDOW_WIDTH                   = "window.width"
	PA_PROP_WINDOW_HEIGHT                  = "window.height"
	PA_PROP_WINDOW_HPOS                    = "window.hpos"
	PA_PROP_WINDOW_VPOS                    = "window.vpos"
	PA_PROP_WINDOW_DESKTOP                 = "window.desktop"
	PA_PROP_WINDOW_X11_DISPLAY             = "window.x11.display"
	PA_PROP_WINDOW_X11_SCREEN              = "window.x11.screen"
	PA_PROP_WINDOW_X11_MONITOR             = "window.x11.monitor"
	PA_PROP_WINDOW_X11_XID                 = "window.x11.xid"
	PA_PROP_APPLICATION_NAME               = "application.name"
	PA_PROP_APPLICATION_ID                 = "application.id"
	PA_PROP_APPLICATION_VERSION            = "application.version"
	PA_PROP_APPLICATION_ICON_NAME          = "application.icon_name"
	PA_PROP_APPLICATION_LANGUAGE           = "application.language"
	PA_PROP_APPLICATION_PROCESS_ID         = "application.process.id"
	PA_PROP_APPLICATION_PROCESS_BINARY     = "application.process.binary"
	PA_PROP_APPLICATION_PROCESS_USER       = "application.process.user"
	PA_PROP_APPLICATION_PROCESS_HOST       = "application.process.host"
	PA_PROP_APPLICATION_PROCESS_MACHINE_ID = "application.process.machine_id"
	PA_PROP_APPLICATION_PROCESS_SESSION_ID = "application.process.session_id"
	PA_PROP_DEVICE_STRING                  = "device.string"
	PA_PROP_DEVICE_API                     = "device.api"
	PA_PROP_DEVICE_DESCRIPTION             = "device.description"
	PA_PROP_DEVICE_BUS_PATH                = "device.bus_path"
	PA_PROP_DEVICE_SERIAL                  = "device.serial"
	PA_PROP_DEVICE_VENDOR_ID               = "device.vendor.id"
	PA_PROP_DEVICE_VENDOR_NAME             = "device.vendor.name"
	PA_PROP_DEVICE_PRODUCT_ID              = "device.product.id"
	PA_PROP_DEVICE_PRODUCT_NAME            = "device.product.name"
	PA_PROP_DEVICE_CLASS                   = "device.class"
	PA_PROP_DEVICE_FORM_FACTOR             = "device.form_factor"
	PA_PROP_DEVICE_BUS                     = "device.bus"
	PA_PROP_DEVICE_ICON_NAME               = "device.icon_name"
	PA_PROP_DEVICE_ACCESS_MODE             = "device.access_mode"
	PA_PROP_DEVICE_MASTER_DEVICE           = "device.master_device"
	PA_PROP_DEVICE_BUFFERING_BUFFER_SIZE   = "device.buffering.buffer_size"
	PA_PROP_DEVICE_BUFFERING_FRAGMENT_SIZE = "device.buffering.fragment_size"
	PA_PROP_DEVICE_PROFILE_NAME            = "device.profile.name"
	PA_PROP_DEVICE_INTENDED_ROLES          = "device.intended_roles"
	PA_PROP_DEVICE_PROFILE_DESCRIPTION     = "device.profile.description"
	PA_PROP_MODULE_AUTHOR                  = "module.author"
	PA_PROP_MODULE_DESCRIPTION             = "module.description"
	PA_PROP_MODULE_USAGE                   = "module.usage"
	PA_PROP_MODULE_VERSION                 = "module.version"
	PA_PROP_FORMAT_SAMPLE_FORMAT           = "format.sample_format"
	PA_PROP_FORMAT_RATE                    = "format.rate"
	PA_PROP_FORMAT_CHANNELS                = "format.channels"
	PA_PROP_FORMAT_CHANNEL_MAP             = "format.channel_map"
)

type PropList map[string]string

func (p PropList) toC() *C.pa_proplist {
	core := C.pa_proplist_new()
	for k, v := range p {
		ck := C.CString(k)
		cv := C.CString(v)
		C.pa_proplist_sets(core, ck, cv)
		C.free(unsafe.Pointer(ck))
		C.free(unsafe.Pointer(cv))
	}
	return core
}
