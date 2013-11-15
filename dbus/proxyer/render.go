package main

import "text/template"
import "dlib/dbus"
import "io"
import "log"

func getGlobalTemplate() string {
	if INFOS.Config.Language == "PyQt" {
		return __GLOBAL_TEMPLATE_PyQT
	}
	return __GLOBAL_TEMPLATE
}
func getInterfaceTemplate() string {
	if INFOS.Config.Language == "PyQt" {
		return __IFC_TEMPLATE_PyQt
	}
	return __IFC_TEMPLATE
}

func getMainTemplate() *template.Template {
	return template.Must(template.New("main").Funcs(template.FuncMap{
		"GetBusType": func() string { return INFOS.Config.BusType },
		"PkgName":    func() string { return INFOS.Config.PkgName },
	}).Parse(getGlobalTemplate()))
}

func GenInterfaceCode(lang string, pkgName string, info dbus.InterfaceInfo, writer io.Writer, dest, ifc_name, exportName string) {
	if lang == "Golang" {
		filterKeyWord(getGoKeyword, &info)
	}
	log.Println("d:", dest, "i:", ifc_name, "e:", exportName)
	funcs := template.FuncMap{
		"PkgName":    func() string { return pkgName },
		"OBJ_NAME":   func() string { return "obj" },
		"TypeFor":    func(s string) string { return dbus.TypeFor(s) },
		"DestName":   func() string { return dest },
		"IfcName":    func() string { return ifc_name },
		"ExportName": func() string { return exportName },
		"GetParamterNames": func(args []dbus.ArgInfo) (ret string) {
			for _, arg := range args {
				if arg.Direction != "out" {
					ret += ", "
					ret += arg.Name
				}
			}
			return
		},
		"GetParamterOuts": func(args []dbus.ArgInfo) (ret string) {
			var notFirst = false
			for _, arg := range args {
				if arg.Direction != "in" {
					if notFirst {
						ret += ","
					}
					notFirst = true
					ret += "&" + arg.Name
				}
			}
			return
		},
		"GetParamterOutsProto": func(args []dbus.ArgInfo) (ret string) {
			var notFirst = false
			for _, arg := range args {
				if arg.Direction != "in" {
					if notFirst {
						ret += ","
					}
					notFirst = true
					ret += arg.Name + " " + dbus.TypeFor(arg.Type)
				}
			}
			return
		},
		"GetParamterInsProto": func(args []dbus.ArgInfo) (ret string) {
			var notFirst = false
			for _, arg := range args {
				if arg.Direction != "out" {
					if notFirst {
						ret += ","
					}
					notFirst = true
					ret += arg.Name + " " + dbus.TypeFor(arg.Type)
				}
			}
			return
		},
	}
	templ := template.Must(template.New(exportName).Funcs(funcs).Parse(getInterfaceTemplate()))
	templ.Execute(writer, info)
}
