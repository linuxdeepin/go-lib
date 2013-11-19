package main

import "fmt"
import "os"
import "os/exec"
import "path"
import "strings"
import "text/template"

var __IFC_TEMPLATE_INIT_QML = `/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
#include <QtDBus>
QVariant tryConvert(const QVariant&);
`

var __IFC_TEMPLATE_QML = `
#ifndef __{{ExportName}}_H__
#define __{{ExportName}}_H__

class {{ExportName}}Proxyer: public QDBusAbstractInterface
{
    Q_OBJECT
public:
    {{ExportName}}Proxyer(const QString &path)
          :QDBusAbstractInterface("{{DestName}}", path, "{{IfcName}}", QDBusConnection::{{BusType}}Bus(), 0)
    {
    }

    ~{{ExportName}}Proxyer()
    {
    };
{{range .Properties}}
    Q_PROPERTY(QVariant {{.Name}} READ {{.Name}} WRITE set{{.Name}})
    QVariant {{.Name}}() { return tryConvert(property("{{.Name}}")); }
    void set{{.Name}}(const QVariant &v) { setProperty("{{.Name}}", v); }
    {{end}}

Q_SIGNALS:{{range .Signals}}
    void {{.Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}} {{$e.Name}}{{end}});{{end}}
};

class {{ExportName}} : public QObject 
{
    Q_OBJECT
private:
    QString m_path;
    Q_SLOT void _propertiesChanged(const QDBusMessage &msg) {
	    QList<QVariant> arguments = msg.arguments();
	    if (3 != arguments.count())
	    	return;
	    QString interfaceName = msg.arguments().at(0).toString();
	    if (interfaceName != "{{IfcName}}")
		    return;

	    QVariantMap changedProps = qdbus_cast<QVariantMap>(arguments.at(1).value<QDBusArgument>());
	    foreach(const QString &prop, changedProps.keys()) {
		    if (0) { {{range .Properties}}
		    } else if (prop == "{{.Name}}") {
			    Q_EMIT {{Lower .Name}}Changed(changedProps.value(prop));{{end}}
		    }
	    }
    }
    void _rebuild() 
    { 
	  delete m_ifc;
          m_ifc = new {{ExportName}}Proxyer(m_path);{{range .Signals}}
	  QObject::connect(m_ifc, SIGNAL({{.Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}}{{end}})), 
	  		this, SIGNAL({{Lower .Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}}{{end}})));{{end}}
    }
public:
    Q_PROPERTY(QString path READ path WRITE setPath NOTIFY pathChanged)
    const QString path() {
	    return m_path;
    }
    void setPath(const QString& path) {
	    QDBusConnection::{{BusType}}Bus().disconnect("{{DestName}}", m_path, "org.freedesktop.DBus.Properties", "PropertiesChanged",
	    				 this, SLOT(_propertiesChanged(QDBusMessage)));
	    m_path = path;
	    QDBusConnection::{{BusType}}Bus().connect("{{DestName}}", m_path, "org.freedesktop.DBus.Properties", "PropertiesChanged",
	    				"sa{sv}as", this, SLOT(_propertiesChanged(QDBusMessage)));
	    _rebuild();
    }
    Q_SIGNAL void pathChanged(QString);

    {{ExportName}}(QObject *parent=0) : QObject(parent), m_ifc(new {{ExportName}}Proxyer("{{Ifc2Obj IfcName}}"))
    {
	    QDBusConnection::{{BusType}}Bus().connect("{{DestName}}", m_path, "org.freedesktop.DBus.Properties", "PropertiesChanged",
	    				"sa{sv}as", this, SLOT(_propertiesChanged(QDBusMessage)));
    }
    {{range .Properties}}
    Q_PROPERTY(QVariant {{Lower .Name}} READ {{.Name}} WRITE set{{.Name}} NOTIFY {{Lower .Name}}Changed){{end}}

    //Property read methods{{range .Properties}}
    const QVariant {{.Name}}() { return tryConvert(m_ifc->property("{{.Name}}")); }{{end}}
    //Property set methods :TODO check access{{range .Properties}}
    void set{{.Name}}(const QVariant &v) { m_ifc->setProperty("{{.Name}}", v); }{{end}}

public Q_SLOTS:{{range .Methods}}
    QVariant {{.Name}}({{range $i, $e := GetOuts .Args}}{{if ne $i 0}}, {{end}}const QVariant &{{.Name}}{{end}}) {
	    QList<QVariant> argumentList;{{range GetOuts .Args}}{{if NormaliseQDBus .Type}}
	    {{.Name}} = {{NormaliseQDBus .Type}}{{end}}{{end}}
	    argumentList{{range GetOuts .Args}} << {{.Name}}{{end}};

	    QDBusPendingReply<> call = m_ifc->asyncCallWithArgumentList(QLatin1String("{{.Name}}"), argumentList);
	    call.waitForFinished();
	    if (call.isValid()) {
		    QList<QVariant> args = call.reply().arguments();
		    switch (args.size()) {
			    case 0: return QVariant();
			    case 1: {
				    return tryConvert(args[0]);
			    }
		    default:
			    {
				    for (int i=0; i<args.size(); i++) {
					    args[i] = tryConvert(args[i]);
				    }
				    return args;
			    }
		    }
	    } else {
		    qDebug() << "Error:" << call.error().message();
		    return QVariant();
	    }
    }
{{end}}

Q_SIGNALS:
//Property changed notify signal{{range .Properties}}
    void {{Lower .Name}}Changed(QVariant);{{end}}

//DBus Interface's signal{{range .Signals}}
    void {{Lower .Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}} {{$e.Name}}{{end}});{{end}}
private:
    {{ExportName}}Proxyer *m_ifc;
};

#endif
`

