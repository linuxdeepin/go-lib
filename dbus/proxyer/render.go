package main

import "text/template"
import "dlib/dbus"
import "io"
import "log"

func getGlobalTemplate() string {
	if INFOS.Config.Target == "PyQt" {
		return __GLOBAL_TEMPLATE_PyQT
	}
	return __GLOBAL_TEMPLATE
}
func getInterfaceTemplate() string {
	if INFOS.Config.Target == "PyQt" {
		return __IFC_TEMPLATE_PyQt
	}
	return __IFC_TEMPLATE
}
func getInterfaceInitTemplate() string {
	if INFOS.Config.Target == "PyQt" {
		return __IFC_TEMPLATE_INIT_PyQt
	}
	return __IFC_TEMPLATE_INIT
}

func renderMain(writer io.Writer) {
	template.Must(template.New("main").Funcs(template.FuncMap{
		"GetBusType": func() string { return INFOS.Config.BusType },
		"PkgName":    func() string { return INFOS.Config.PkgName },
	}).Parse(getGlobalTemplate())).Execute(writer, nil)
}

func renderInterfaceInit(writer io.Writer) {
	template.Must(template.New("IfcInit").Funcs(template.FuncMap{
		"GetBusType": func() string { return INFOS.Config.BusType },
		"PkgName":    func() string { return INFOS.Config.PkgName },
		"HasSignals": func() bool { return true },
	}).Parse(getInterfaceInitTemplate())).Execute(writer, nil)
}

func renderInterface(lang string, pkgName string, info dbus.InterfaceInfo, writer io.Writer, dest, ifc_name, exportName string) {
	if lang == "GoLang" {
		filterKeyWord(getGoKeyword, &info)
	} else if lang == "PyQt" {
		filterKeyWord(getPyQtKeyword, &info)
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

func renderTest(testPath, pkgName string, objName string, writer io.Writer, info dbus.InterfaceInfo) {
	funcs := template.FuncMap{
		"TestPath": func() string { return testPath },
		"PkgName":  func() string { return pkgName },
		"ObjName":  func() string { return objName },
		/*"GetTestValue": func(args []dbus.ArgInfo) string {*/
		/*},*/
	}
	template.Must(template.New("testing").Funcs(funcs).Parse(__TEST_TEMPLATE)).Execute(writer, info)
}
