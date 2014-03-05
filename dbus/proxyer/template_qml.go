package main

import "fmt"
import "os"
import "log"
import "os/exec"
import "path"
import "strings"
import "text/template"

var __IFC_TEMPLATE_INIT_QML = `/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
#include <QtDBus>
QVariant unmarsh(const QVariant&);
QVariant marsh(QDBusArgument target, const QVariant& arg, const QString& sig);
`

var __IFC_TEMPLATE_QML = `
#ifndef __{{ExportName}}_H__
#define __{{ExportName}}_H__

class {{ExportName}}Proxyer: public QDBusAbstractInterface
{
    Q_OBJECT
public:
    {{ExportName}}Proxyer(const QString &path, QObject* parent)
          :QDBusAbstractInterface("{{DestName}}", path, "{{IfcName}}", QDBusConnection::{{BusType}}Bus(), parent)
    {
	    if (!isValid()) {
		    qDebug() << "Create {{ExportName}} remote object failed : " << lastError().message();
	    }
    }

{{range .Properties}}
    Q_PROPERTY(QVariant {{.Name}} READ {{.Name}} {{if PropWritable .}}WRITE set{{.Name}}{{end}})
    QVariant {{.Name}}() { return unmarsh(property("{{.Name}}")); }
    {{if PropWritable .}}void set{{.Name}}(const QVariant &v) { setProperty("{{.Name}}", v); }{{end}}
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
			    Q_EMIT ___{{Lower .Name}}Changed___(unmarsh(changedProps.value(prop)));{{end}}
		    }
	    }
    }
    void _rebuild() 
    { 
	  delete m_ifc;
          m_ifc = new {{ExportName}}Proxyer(m_path, this);
	  _setupSignalHandle();
    }
    void _setupSignalHandle() {
{{range .Signals}}
	  QObject::connect(m_ifc, SIGNAL({{.Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}}{{end}})), 
	  		this, SIGNAL({{Lower .Name}}({{range $i, $e := .Args}}{{if ne $i 0}},{{end}}{{getQType $e.Type}}{{end}})));
{{end}}
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

    {{ExportName}}(QObject *parent=0) : QObject(parent), m_ifc(new {{ExportName}}Proxyer("{{Ifc2Obj IfcName}}", this))
    {
	    _setupSignalHandle();
	    QDBusConnection::{{BusType}}Bus().connect("{{DestName}}", m_path, "org.freedesktop.DBus.Properties", "PropertiesChanged",
	    				"sa{sv}as", this, SLOT(_propertiesChanged(QDBusMessage)));
    }
    {{range .Properties}}
    Q_PROPERTY(QVariant {{Lower .Name}} READ {{.Name}} {{if PropWritable .}}WRITE set{{.Name}}{{end}} NOTIFY ___{{Lower .Name}}Changed___){{end}}

    //Property read methods{{range .Properties}}
    const QVariant {{.Name}}() { return unmarsh(m_ifc->property("{{.Name}}")); }{{end}}
    //Property set methods :TODO check access{{range .Properties}}{{if PropWritable .}}
    void set{{.Name}}(const QVariant &v) {
	    QVariant marshedValue = marsh(QDBusArgument(), v, "{{.Type}}");
	    m_ifc->setProperty("{{.Name}}", marshedValue);
	    Q_EMIT ___{{Lower .Name}}Changed___(marshedValue);
    }{{end}}{{end}}

public Q_SLOTS:{{range .Methods}}
    QVariant {{.Name}}({{range $i, $e := GetOuts .Args}}{{if ne $i 0}}, {{end}}const QVariant &{{.Name}}{{end}}) {
	    QList<QVariant> argumentList;
	    argumentList{{range GetOuts .Args}} << marsh(QDBusArgument(), {{.Name}}, "{{.Type}}"){{end}};

	    QDBusPendingReply<> call = m_ifc->asyncCallWithArgumentList(QLatin1String("{{.Name}}"), argumentList);
	    call.waitForFinished();
	    if (call.isValid()) {
		    QList<QVariant> args = call.reply().arguments();
		    switch (args.size()) {
			    case 0: return QVariant();
			    case 1: {
				    return unmarsh(args[0]);
			    }
		    default:
			    {
				    for (int i=0; i<args.size(); i++) {
					    args[i] = unmarsh(args[i]);
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
    void ___{{Lower .Name}}Changed___(QVariant);{{end}}

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
        void registerTypes(const char* uri) { {{range .Interfaces}}
             qmlRegisterType<{{.ObjectName}}>(uri, 1, 0, "{{.ObjectName}}");{{end}}
        }
};
` + _templateMarshUnMarsh + `
#endif
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


test.depends = lib/$(TARGET)
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
	pkgName := getQMLPkgName("DBus." + INFOS.Config.DestName)
	os.MkdirAll(INFOS.Config.OutputDir+"/"+strings.Replace(pkgName, ".", "/", -1), 0755)
	cmd_str := fmt.Sprintf("cd %s && ln -sv %s lib && qmake", INFOS.Config.OutputDir, strings.Replace(pkgName, ".", "/", -1))
	cmd := exec.Command("bash", "-c", cmd_str)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Run: " + cmd_str + " failed(Did you have an valid qmake?) testQML code will not generated!")
	}
	qmldir, err := os.Create(path.Join(INFOS.Config.OutputDir, "lib", "qmldir"))
	if err != nil {
		panic(err)
	}
	qmldir.WriteString("module " + pkgName + "\n")
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
		"PkgName":          func() string { return pkgName },
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

var sigToQtType = map[string]string{
	"y": "uchar",
	"b": "bool",
	"n": "short",
	"q": "ushort",
	"i": "int",
	"u": "uint",
	"x": "qlonglong",
	"t": "qulonglong",
	"d": "double",
	"s": "QString",
	"g": "QDBusVariant",
	"o": "QDBusObjectPath",
	"v": "QDBusSignature",
	"h": "uint",
}

func getQtSignaturesType() (sigs map[string]string) {
	sigs = make(map[string]string)
	var store func(string)
	store = func(sig string) {
		if v, ok := sigToQtType[sig]; ok {
			sigs[sig] = v
		} else if sig == "as" {
			sigs[sig] = "QStringList"
		} else if sig == "so" {
			fmt.Println("Warning: `so` isn't supported")
			sigs[sig] = "QStringList"
		} else if sig[0] == 'a' {
			if sig[1] == '(' {
				sigs[sig] = "QVariantList"
			} else if sig[1] == '{' {
				if len(sig) < 5 { // a{xx} has at least five characters
					return
				}
				store(sig[2:3])
				store(sig[3:strings.LastIndex(sig, "}")])
				sigs[sig] = "QVariantMap"
			} else {
				store(sig[1:])
				sigs[sig] = "QList< " + sigs[sig[1:]] + " >"
			}
		} else if sig[0] == '(' {
			sigs[sig] = "QVariantList"
		} else {
			panic(fmt.Sprintf("parse signature failed:%q\n", sig))
		}
	}
	for _, ifc := range INFOS.Interfaces {
		info := GetInterfaceInfo(ifc)
		for _, m := range info.Methods {
			for _, a := range m.Args {
				store(a.Type)
			}
		}
		for _, p := range info.Properties {
			store(p.Type)
		}
		for _, s := range info.Signals {
			for _, ss := range s.Args {
				store(ss.Type)
			}
		}
	}
	return sigs
}

var _templateMarshUnMarsh = `
inline
int getTypeId(const QString& sig) {
    //TODO: this should staticly generate by xml info
    if (0) { {{ range $key, $value := GetQtSignaturesType }}
    } else if (sig == "{{$key}}") {
	    return qDBusRegisterMetaType<{{$value}} >();{{end}}
    } else {
	    qDebug() << "Didn't support getTypeId" << sig << " please report it to snyh@snyh.org";
    }
}