var __GLOBAL_TEMPLATE_QML = `
#ifndef __DBUS_H__
#define __DBUS_H__

{{range GetModules}}
#include "{{.}}.h"{{end}}
#include <QQmlExtensionPlugin>
#include <qqml.h>

class DBusPlugin: public QQmlExtensionPlugin
{
    Q_OBJECT
	Q_PLUGIN_METADATA(IID "com.deepin.dde.daemon.DBus")

    public:
	void registerTypes(const char* uri) {
		qDebug() << "D:" << uri;{{range .Interfaces}}
	    qmlRegisterType<{{.ObjectName}}>(uri, 1, 0, "{{.ObjectName}}");{{end}}
    }
};
#endif


inline 
QVariant parse(const QDBusArgument &argument)
{
    switch (argument.currentType()) {
    case QDBusArgument::BasicType: {
        QVariant v = argument.asVariant();
        if (v.userType() == qMetaTypeId<QDBusObjectPath>())
            return v.value<QDBusObjectPath>().path();
        else if (v.userType() == qMetaTypeId<QDBusSignature>())
            return v.value<QDBusSignature>().signature();
        else
            return v;
    }
    case QDBusArgument::VariantType: {
        QVariant v = argument.asVariant().value<QDBusVariant>().variant();
        if (v.userType() == qMetaTypeId<QDBusArgument>())
            return parse(v.value<QDBusArgument>());
        else
            return v;
    }
    case QDBusArgument::ArrayType: {
        QVariantList list;
        argument.beginArray();
        while (!argument.atEnd())
            list.append(parse(argument));
        argument.endArray();
        return list;
    }
    case QDBusArgument::StructureType: {
        QVariantList list;
        argument.beginStructure();
        while (!argument.atEnd())
            list.append(parse(argument));
        argument.endStructure();
        return QVariant::fromValue(list);
    }
    case QDBusArgument::MapType: {
        QVariantMap map;
        argument.beginMap();
        while (!argument.atEnd()) {
            argument.beginMapEntry();
            QVariant key = parse(argument);
            QVariant value = parse(argument);
            map.insert(key.toString(), value);
            argument.endMapEntry();
        }
        argument.endMap();
        return map;
    }
    default:
        return QVariant();
        break;
    }
}

QVariant tryConvert(const QVariant& v) {
	if (QString(v.typeName()).startsWith("QList")) {
		QVariantList list;
		foreach(const QDBusObjectPath &p, v.value<QList<QDBusObjectPath> >()) {
			list.append(tryConvert(QVariant::fromValue(p)));
		}
		return list;
	} else if (v.userType() == qMetaTypeId<QDBusObjectPath>()) {
		return QVariant::fromValue(v.value<QDBusObjectPath>().path());
	} else if (v.userType() == qMetaTypeId<QDBusArgument>()) {
		return tryConvert(parse(v.value<QDBusArgument>()));
	} else if (v.userType() == qMetaTypeId<QByteArray>()) {
		return QString(v.value<QByteArray>());
	}
	return v;
}
`
var __PROJECT_TEMPL_QML = `
TEMPLATE=lib
CONFIC += plugin
QT += qml dbus

TARGET = {{PkgName}}
DESTDIR = lib

OBJECTS_DIRS = tmp
MOC_DIR = tmp

HEADERS += plugin.h {{range GetModules}}{{.}}.h {{end}}


test.depends = {{PkgName}}/$(TARGET)
test.commands = (qmlscene -I . test.qml)
QMAKE_EXTRA_TARGETS += test
`

