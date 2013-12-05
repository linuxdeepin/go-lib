dlib/dbus
-------

The core dbus communication code is mainly use [go.dbus](http://github.com/guelfey/go.dbus) with little modify.

This library is mainly create an easy frame-less dbus golang interface.

### Features
* Auto export struct in golang to dbus,
* implement an dbus.Property interface to auto manager property's set/get/notify
* an dbus-binding code generator that can generate dbus-binding for QML/PyQt/Golang.