inline
QVariant qstring2dbus(QString value, char sig) {
    switch (sig) {
        case 'y':
            return QVariant::fromValue(uchar(value[0].toLatin1()));
        case 'n':
            return QVariant::fromValue(value.toShort());
        case 'q':
            return QVariant::fromValue(value.toUShort());
        case 'i':
            return QVariant::fromValue(value.toInt());
        case 'u':
            return QVariant::fromValue(value.toUInt());
        case 'x':
            return QVariant::fromValue(value.toLongLong());
        case 't':
            return QVariant::fromValue(value.toULongLong());
        case 'd':
            return QVariant::fromValue(value.toDouble());
        case 's':
            return QVariant::fromValue(value);
        case 'o':
            return QVariant::fromValue(QDBusObjectPath(value));
        case 'v':
            return QVariant::fromValue(QDBusSignature(value));
        default:
            qDebug() << "Dict entry key should be an basic dbus type not an " << sig;
            return QVariant();
    }
}

QList<QString> splitStructureSignature(const QString& sig) {
    if (sig.size() < 3 || sig[0] != '(' || sig[sig.size()-1] != ')') {
        return QList<QString>();
    }

    QList<QString> sigs;

    QString tmp = sig.mid(1, sig.size()-2);
    while (tmp.size() != 0) {
        switch (tmp[0].toLatin1()) {
            case 'a':
                if (tmp.size() < 2) {
                    return QList<QString>();
                }
                if (tmp[1] == '{') {
                    int lastIndex = tmp.lastIndexOf('}') + 1;
                    if (lastIndex == 0) return QList<QString>();
                    sigs.append(tmp.mid(0, lastIndex));
                    tmp = tmp.mid(lastIndex);
                    break;
                } else if (tmp[1] == '(') {
                    int lastIndex = tmp.lastIndexOf(')') + 1;
                    if (lastIndex == 0) return QList<QString>();
                    sigs.append(tmp.mid(0, lastIndex));
                    tmp = tmp.mid(lastIndex);
                    break;
                } else {
                    sigs.append(tmp.mid(0, 2));
                    tmp = tmp.mid(2);
                    break;
                }
            case '(': {
                          int lastIndex = tmp.lastIndexOf(')') + 1;
                          if (lastIndex == 0) return QList<QString>();
                          sigs.append(tmp.mid(0, lastIndex));
                          tmp = tmp.mid(lastIndex);
                          break;
                      }
            case 'y': case 'b': case 'n': case 'q':
            case 'i': case 'u': case 'x': case 't':
            case 'd': case 's': case 'o': case 'g':
            case 'h': case 'v':
                sigs.append(QString(tmp[0]));
                tmp = tmp.mid(1, tmp.size() - 1);
                break;
            default:
                return QList<QString>();
        }
    }
    return sigs;
}

