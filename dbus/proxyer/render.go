package main

import "strings"
import "text/template"
import "dlib/dbus"
import "io"
import "log"

func lower(str string) string   { return strings.ToLower(str[:1]) + str[1:] }
func upper(str string) string   { return strings.ToUpper(str[:1]) + str[1:] }
func ifc2obj(ifc string) string { return "/" + strings.Replace(ifc, ".", "/", -1) }

var TEMPLs = map[string]string{
	"GLOBAL_PyQt":   __GLOBAL_TEMPLATE_PyQt,
	"GLOBAL_GoLang": __GLOBAL_TEMPLATE_GoLang,
	"GLOBAL_QML":    __GLOBAL_TEMPLATE_QML,

	"IFC_PyQt":   __IFC_TEMPLATE_PyQt,
	"IFC_GoLang": __IFC_TEMPLATE_GoLang,
	"IFC_QML":    __IFC_TEMPLATE_QML,

	"IFC_INIT_PyQt":   __IFC_TEMPLATE_INIT_PyQt,
	"IFC_INIT_GoLang": __IFC_TEMPLATE_INIT_GoLang,
	"IFC_INIT_QML":    __IFC_TEMPLATE_INIT_QML,
}

func renderMain(writer io.Writer) {
	template.Must(template.New("main").Funcs(template.FuncMap{
		"Lower":   lower,
		"Upper":   upper,
		"BusType": func() string { return INFOS.Config.BusType },
		"PkgName": func() string { return INFOS.Config.PkgName },
		"GetModules": func() map[string]string {
			r := make(map[string]string)
			for _, ifc := range INFOS.Interfaces {
				r[ifc.OutFile] = ifc.OutFile
			}
			return r
		},
	}).Parse(TEMPLs["GLOBAL_"+INFOS.Config.Target])).Execute(writer, INFOS)
}

func renderInterfaceInit(writer io.Writer) {
	template.Must(template.New("IfcInit").Funcs(template.FuncMap{
		"BusType":    func() string { return INFOS.Config.BusType },
		"PkgName":    func() string { return INFOS.Config.PkgName },
		"HasSignals": func() bool { return true },
	}).Parse(TEMPLs["IFC_INIT_"+INFOS.Config.Target])).Execute(writer, nil)
}

func renderInterface(target string, pkgName string, info dbus.InterfaceInfo, writer io.Writer, dest, ifc_name, exportName string) {
	if target == "GoLang" {
		filterKeyWord(getGoKeyword, &info)
	} else if target == "PyQt" {
		filterKeyWord(getPyQtKeyword, &info)
	} else if target == "QML" {
		filterKeyWord(getGoKeyword, &info)
	}
	log.Println("d:", dest, "i:", ifc_name, "e:", exportName)
	funcs := template.FuncMap{
		"Lower":          lower,
		"Upper":          upper,
		"BusType":        func() string { return INFOS.Config.BusType },
		"PkgName":        func() string { return pkgName },
		"OBJ_NAME":       func() string { return "obj" },
		"TypeFor":        func(s string) string { return dbus.TypeFor(s) },
		"getQType":       getQType,
		"DestName":       func() string { return dest },
		"IfcName":        func() string { return ifc_name },
		"ExportName":     func() string { return exportName },
		"NormaliseQDBus": normaliseQDBus,
		"Ifc2Obj":        ifc2obj,
		"GetOuts": func(args []dbus.ArgInfo) []dbus.ArgInfo {
			ret := make([]dbus.ArgInfo, 0)
			for _, a := range args {
				if a.Direction != "out" {
					ret = append(ret, a)
				}
			}
			return ret
		},
		"CalcArgNum": func(args []dbus.ArgInfo, direction string) (r int) {
			for _, arg := range args {
				if arg.Direction == direction {
					r++
				}
			}
			return
		},
		"Repeat": func(str string, sep string, times int) (r string) {
			for i := 0; i < times; i++ {
				if i != 0 {
					r += sep
				}
				r += str
			}
			return
		},
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
	templ := template.Must(template.New(exportName).Funcs(funcs).Parse(TEMPLs["IFC_"+INFOS.Config.Target]))
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
