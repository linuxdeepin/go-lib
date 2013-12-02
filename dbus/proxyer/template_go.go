package main

var __GLOBAL_TEMPLATE_GoLang = `
package {{PkgName}}
import "dlib/dbus"
var __conn *dbus.Conn = nil
func getBus() *dbus.Conn {
	if __conn  == nil {
		var err error
		__conn, err = dbus.{{BusType}}Bus()
		if err != nil {
			panic(err)
		}
	}
	return __conn
}
`

var __IFC_TEMPLATE_INIT_GoLang = `/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package {{PkgName}}
import "dlib/dbus"
import "reflect"
var _ = reflect.TypeOf /*prevent compile error*/
`

var __IFC_TEMPLATE_GoLang = `
type {{ExportName}} struct {
	Path dbus.ObjectPath
	core *dbus.Object
	{{if .Signals}}signal_chan chan *dbus.Signal{{end}}
}
{{$obj_name := .Name}}
{{range .Methods }}
func ({{OBJ_NAME}} {{ExportName }}) {{.Name}} ({{GetParamterInsProto .Args}}) ({{GetParamterOutsProto .Args}}) {
	{{OBJ_NAME}}.core.Call("{{$obj_name}}.{{.Name}}", 0{{GetParamterNames .Args}}).Store({{GetParamterOuts .Args}})
	return
}
{{end}}

{{range .Signals}}
func ({{OBJ_NAME}} {{ExportName}}) Connect{{.Name}}(callback func({{GetParamterOutsProto .Args}})) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string({{OBJ_NAME}}.core.Path())+"', interface='{{IfcName}}',sender='{{DestName}}',member='{{.Name}}'")
	__conn.Signal({{OBJ_NAME}}.signal_chan)
	go func() {
		for v := range({{OBJ_NAME}}.signal_chan) {
			if v.Name != "{{IfcName}}.{{.Name}}" || {{len .Args}} != len(v.Body) {
				continue
			}
			{{range $index, $arg := .Args}}if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*{{TypeFor $arg.Type}})(nil)).Elem() {
				continue
			}
			{{end}}

			callback({{range $index, $arg := .Args}}{{if $index}},{{end}}v.Body[{{$index}}].({{TypeFor $arg.Type}}){{end}})
		}
	}()

}
{{end}}

{{range .Properties}}
func ({{OBJ_NAME}} *{{ExportName}}) Set{{.Name}}({{.Name}} {{TypeFor .Type}}) {
	{{OBJ_NAME}}.core.Call("org.freedesktop.DBus.Properties.Set", 0, "{{IfcName}}", "{{.Name}}", {{.Name}})
}
{{ $convert := TryConvertObjectPath . }}
func ({{OBJ_NAME}} {{ExportName}}) Get{{.Name}}() (ret {{GetObjectPathType .}}) {
	var r dbus.Variant
	err := {{OBJ_NAME}}.core.Call("org.freedesktop.DBus.Properties.Get", 0, "{{IfcName}}", "{{.Name}}").Store(&r)
	if err == nil && r.Signature().String() == "{{.Type}}" { {{ if $convert }}
		before := r.Value().({{TypeFor .Type}})
		{{$convert}}
		return after{{else}}
		return r.Value().({{TypeFor .Type}}){{end}}
	}  else {
		panic(err)
	}
	return
}
{{end}}

func Get{{ExportName}}(path string) *{{ExportName}} {
	return  &{{ExportName}}{dbus.ObjectPath(path), getBus().Object("{{DestName}}", dbus.ObjectPath(path)){{if .Signals}},make(chan *dbus.Signal){{end}}}
}

`

var __TEST_TEMPLATE = `/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
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
