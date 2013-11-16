package main

var __GLOBAL_TEMPLATE_PyQT = `#! /usr/bin/env python
# This file is auto generate by dlib/dbus/proxyer @linuxdeepin.com . Don't edit it

from PyQt5.QtCore import QObject, pyqtSlot{{range GetModules}}
from {{.}} import *{{end}}

class DBusFactory(QObject):
    def __init__(self, parent=None):
        super(DBusFactory, self).__init__(parent)
        self.__objects = {}
{{range .Interfaces}}
    @pyqtSlot(str, result=QObject)
    def get{{.ObjectName}}(self, path):
        if hasattr(self.__objects, path):
            return self.__objects[path]
        else:
            obj = {{.ObjectName}}(path)
            self.__objects[path]=obj
            return obj
{{end}}
`

var __IFC_TEMPLATE_INIT_PyQt = `#! /usr/bin/env python
# This file is auto generate by dlib/dbus/proxyer @linuxdeepin.com . Don't edit it
from PyQt5.QtCore import QObject, pyqtSlot, pyqtSignal, pyqtProperty, QVariant
from PyQt5.QtDBus import QDBusAbstractInterface, QDBusConnection, QDBusReply, QDBusMessage, QDBusInterface, QDBusError
`

var __IFC_TEMPLATE_PyQt = `
class {{ExportName}}(QObject):
    def connectSignal(self, signal):
        getattr({{ExportName}}.Proxyer, signal).connect(getattr(view.rootObject(), "on%s" % signal))
    class Proxyer(QDBusAbstractInterface):{{range .Signals}}
       {{.Name}} = pyqtSignal(QDBusMessage)
{{end}}
       def __init__(self, bus, path, parent=None):
           super({{ExportName}}.Proxyer, self).__init__("{{DestName}}", path, "{{IfcName}}", bus, parent)



    def __init__(self, path, parent=None):
        self.path = path
        super({{ExportName}}, self).__init__(parent)
        bus = QDBusConnection.{{GetBusType}}Bus()
        self._proxyer = {{ExportName}}.Proxyer(bus, path, self)
{{with .Properties}}
        self._propIfc = QDBusInterface("{{DestName}}", self.path, "org.freedesktop.DBus.Properties", bus, parent)
{{end}}
{{range .Properties}}
    @pyqtProperty(QVariant)
    def {{.Name}}(self):
        return QDBusReply(self._propIfc.call("Get", "{{IfcName}}", "{{.Name}}")).value()
    @{{.Name}}.setter
    def {{.Name}}(self, value):
        self._propIfc.asynCall("Set", "{{IfcName}}", "{{.Name}}", value)
{{end}}
{{range .Methods }}{{$outNum := CalcArgNum .Args "out"}}{{$inNum := (CalcArgNum .Args "in")}}
    @pyqtSlot({{Repeat "QVariant" ", " $inNum}}{{if gt $outNum 0}}{{if gt $inNum 0}}, {{end}}result=QVariant{{end}})
    def {{.Name}} (self{{range .Args}}{{if eq .Direction "in"}}, {{.Name}}{{end}}{{end}}):
        msg = self._proxyer.call("{{.Name}}" {{GetParamterNames .Args}})
        reply = QDBusReply(msg)
        if reply.isValid():{{if gt $outNum 1}}
                return list(msg.arguments()){{else}}{{if eq $outNum 0}}
                pass{{else}}
                return reply.value(){{end}}{{end}}
        else:
                raise(Exception(reply.error().message()))
{{end}}
`
var m = `
func Get{{ExportName}}(path string) *{{ExportName}} {
	return  &{{ExportName}}{dbus.ObjectPath(path), getBus().Object("{{DestName}}", dbus.ObjectPath(path)){{if .Signals}},make(chan *dbus.Signal){{end}}}
}

`

var __TEST_TEMPLATE_PyQt = `/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package {{PkgName}}
import "testing"
{{range .Methods}}
func Test{{ObjName}}Method{{.Name}} (t *testing.T) {
	{{/*
	rnd := rand.New(rand.NewSource(99))
	r := Get{{ObjName}}("{{TestPath}}").{{.Name}}({{.Args}})
--*/}}

}
{{end}}

{{range .Properties}}
func Test{{ObjName}}Property{{.Name}} (t *testing.T) {
	t.Log("Get the property {{.Name}} of object {{ObjName}} ===> ",
		Get{{ObjName}}("{{TestPath}}").Get{{.Name}}())
}
{{end}}

{{range .Signals}}
func Test{{ObjName}}Signal{{.Name}} (t *testing.T) {
}
{{end}}
`