var __TEST_QML = `
import {{PkgName}} 1.0
import QtQuick 2.0
import QtQuick.Controls 1.0

Item { {{range .Interfaces}}
    {{.ObjectName}} {
       id: "{{Lower .ObjectName}}ID"
       // path: "{{Ifc2Obj .Interface}}"
    } {{end}}
    width: 400; height: 400
    TabView {
	    anchors.fill  : parent
	    {{range .Interfaces}}
	    Tab {   {{$ifc := GetInterfaceInfo .}} {{$objName := Lower .ObjectName }}
		    title: "{{.ObjectName}}"
		    Column {
			    {{range $ifc.Properties}}
			    Row {
				    Label {
					    text: "{{.Name}}:"
				    }
				    Text {
					    text: JSON.stringify({{$objName}}ID.{{Lower .Name}})
				    }
			    }{{end}}
		    }
	    }
	    {{end}}
    }
}
`

func renderQMLProject() {
	writer, err := os.Create(path.Join(INFOS.Config.OutputDir, "tt.pro"))
	if err != nil {
		panic(err)
	}
	template.Must(template.New("main").Funcs(template.FuncMap{
		"BusType": func() string { return INFOS.Config.BusType },
		"PkgName": func() string { return INFOS.Config.PkgName },
		"GetModules": func() map[string]string {
			r := make(map[string]string)
			for _, ifc := range INFOS.Interfaces {
				r[ifc.OutFile] = ifc.OutFile
			}
			return r
		},
	}).Parse(__PROJECT_TEMPL_QML)).Execute(writer, INFOS)
	writer.Close()
}
func testQML() {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && qmake", INFOS.Config.OutputDir))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	/*qmldir, err := os.Create(path.Join(INFOS.Config.OutputDir, INFOS.Config.PkgName, "qmldir"))*/
	qmldir, err := os.Create(path.Join(INFOS.Config.OutputDir, "lib", "qmldir"))
	if err != nil {
		panic(err)
	}
	qmldir.WriteString("module " + INFOS.Config.PkgName + "\n")
	qmldir.WriteString("plugin " + INFOS.Config.PkgName)
	qmldir.Close()

	writer, err := os.Create(path.Join(INFOS.Config.OutputDir, "test.qml"))
	if err != nil {
		panic(err)
	}
	template.Must(template.New("qmltest").Funcs(template.FuncMap{
		"Lower":            lower,
		"GetInterfaceInfo": GetInterfaceInfo,
		"BusType":          func() string { return INFOS.Config.BusType },
		"PkgName":          func() string { return INFOS.Config.PkgName },
		"Ifc2Obj":          ifc2obj,
		"GetModules": func() map[string]string {
			r := make(map[string]string)
			for _, ifc := range INFOS.Interfaces {
				r[ifc.OutFile] = ifc.OutFile
			}
			return r
		},
	}).Parse(__TEST_QML)).Execute(writer, INFOS)

}
func qtPropertyFilter(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "QMap") {
		return "QVariantMap"
	} else if strings.HasPrefix(s, "QList") {
		return "QVariantList"
	} else if strings.HasPrefix(s, "QValueList") {
		return "QVariantValueList"
	}
	return s
}
