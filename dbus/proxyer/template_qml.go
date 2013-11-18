package main

import "os"
import "path"
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
    Q_PROPERTY({{getQType .Type}} {{.Name}} NOTIFY {{.Name}}Changed)
    Q_SIGNAL void {{.Name}}Changed({{getQType .Type}} {{.Name}});{{end}}

Q_SIGNALS:{{range .Signals}}
    void {{.Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}} {{$e.Name}}{{end}});{{end}}
};

class {{ExportName}} : public QObject 
{
    Q_OBJECT
private:
    QString m_path;
    void rebuild() 
    { 
	  delete m_ifc;
          m_ifc = new {{ExportName}}Proxyer(m_path);{{range .Signals}}
	  QObject::connect(m_ifc, SIGNAL({{.Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}}{{end}})), 
	  		this, SIGNAL({{Lower .Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}}{{end}})));{{end}}
          //TODO: Signal arguments convert{{range .Properties}}
	  QObject::connect(m_ifc, SIGNAL({{.Name}}({{getQType .Type}})), this, SIGNAL({{Lower .Name}}({{getQType .Type}})));{{end}}
    }
public:
    Q_PROPERTY(QString path READ path WRITE setPath NOTIFY pathChanged)
    const QString path() {
	    return m_path;
    }
    void setPath(const QString& path) {
	    m_path = path;
	    rebuild();
    }
    Q_SIGNAL void pathChanged(QString);

    {{ExportName}}(QObject *parent=0) : QObject(parent), m_ifc(new {{ExportName}}Proxyer("{{IfcName}}"))
    {
    }
    {{range .Properties}}
    Q_PROPERTY(QVariant {{Lower .Name}} READ {{.Name}} NOTIFY {{Lower .Name}}Changed){{end}}

    //Property read methods{{range .Properties}}
    const QVariant {{.Name}}() {
	    return tryConvert(m_ifc->property("{{.Name}}"));
    }{{end}}

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
    void {{Lower .Name}}Changed();{{end}}

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
		qDebug() << "D:" << uri;
		{{range GetModules}}
	    qmlRegisterType<{{Upper .}}>(uri, 1, 0, "{{Upper .}}");{{end}}
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
	return v;
	if (QString(v.typeName()).startsWith("QList")) {
		QVariantList list;
		foreach(const QDBusObjectPath &p, v.value<QList<QDBusObjectPath> >()) {
			/*qDebug() << "Prop :" << p.path();*/
			list.append(tryConvert(QVariant::fromValue(p)));
		}
		return list;
	} else if (v.userType() == qMetaTypeId<QDBusObjectPath>()) {
		return QVariant::fromValue(v.value<QDBusObjectPath>().path());
	} else if (v.userType() == qMetaTypeId<QDBusArgument>()) {
		return tryConvert(parse(v.value<QDBusArgument>()));
	}
	return v;
}
`
var __PROJECT_TEMPL_QML = `
TEMPLATE=lib
CONFIC += plugin
QT += qml dbus

QMAKE_CC=clang
QMAKE_CXX=clang++

TARGET = {{PkgName}}
DESTDIR = lib

OBJECTS_DIRS = tmp
MOC_DIR = tmp

HEADERS += plugin.h {{range GetModules}}{{.}}.h {{end}}
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
