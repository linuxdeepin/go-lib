// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package gsprop

import (
	"encoding/xml"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	gio "github.com/linuxdeepin/go-gir/gio-2.0"
	"github.com/linuxdeepin/go-lib/dbusutilv1"
	"github.com/linuxdeepin/go-lib/gsettings"
)

const demoSchemaId = "ca.desrt.dconf-editor.Demo"

type srvObject1 struct {
	Bool   Bool   `prop:"access:rw"`
	Enum   Enum   `prop:"access:rw"`
	Int    Int    `prop:"access:rw"`
	Uint   Uint   `prop:"access:rw"`
	Double Double `prop:"access:rw"`
	String String `prop:"access:rw"`
	Strv   Strv   `prop:"access:rw"`
}

const srvObj1Interface = "org.deepin.dde.lib.gsprop.Object1"

func (o *srvObject1) GetExportedMethods() dbusutilv1.ExportedMethods {
	return nil
}

func TestAll(t *testing.T) {
	err := exec.Command("gsettings", "list-recursively", demoSchemaId).Run()
	if err != nil {
		t.Skip("failed to exec gsettings list-recursively")
	}

	script := `
gsettings set ca.desrt.dconf-editor.Demo boolean true
gsettings set ca.desrt.dconf-editor.Demo enumeration Blue
gsettings set ca.desrt.dconf-editor.Demo integer-32-signed -32
gsettings set ca.desrt.dconf-editor.Demo integer-32-unsigned 132
gsettings set ca.desrt.dconf-editor.Demo double 1.5
gsettings set ca.desrt.dconf-editor.Demo string 'hello world'
gsettings set ca.desrt.dconf-editor.Demo string-array '["go","perl","python", "c#"]'
`
	sh := exec.Command("sh")
	sh.Stdin = strings.NewReader(script)
	err = sh.Run()
	if err != nil {
		t.Error("Unexpected error executing script:", err)
	}

	gs := gio.NewSettings(demoSchemaId)

	srvObj1 := &srvObject1{}
	srvObj1.Bool.Bind(gs, "boolean")
	srvObj1.Enum.Bind(gs, "enumeration")
	srvObj1.Int.Bind(gs, "integer-32-signed")
	srvObj1.Uint.Bind(gs, "integer-32-unsigned")
	srvObj1.Double.Bind(gs, "double")
	srvObj1.String.Bind(gs, "string")
	srvObj1.Strv.Bind(gs, "string-array")

	service, err := dbusutilv1.NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	const srvObj1Path = "/org/deepin/dde/lib/gsprop/Object1"
	err = service.Export(srvObj1Path, srvObj1Interface, srvObj1)
	if err != nil {
		t.Error("Unexpected error export srvObj1:", err)
	}

	err = service.RequestName(srvObj1Interface)
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj1 := service.Conn().Object(srvObj1Interface, srvObj1Path)
	var introspection string
	err = clientObj1.Call(
		"org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&introspection)
	if err != nil {
		t.Error("Unexpected error calling Introspectable method:", err)
	}

	var iNode introspect.Node
	err = xml.Unmarshal([]byte(introspection), &iNode)
	if err != nil {
		t.Error("Unexpected error unmarshaling xml:", err)
	}
	xmlData, err := xml.MarshalIndent(&iNode, "", "\t")
	if err != nil {
		t.Error("Unexpected marshal indent:", xmlData)
	}
	t.Logf("%s\n", xmlData)

	var pBool bool
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Bool").Store(&pBool)
	if err != nil {
		t.Error("Unexpected error getting the Bool property:", err)
	}

	if pBool != true {
		t.Error("pBool expected true")
	}

	var pEnum uint32
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Enum").Store(&pEnum)
	if err != nil {
		t.Error("Unexpected error getting the Enum property:", err)
	}
	if pEnum != 2 {
		t.Errorf("pEnum expected 2 but got %d", pEnum)
	}

	var pInt int32
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Int").Store(&pInt)
	if err != nil {
		t.Error("Unexpected error getting the Int property:", err)
	}
	if pInt != -32 {
		t.Errorf("pInt expected -32 but got %d", pInt)
	}

	var pUint uint32
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Uint").Store(&pUint)
	if err != nil {
		t.Error("Unexpected error getting the Uint property:", err)
	}
	if pUint != 132 {
		t.Errorf("pUint expected 132 but got %d", pUint)
	}

	var pDouble float64
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Double").Store(&pDouble)
	if err != nil {
		t.Error("Unexpected error getting the Double property:", err)
	}
	if pDouble != 1.5 {
		t.Errorf("pDobule expected 1.5 but got %v", pDouble)
	}

	var pString string
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "String").Store(&pString)
	if err != nil {
		t.Error("Unexpected error getting the String property:", err)
	}
	pStringExpected := "hello world"
	if pString != pStringExpected {
		t.Errorf("pString expected %q but got %q", pStringExpected, pString)
	}

	var pStrv []string
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Strv").Store(&pStrv)
	if err != nil {
		t.Error("Unexpected error getting the Strv property:", err)
	}

	pStrvExpected := []string{"go", "perl", "python", "c#"}
	if !strvEqual(pStrvExpected, pStrv) {
		t.Errorf("pStrv expected %#v but got %#v", pStrvExpected, pStrv)
	}

	// GetAll
	var props map[string]dbus.Variant
	err = clientObj1.Call("org.freedesktop.DBus.Properties.GetAll", 0,
		srvObj1Interface).Store(&props)
	if err != nil {
		t.Error("Unexpected error getting all properties:", err)
	}
	t.Logf("props: %#v\n", props)

	// set bool
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Bool", dbus.MakeVariant(false)).Err
	if err != nil {
		t.Error("Unexpected error setting the Bool property:", err)
	}
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Bool").Store(&pBool)
	if err != nil {
		t.Error("Unexpected error getting the Bool property:", err)
	}

	if pBool != false {
		t.Error("pBool expected false")
	}

	// set enum
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Enum", dbus.MakeVariant(int32(1))).Err
	if err != nil {
		t.Error("Unexpected error setting the Enum property:", err)
	}

	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Enum").Store(&pEnum)
	if err != nil {
		t.Error("Unexpected error getting the Enum property:", err)
	}
	if pEnum != 1 {
		t.Errorf("pEnum expected 1 but got %d", pEnum)
	}

	// set enum failed
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Enum", dbus.MakeVariant(int32(999))).Err
	if err == nil {
		t.Error("Expected error setting the Enum property")
	}

	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Enum").Store(&pEnum)
	if err != nil {
		t.Error("Unexpected error getting the Enum property:", err)
	}
	if pEnum != 1 {
		t.Errorf("pEnum expected 1 but got %d", pEnum)
	}

	// set int
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Int", dbus.MakeVariant(int32(110))).Err
	if err != nil {
		t.Error("Unexpected error setting the Int property:", err)
	}

	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Int").Store(&pInt)
	if err != nil {
		t.Error("Unexpected error getting the Int property:", err)
	}
	if pInt != 110 {
		t.Errorf("pInt expected 110 but got %d", pInt)
	}

	// set uint
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Uint", dbus.MakeVariant(uint32(120))).Err
	if err != nil {
		t.Error("Unexpected error setting the Uint property:", err)
	}

	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Uint").Store(&pUint)
	if err != nil {
		t.Error("Unexpected error getting the Uint property:", err)
	}
	if pUint != 120 {
		t.Errorf("pUint expected 120 but got %d", pUint)
	}

	// set double
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Double", dbus.MakeVariant(float64(3.5))).Err
	if err != nil {
		t.Error("Unexpected error setting the Double property:", err)
	}

	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Double").Store(&pDouble)
	if err != nil {
		t.Error("Unexpected error getting the Double property:", err)
	}
	if pDouble != 3.5 {
		t.Errorf("pDouble expected 3.5 but got %v", pDouble)
	}

	// set string
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "String", dbus.MakeVariant("deepin")).Err
	if err != nil {
		t.Error("Unexpected error setting the String property:", err)
	}

	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "String").Store(&pString)
	if err != nil {
		t.Error("Unexpected error getting the String property:", err)
	}
	if pString != "deepin" {
		t.Errorf("pString expected %q but got %q", "deepin", pString)
	}

	// set strv
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Strv", dbus.MakeVariant([]string{"a", "b", "c"})).Err
	if err != nil {
		t.Error("Unexpected error setting the Strv property:", err)
	}

	pStrv = nil
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Get", 0,
		srvObj1Interface, "Strv").Store(&pStrv)
	if err != nil {
		t.Error("Unexpected error getting the Strv property:", err)
	}
	pStrvExpected = []string{"a", "b", "c"}
	if !strvEqual(pStrv, pStrvExpected) {
		t.Errorf("pStrv expected %#v but got %#v", pStrvExpected, pStrv)
	}
	gio.SettingsSync()

	// property changed signal
	go func() {
		_ = gsettings.StartMonitor()
	}()

	rule := dbusutilv1.NewMatchRuleBuilder().ExtPropertiesChanged(srvObj1Path,
		srvObj1Interface).Build()

	err = rule.AddTo(service.Conn())
	if err != nil {
		t.Error("Unexpected error adding rule to service conn:", err)
	}

	ch := make(chan int)
	go processSignal(service.Conn(), func(signal *dbus.Signal) bool {
		if signal.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" &&
			signal.Path == srvObj1Path {

			ch <- 1

			interfaceName := signal.Body[0].(string)
			if interfaceName != srvObj1Interface {
				t.Errorf("interfaceName expected %q but got %q", srvObj1Interface,
					interfaceName)
			}

			propVarMap := signal.Body[1].(map[string]dbus.Variant)
			variant, ok := propVarMap["Bool"]
			if !ok {
				t.Error("propVarMap expected has key Prop4")
			}
			val := variant.Value().(bool)
			if val != true {
				t.Errorf("val expected true")
			}
			return false
		}
		return true
	})

	// set bool
	err = clientObj1.Call("org.freedesktop.DBus.Properties.Set", 0,
		srvObj1Interface, "Bool", dbus.MakeVariant(true)).Err
	if err != nil {
		t.Error("Unexpected error setting the Bool property:", err)
	}

	select {
	case <-ch:
		close(ch)
	case <-time.After(5 * time.Second):
		t.Log("Failed to announce that the Bool PropertyChanged signal emitted")
	}

	err = rule.RemoveFrom(service.Conn())
	if err != nil {
		t.Error("Unexpected error removing rule from service conn:", err)
	}

}

// copy from dbusutilv1 service_test.go
func processSignal(conn *dbus.Conn, fn func(signal *dbus.Signal) bool) {
	signalChan := make(chan *dbus.Signal, 10)
	conn.Signal(signalChan)

	for sig := range signalChan {
		shouldContinue := fn(sig)
		if !shouldContinue {
			break
		}
	}
	conn.RemoveSignal(signalChan)
	close(signalChan)
}
