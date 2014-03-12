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
import "fmt"
import "errors"
import "strings"
/*prevent compile error*/
var _ = fmt.Println
var _ = runtime.SetFinalizer
var _ = sync.NewCond
var _ = reflect.TypeOf
var _ = property.BaseObserver{}
`

var __IFC_TEMPLATE_GoLang = `
type {{ExportName}} struct {
	Path dbus.ObjectPath
	DestName string
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
	ch := make(chan *dbus.Signal, 30)
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
func Destroy{{ExportName}}(obj *{{ExportName}}) {
	obj.signalsLocker.Lock()
	for ch, _ := range obj.signals {
		getBus().DetachSignal(ch)
		close(ch)
	}
	obj.signals = make(map[chan *dbus.Signal]bool)
	obj.signalsLocker.Unlock()
}
{{end}}

{{$obj_name := .Name}}
{{range .Methods }}
func ({{OBJ_NAME}} {{ExportName }}) {{.Name}} ({{GetParamterInsProto .Args}}) ({{GetParamterOutsProto .Args}} {{with GetParamterOuts .Args}},{{end}}_err error) {
	_err = {{OBJ_NAME}}.core.Call("{{$obj_name}}.{{.Name}}", 0{{GetParamterNames .Args}}).Store({{GetParamterOuts .Args}})
	if _err != nil {
		fmt.Println(_err)
	}
	return
}
{{end}}

{{range .Signals}}
func ({{OBJ_NAME}} {{ExportName}}) Connect{{.Name}}(callback func({{GetParamterOutsProto .Args}})) func() {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string({{OBJ_NAME}}.Path)+"', interface='{{IfcName}}',sender='{{OBJ_NAME}}.DestName',member='{{.Name}}'")
	sigChan := {{OBJ_NAME}}._createSignalChan()
	go func() {
		for v := range(sigChan) {
			if v.Path != {{OBJ_NAME}}.Path || v.Name != "{{IfcName}}.{{.Name}}" || {{len .Args}} != len(v.Body) {
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
		fmt.Println("The property {{.Name}} of {{IfcName}} is an {{TypeFor .Type}} but Set with an ", reflect.TypeOf(v))
	}
}
func (this *dbusProperty{{ExportName}}{{.Name}}) Set(v {{TypeFor .Type}}) {
	this.SetValue(v)
}{{else}}
func (this *dbusProperty{{ExportName}}{{.Name}}) SetValue(notwritable interface{}) {
	fmt.Println("{{IfcName}}.{{.Name}} is not writable")
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
		fmt.Println("dbusProperty:{{.Name}} error:", err, "at {{IfcName}}")
		return *new({{TypeFor .Type}})
	}
}
func (this *dbusProperty{{ExportName}}{{.Name}}) GetType() reflect.Type {
	return reflect.TypeOf((*{{TypeFor .Type}})(nil)).Elem()
}
{{end}}

func New{{ExportName}}(destName string, path dbus.ObjectPath) (*{{ExportName}}, error) {
	if !path.IsValid() {
		return nil, errors.New("The path of '" + string(path) + "' is invalid.")
	}

	core := getBus().Object(destName, path)
	var v string
	core.Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&v)
	if strings.Index(v, "{{IfcName}}") == -1 {
		return nil, errors.New("'" + string(path) + "' hasn't interface '{{IfcName}}'.")
	}

	obj := &{{ExportName}}{Path:path, DestName:destName, core:core{{if or .Signals .Properties}},signals:make(map[chan *dbus.Signal]bool){{end}}}
	{{range .Properties}}
	obj.{{.Name}} = &dbusProperty{{ExportName}}{{.Name}}{&property.BaseObserver{}, core}{{end}}
{{with .Properties}}
	getBus().BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',path='"+string(path)+"',interface='org.freedesktop.DBus.Properties',sender='"+destName+"',member='PropertiesChanged'")
	getBus().BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',path='"+string(path)+"',interface='{{IfcName}}',sender='"+destName+"',member='PropertiesChanged'")
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
{{if or .Properties .Signals}}runtime.SetFinalizer(obj, func(_obj *{{ExportName}}) { Destroy{{ExportName}}(_obj) }){{end}}
	return obj, nil
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
