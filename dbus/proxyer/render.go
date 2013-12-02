package main

import "text/template"
import "dlib/dbus"
import "io"
import "log"

var TEMPLs = map[string]string{
	"GLOBAL_pyqt":   __GLOBAL_TEMPLATE_PyQt,
	"GLOBAL_golang": __GLOBAL_TEMPLATE_GoLang,
	"GLOBAL_qml":    __GLOBAL_TEMPLATE_QML,

	"IFC_pyqt":   __IFC_TEMPLATE_PyQt,
	"IFC_golang": __IFC_TEMPLATE_GoLang,
	"IFC_qml":    __IFC_TEMPLATE_QML,

	"IFC_INIT_pyqt":   __IFC_TEMPLATE_INIT_PyQt,
	"IFC_INIT_golang": __IFC_TEMPLATE_INIT_GoLang,
	"IFC_INIT_qml":    __IFC_TEMPLATE_INIT_QML,
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

func renderInterface(info dbus.InterfaceInfo, writer io.Writer, ifc_name, exportName string) {
	if INFOS.Config.Target == GoLang {
		filterKeyWord(getGoKeyword, &info)
	} else if INFOS.Config.Target == PyQt {
		filterKeyWord(getPyQtKeyword, &info)
	} else if INFOS.Config.Target == QML {
		filterKeyWord(getGoKeyword, &info)
	}
	log.Printf("Generate %q code for service:%q interface:%q ObjectName:%q", INFOS.Config.Target, INFOS.Config.DestName, ifc_name, exportName)
	funcs := template.FuncMap{
		"Lower":          lower,
		"Upper":          upper,
		"BusType":        func() string { return INFOS.Config.BusType },
		"PkgName":        func() string { return INFOS.Config.PkgName },
		"OBJ_NAME":       func() string { return "obj" },
		"TypeFor":        func(s string) string { return dbus.TypeFor(s) },
		"getQType":       getQType,
		"DestName":       func() string { return INFOS.Config.DestName },
		"IfcName":        func() string { return ifc_name },
		"ExportName":     func() string { return exportName },
		"NormaliseQDBus": normaliseQDBus,
		"Ifc2Obj":        ifc2obj,
		"PropWritable":   func(prop dbus.PropertyInfo) bool { return prop.Access == "readwrite" },
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
		"TryConvertObjectPath": func(prop dbus.PropertyInfo) string {
			if v := getObjectPathConvert("Property", prop.Annotations); v != "" {
				switch INFOS.Config.Target {
				case GoLang:
					return tryConvertObjectPathGo(prop.Type, v)
				case QML:
					return tryConvertObjectPathQML(prop.Type, v)
				}
			}
			return ""
		},
		"GetObjectPathType": func(prop dbus.PropertyInfo) (ret string) {
			if v := getObjectPathConvert("Property", prop.Annotations); v != "" {
				switch INFOS.Config.Target {
				case GoLang:
					ret, _ = guessTypeGo(prop.Type, v)
				case QML:
					ret, _ = guessTypeQML(prop.Type, v)
				}
				return
			}
			return dbus.TypeFor(prop.Type)
		},
	}
	templ := template.Must(template.New(exportName).Funcs(funcs).Parse(TEMPLs["IFC_"+INFOS.Config.Target]))
	templ.Execute(writer, info)
}

func renderTest(testPath, objName string, writer io.Writer, info dbus.InterfaceInfo) {
	funcs := template.FuncMap{
		"TestPath": func() string { return testPath },
		"PkgName":  func() string { return INFOS.Config.PkgName },
		"ObjName":  func() string { return objName },
		/*"GetTestValue": func(args []dbus.ArgInfo) string {*/
		/*},*/
	}
	template.Must(template.New("testing").Funcs(funcs).Parse(__TEST_TEMPLATE)).Execute(writer, info)
}