QVariant marsh(QDBusArgument target, const QVariant& arg, const QString& sig) {
    if (sig.size() == 0) {
        return QVariant::fromValue(target);
    }
    switch (sig[0].toLatin1()) {
        case 'y':
            target << qstring2dbus(arg.value<QString>(), 'y').value<uchar>();
            return QVariant::fromValue(target);
        case 'b':
            target << arg.value<bool>();
            return QVariant::fromValue(target);
        case 'n':
            target << arg.value<short>();
            return QVariant::fromValue(target);
        case 'q':
            target << arg.value<ushort>();
            return QVariant::fromValue(target);
        case 'i':
            target << arg.value<qint32>();
            return QVariant::fromValue(target);
        case 'u':
            target << arg.value<quint32>();
            return QVariant::fromValue(target);
        case 'x':
            target << arg.value<qlonglong>();
            return QVariant::fromValue(target);
        case 't':
            target << arg.value<qulonglong>();
            return QVariant::fromValue(target);
        case 'd':
            target << arg.value<double>();
            return QVariant::fromValue(target);
        case 's':
            target << arg.value<QString>();
            return QVariant::fromValue(target);
        case 'o':
            target << QDBusObjectPath(arg.value<QString>());
            return QVariant::fromValue(target);
        case 'g':
            target << QDBusSignature(arg.value<QString>());
            return QVariant::fromValue(target);

        case 'a':
            {
                if (sig.size() < 2) { return QVariant(); }
                char s = sig[1].toLatin1();
                if (s == '{') {
                    char key_sig = sig[2].toLatin1();
                    QString value_sig = sig.mid(3, sig.lastIndexOf('}') - 3);
                    target.beginMap(getTypeId(QString(key_sig)), getTypeId(value_sig));
                    //qDebug() << "BeginMap:" << key_sig << value_sig;
                    foreach(const QString& key, arg.value<QVariantMap>().keys()) {
                        //qDebug() << "KEY:" << key;
                        target.beginMapEntry();
                        //qDebug() <<"beginMapEntry";
                        marsh(target, qstring2dbus(key, key_sig), QString(key_sig));
                        marsh(target, arg.value<QVariantMap>()[key], value_sig);
                        //qDebug() <<"EndMapEntry";
                        target.endMapEntry();
                    }
                    //qDebug() << "EndMap";
                    target.endMap();
                    return QVariant::fromValue(target);
                } else {
                    QString next = sig.right(sig.size() - 1);
                    target.beginArray(getTypeId(next));
                    foreach(const QVariant& v, arg.value<QVariantList>()) {
                        marsh(target, v, next);
                    }
                    target.endArray();
                    return QVariant::fromValue(target);
                }
            }
        case '(':
            {
                QList<QString> sigs = splitStructureSignature(sig);
                QVariantList values = arg.value<QVariantList>();
                if (values.size() != sigs.size()) {
                    qDebug() << "structure (" << arg << ") didn't match signature :" << sigs;
                    return QVariant();
                }
                target.beginStructure();
                for (int i=0; i < sigs.size(); i++) {
                    marsh(target, values[i], sigs[i]);
                }
                target.endStructure();
                return QVariant::fromValue(target);
            }
        default:
            qDebug() << "Panic didn't support marsh" << sig;
    }
    return QVariant::fromValue(target);
}

inline
QVariant unmarshDBus(const QDBusArgument &argument)
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
            return unmarshDBus(v.value<QDBusArgument>());
        else
            return v;
    }
    case QDBusArgument::ArrayType: {
        QVariantList list;
        argument.beginArray();
        while (!argument.atEnd())
            list.append(unmarshDBus(argument));
        argument.endArray();
        return list;
    }
    case QDBusArgument::StructureType: {
        QVariantList list;
        argument.beginStructure();
        while (!argument.atEnd())
            list.append(unmarshDBus(argument));
        argument.endStructure();
        return QVariant::fromValue(list);
    }
    case QDBusArgument::MapType: {
        QVariantMap map;
        argument.beginMap();
        while (!argument.atEnd()) {
            argument.beginMapEntry();
            QVariant key = unmarshDBus(argument);
            QVariant value = unmarshDBus(argument);
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

QVariant unmarsh(const QVariant& v) {
    if (v.userType() == qMetaTypeId<QDBusObjectPath>()) {
        return QVariant::fromValue(v.value<QDBusObjectPath>().path());
    } else if (v.userType() == qMetaTypeId<QDBusArgument>()) {
        return unmarsh(unmarshDBus(v.value<QDBusArgument>()));
    } else if (v.userType() == qMetaTypeId<QByteArray>()) {
        return QString(v.value<QByteArray>());
    }
    return v;
}
`
