// +build ignore
package main

import "testing"

func TestQType(t *testing.T) {
	if getQType("u") != "uint" {
		t.Fatal(` "u" != "uint" `)
	}
	if getQType("au") != "QList<uint >" {
		t.Fatal(` "au" != "QList<uint >" ` + getQType("au"))
	}
	if getQType("ao") != "QList<QDBusObjectPath >" {
		t.Fatal(` "u" != "QList<QDBusObjectPath >" `)
	}
	if getQType("as") != "QList<QString >" {
		t.Fatal(` "u" != "QList<QString>" `)
	}
	if getQType("av") != "QList<QDBusVariant >" {
		t.Fatal(` "u" != "QList<QVariant>" `)
	}
	if getQType("a{ss}") != "QMap<QString, QString >" {
		t.Fatal(` "u" != "QMap<QString, QString >" `)
	}
}
