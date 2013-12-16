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
import "dlib/dbus/property"
import "reflect"
import "sync"
import "runtime"
import "log"
/*prevent compile error*/
var _ = runtime.SetFinalizer
var _ = sync.NewCond
var _ = reflect.TypeOf
var _ = property.BaseObserver{}
`

var __IFC_TEMPLATE_GoLang = `
type {{ExportName}} struct {
	Path dbus.ObjectPath
	core *dbus.Object
{{if or .Properties .Signals}}
	signals map[chan *dbus.Signal]bool
	signalsLocker sync.Mutex
{{end}}
	{{range .Properties}}
	{{.Name}} *dbusProperty{{ExportName}}{{.Name}}{{end}}
}
{{if or .Properties .Signals}}
func ({{OBJ_NAME}} {{ExportName}}) _createSignalChan() chan *dbus.Signal {
	{{OBJ_NAME}}.signalsLocker.Lock()
	ch := make(chan *dbus.Signal)
	getBus().Signal(ch)
	{{OBJ_NAME}}.signals[ch] = false
	{{OBJ_NAME}}.signalsLocker.Unlock()
	return ch
}
func ({{OBJ_NAME}} {{ExportName}}) _deleteSignalChan(ch chan *dbus.Signal) {
	{{OBJ_NAME}}.signalsLocker.Lock()
	delete({{OBJ_NAME}}.signals, ch)
	getBus().DetachSignal(ch)
	close(ch)
	{{OBJ_NAME}}.signalsLocker.Unlock()
}
func destroy{{ExportName}}(obj *{{ExportName}}) {
	obj.signalsLocker.Lock()
	for ch, _ := range obj.signals {
		delete({{OBJ_NAME}}.signals, ch)
		getBus().DetachSignal(ch)
		close(ch)
	}
	log.Printf("Debug: run destroy{{ExportName}}", obj)
	obj.signalsLocker.Unlock()
}
{{end}}

{{$obj_name := .Name}}
{{range .Methods }}
func ({{OBJ_NAME}} {{ExportName }}) {{.Name}} ({{GetParamterInsProto .Args}}) ({{GetParamterOutsProto .Args}}) {
	err := {{OBJ_NAME}}.core.Call("{{$obj_name}}.{{.Name}}", 0{{GetParamterNames .Args}}).Store({{GetParamterOuts .Args}})
	if err != nil {
		log.Println("Invoked", {{OBJ_NAME}}.Path, ":{{$obj_name}}.{{.Name}}("{{GetParamterNames .Args}}, ") failed:", err)
	}
	return
}
{{end}}

{{range .Signals}}
func ({{OBJ_NAME}} {{ExportName}}) Connect{{.Name}}(callback func({{GetParamterOutsProto .Args}})) func() {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string({{OBJ_NAME}}.core.Path())+"', interface='{{IfcName}}',sender='{{DestName}}',member='{{.Name}}'")
	sigChan := {{OBJ_NAME}}._createSignalChan()
	go func() {
		for v := range(sigChan) {
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
	return func() {
		{{OBJ_NAME}}._deleteSignalChan(sigChan)
	}
}
{{end}}

{{range .Properties}}
type dbusProperty{{ExportName}}{{.Name}} struct{
	*property.BaseObserver
	core *dbus.Object
}
{{if PropWritable .}}func (this *dbusProperty{{ExportName}}{{.Name}}) SetValue(v interface{}/*{{TypeFor .Type}}*/) {
	if reflect.TypeOf(v) == reflect.TypeOf((*{{TypeFor .Type}})(nil)).Elem() {
		this.core.Call("org.freedesktop.DBus.Properties.Set", 0, "{{IfcName}}", "{{.Name}}", dbus.MakeVariant(v))
	} else {
		log.Println("The property {{.Name}} of {{IfcName}} is an {{TypeFor .Type}} but Set with an ", reflect.TypeOf(v))
	}
}
func (this *dbusProperty{{ExportName}}{{.Name}}) Set(v {{TypeFor .Type}}) {
	this.SetValue(v)
}{{else}}
func (this *dbusProperty{{ExportName}}{{.Name}}) SetValue(notwritable interface{}) {
	log.Printf("{{IfcName}}.{{.Name}} is not writable")
}{{end}}
{{ $convert := TryConvertObjectPath . }}
func (this *dbusProperty{{ExportName}}{{.Name}}) Get() {{GetObjectPathType .}} {
	return this.GetValue().({{GetObjectPathType .}})
}
func (this *dbusProperty{{ExportName}}{{.Name}}) GetValue() interface{} /*{{GetObjectPathType .}}*/ {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "{{IfcName}}", "{{.Name}}").Store(&r)
	if err == nil && r.Signature().String() == "{{.Type}}" { {{ if $convert }}
		before := r.Value().({{TypeFor .Type}})
		{{$convert}}
		return after{{else}}
		return r.Value().({{TypeFor .Type}}){{end}}
	}  else {
		log.Println("dbusProperty:{{.Name}} error:", err, "at {{IfcName}}")
		return *new({{TypeFor .Type}})
	}
}
func (this *dbusProperty{{ExportName}}{{.Name}}) GetType() reflect.Type {
	return reflect.TypeOf((*{{TypeFor .Type}})(nil)).Elem()
}
{{end}}

func Get{{ExportName}}(path string) *{{ExportName}} {
	core := getBus().Object("{{DestName}}", dbus.ObjectPath(path))
	obj := &{{ExportName}}{Path:dbus.ObjectPath(path), core:core{{if or .Signals .Properties}},signals:make(map[chan *dbus.Signal]bool){{end}}}
	{{range .Properties}}
	obj.{{.Name}} = &dbusProperty{{ExportName}}{{.Name}}{&property.BaseObserver{}, core}{{end}}
{{with .Properties}}
	getBus().BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',path='"+path+"',interface='org.freedesktop.DBus.Properties',sender='{{DestName}}',member='PropertiesChanged'")
	getBus().BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',path='"+path+"',interface='{{IfcName}}',sender='{{DestName}}',member='PropertiesChanged'")
	sigChan := obj._createSignalChan()
	go func() {
		typeString := reflect.TypeOf("")
		typeKeyValues := reflect.TypeOf(map[string]dbus.Variant{})
		typeArrayValues := reflect.TypeOf([]string{})
		for v := range(sigChan) {
			if v.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" &&
				len(v.Body) == 3 &&
				reflect.TypeOf(v.Body[0]) == typeString &&
				reflect.TypeOf(v.Body[1]) == typeKeyValues &&
				reflect.TypeOf(v.Body[2]) == typeArrayValues &&
				v.Body[0].(string) != "{{IfcName}}" {
				props := v.Body[1].(map[string]dbus.Variant)
				for key, _ := range props {
					if false { {{range .}}
					} else if key == "{{.Name}}" {
						obj.{{.Name}}.Notify()
					{{end}} }
				}
			} else if v.Name == "{{IfcName}}.PropertiesChanged" && len(v.Body) == 1 && reflect.TypeOf(v.Body[0]) == typeKeyValues {
				for key, _ := range v.Body[0].(map[string]dbus.Variant) {
					if false { {{range .}}
					} else if key == "{{.Name}}" {
						obj.{{.Name}}.Notify()
					{{end}} }
				}
			}
		}
	}()
{{end}}
{{if or .Properties .Signals}}runtime.SetFinalizer(obj, func(_obj *{{ExportName}}) { destroy{{ExportName}}(_obj) }){{end}}
	return obj
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
