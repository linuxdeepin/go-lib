// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusutil

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
)

func isSessionBusExists() bool {
	address := fmt.Sprintf("/run/user/%d/bus", os.Getuid())
	_, err := os.Stat(address)
	if err != nil {
		return false
	}

	return true
}

func TestService_GetNameOwner(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	owner, err := service.GetNameOwner(orgFreedesktopDBus)
	if err != nil {
		t.Error("Unexpected error calling GetNameOwner:", err)
	}
	if owner != orgFreedesktopDBus {
		t.Errorf("expected owner %q got %q", orgFreedesktopDBus, owner)
	}

	_, err = service.GetNameOwner("xxx.yyy.zzz.123")
	if err == nil {
		t.Error("Expected error due to service xxx.yyy.zzz.123 not exist")
	}
}

func TestService_NameHasOwner(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	hasOwner, err := service.NameHasOwner(orgFreedesktopDBus)
	if err != nil {
		t.Error("Unexpected error calling NameHasOwner:", err)
	}
	if !hasOwner {
		t.Error("hasOwner expected true")
	}

	hasOwner, err = service.NameHasOwner("xxx.yyy.zzz.123")
	if err != nil {
		t.Error("Unexpected error calling NameHasOwner:", err)
	}
	if hasOwner {
		t.Error("hasOwner expected false")
	}
}

func TestService_RequestName(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	const name = "org.deepin.dde.lib.RequestName"
	err = service.RequestName(name)
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	err = service.ReleaseName(name)
	if err != nil {
		t.Error("Unexpected error calling ReleaseName:", err)
	}
}

type srvObject1 struct {
}

func (o *srvObject1) GetExportedMethods() ExportedMethods {
	return ExportedMethods{
		{
			Name:    "Method1",
			Fn:      o.Method1,
			OutArgs: []string{"num"},
		},
	}
}

func (*srvObject1) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object1"
}

func (*srvObject1) Method1() (int, *dbus.Error) {
	return 1, nil
}

type srvObject12 struct{}

func (o srvObject12) GetExportedMethods() ExportedMethods {
	return nil
}

func (o srvObject12) GetMethodTable() map[string]interface{} {
	return nil
}

func (srvObject12) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object12"
}

type srvString string

func (srvString) GetExportedMethods() ExportedMethods {
	return nil
}

func (srvString) GetInterfaceName() string {
	return "org.deepin.dde.lib.String"
}

func TestService_Export(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	srvObj1 := &srvObject1{}
	err = service.Export("/org/deepin/dde/lib/Object1", srvObj1)
	if err != nil {
		t.Error("Unexpected error exporting srvObj1:", err)
	}

	err = service.RequestName("org.deepin.dde.lib.Object1")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj1 := service.conn.Object("org.deepin.dde.lib.Object1", "/org/deepin/dde/lib/Object1")

	var num int
	err = clientObj1.Call("org.deepin.dde.lib.Object1.Method1", 0).Store(&num)
	if err != nil {
		t.Error("Unexpected error calling srvObj1.Method1:", err)
	}
	if num != 1 {
		t.Errorf("expect 1 but got %d", num)
	}

	if !service.IsExported(srvObj1) {
		t.Error("IsExported expected true")
	}

	_ = service.StopExport(srvObj1)

	if service.IsExported(srvObj1) {
		t.Error("IsExported expected false")
	}

	srvObj12 := srvObject12{}
	err = service.Export("/org/deepin/dde/lib/Object12", srvObj12)
	if err == nil {
		t.Error("Expected error due to srvObj12 is not a struct pointer")
	}

	srvStr := srvString("hello")
	err = service.Export("/org/deepin/dde/lib/String", srvStr)
	if err == nil {
		t.Error("Expected error due to srvStr is not a struct pointer")
	}
}

type srvObject2 struct {
	//nolint
	signals *struct {
		Signal1 struct{}
		Signal2 struct {
			Arg1 string
			Arg2 uint32
		}
	}
}

func (*srvObject2) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject2) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object2"
}

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

func TestService_Emit(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	srvObj2 := &srvObject2{}
	const srvObj2Path = "/org/deepin/dde/lib/Object2"
	err = service.Export(srvObj2Path, srvObj2)
	if err != nil {
		t.Error("Unexpected error exporting srvObj2:", err)
	}

	// Test Signal1
	// rule for watch Signal1
	ruleSig1 := NewMatchRuleBuilder().ExtSignal(srvObj2Path, srvObj2.GetInterfaceName(), "Signal1").Build()
	err = ruleSig1.AddTo(service.conn)
	if err != nil {
		t.Error("unaxpected error adding match rule:", err)
	}

	ch1 := make(chan int)
	go processSignal(service.conn, func(sig *dbus.Signal) bool {
		if sig.Name == "org.deepin.dde.lib.Object2.Signal1" &&
			sig.Path == "/org/deepin/dde/lib/Object2" {
			ch1 <- 1

			if len(sig.Body) != 0 {
				t.Errorf("len(sig.Body) expected 0 got %d", len(sig.Body))
			}
			return false
		}
		return true
	})

	err = service.Emit(srvObj2, "Signal1")
	if err != nil {
		t.Error("Unexpected error emitting Sig1:", err)
	}

	select {
	case <-ch1:
		close(ch1)
	case <-time.After(30 * time.Second):
		t.Log("Failed to announce that the Signal1 was emitted")
	}

	err = ruleSig1.RemoveFrom(service.conn)
	if err != nil {
		t.Error("Unexpected error removing match rule:", err)
	}

	// Test Signal2
	// rule for watch Signal2
	ruleSig2 := NewMatchRuleBuilder().ExtSignal(srvObj2Path, srvObj2.GetInterfaceName(), "Signal2").Build()
	err = ruleSig2.AddTo(service.conn)
	if err != nil {
		t.Error("unaxpected error adding match rule:", err)
	}

	ch2 := make(chan int)

	go processSignal(service.conn, func(sig *dbus.Signal) bool {
		if sig.Name == "org.deepin.dde.lib.Object2.Signal2" &&
			sig.Path == "/org/deepin/dde/lib/Object2" {
			ch2 <- 1

			expectedBody := []interface{}{"hello", uint32(1)}
			if !reflect.DeepEqual(sig.Body, expectedBody) {
				t.Errorf("sig.Body expected %#v got %#v", expectedBody, sig.Body)
			}
			return false
		}
		return true
	})

	err = service.Emit(srvObj2, "Signal2", "hello", uint32(1))
	if err != nil {
		t.Error("Unexpected error emitting Sig2:", err)
	}

	select {
	case <-ch2:
		close(ch2)
	case <-time.After(30 * time.Second):
		t.Log("Failed to announce that the Signal2 was emitted")
	}

	err = ruleSig2.RemoveFrom(service.conn)
	if err != nil {
		t.Error("Unexpected error removing match rule:", err)
	}

	err = service.Emit(srvObj2, "Signal3")
	if err == nil {
		t.Error("Expected error due to srvObj2 don't have signal Signal3")
	}

	err = service.Emit(srvObj2, "Signal2")
	if err == nil {
		t.Error("Expected error due to Signal2 signature not match")
	}

	err = service.Emit(srvObj2, "Signal2", byte(1), uint32(1))
	if err == nil {
		t.Error("Expected error due to Signal2 signature not match")
	}
}

type srvObject3 struct {
	PropsMu sync.RWMutex
	Prop1   string
	Prop2   uint32
}

func (*srvObject3) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject3) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object3"
}

var serviceEmitPropertyChangedTests = []struct {
	PropName string
	Value    interface{}
	Err      bool
	FailMsg  string
}{
	{"Prop1", "abc", false,
		"Unexpected error emitting property Prop1 changed"},

	{"Prop2", uint32(1), false,
		"Unexpected error emitting property Prop2 changed"},

	{"Prop3", "", true,
		"Expected error due to property Prop2 doest not exist"},

	{"Prop1", 0, true,
		"Expected error due to type of Prop1 is not int"},

	{"Prop2", "abc", true,
		"Expected error due to type of Prop2 is not string"},
}

func TestService_EmitPropertyChanged(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	srvObj3 := &srvObject3{}
	const srvObj3Path = "/org/deepin/dde/lib/Object3"
	err = service.Export(srvObj3Path, srvObj3)
	if err != nil {
		t.Error("Unexpected error exporting srvObj3:", err)
	}

	// rule for watch PropertiesChanged signal
	rule := NewMatchRuleBuilder().ExtPropertiesChanged(srvObj3Path,
		srvObj3.GetInterfaceName()).Build()
	err = rule.AddTo(service.conn)
	if err != nil {
		t.Error("Unexpected error adding match rule:", err)
	}
	for _, testCase := range serviceEmitPropertyChangedTests {
		ch := make(chan int)
		if !testCase.Err {
			go processSignal(service.conn, func(sig *dbus.Signal) bool {
				propName := testCase.PropName
				if sig.Path == "/org/deepin/dde/lib/Object3" &&
					sig.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" {
					ch <- 1

					propVarMap := sig.Body[1].(map[string]dbus.Variant)
					_, ok := propVarMap[propName]
					if !ok {
						t.Error("propVarMap expected has key:", propName)
					}
					return false
				}
				return true
			})
		}

		err = service.EmitPropertyChanged(srvObj3, testCase.PropName, testCase.Value)
		if testCase.Err {
			if err == nil {
				t.Error(testCase)
			}
		} else {
			if err != nil {
				t.Error(testCase)
			}

			select {
			case <-ch:
				close(ch)
			case <-time.After(5 * time.Second):
				t.Log("Failed to announce that the PropertiesChanged signal was emitted")
			}
		}
	}

	err = rule.RemoveFrom(service.conn)
	if err != nil {
		t.Error("Unexpected error removing match rule:", err)
	}
}

type srvObject4 struct {
	PropsMu sync.RWMutex
	Prop1   string
	Prop2   uint32
}

func (*srvObject4) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject4) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object4"
}

var serviceEmitPropertiesChangedTest = []struct {
	propValueMap    map[string]interface{}
	invalidateProps []string
	Err             bool
	FailMsg         string
}{
	{
		map[string]interface{}{
			"Prop1": "abc",
			"Prop2": uint32(0),
		},
		nil,
		false,
		"Unexpected error emitting properties changed",
	},

	{
		map[string]interface{}{
			"Prop1": "abc",
		},
		[]string{"Prop2"},
		false,
		"Unexpected error emitting properties changed",
	},

	{
		nil, nil, false,
		"Unexpected error emitting empty properties changed",
	},

	{
		map[string]interface{}{
			"Prop1": "abc",
			"Prop3": 1,
		},
		[]string{"Prop2"},
		true,
		"Expected error due to property Prop3 does not exist",
	},

	{
		map[string]interface{}{
			"Prop1": "abc",
		},
		[]string{"Prop2", "Prop4"},
		true,
		"Expected error due to property Prop4 does not exist",
	},

	{
		map[string]interface{}{
			"Prop1": uint32(1),
		},
		[]string{"Prop2"},
		true,
		"Expected error due to type of property Prop1 is not uint32",
	},

	{
		map[string]interface{}{
			"Prop1": uint32(1),
		},
		[]string{"Prop1"},
		true,
		"Expected error due to Prop1 appears in both propValueMap and invalidateProps",
	},
}

func TestService_EmitPropertiesChanged(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	srvObj4 := &srvObject4{}
	err = service.Export("/org/deepin/dde/lib/Object4", srvObj4)
	if err != nil {
		t.Error("Unexpected error exporting srvObj4:", err)
	}

	for _, testCase := range serviceEmitPropertiesChangedTest {
		err = service.EmitPropertiesChanged(srvObj4, testCase.propValueMap, testCase.invalidateProps...)

		if testCase.Err {
			if err == nil {
				t.Error(testCase)
			}
		} else {
			if err != nil {
				t.Error(testCase)
			}
		}
	}
}

type srvObject5 struct {
	s *Service
}

func (obj *srvObject5) GetExportedMethods() ExportedMethods {
	return ExportedMethods{
		{
			Name: "Method1",
			Fn:   obj.Method1,
		},
	}
}

func (*srvObject5) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object5"
}

func (obj *srvObject5) Method1() *dbus.Error {
	obj.s.DelayAutoQuit()
	return nil
}

func TestService_AutoQuit(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	ch := make(chan struct{})

	bus, err := dbus.SessionBusPrivate()
	if err != nil {
		t.Error("Unexpected error connecting to session bus:", err)
	}

	err = bus.Auth(nil)
	if err != nil {
		t.Error("Unexpected error auth for bus:", err)
	}
	err = bus.Hello()
	if err != nil {
		t.Error("Unexpected error say Hello:", err)
	}

	service := NewService(bus)
	srvObj5 := &srvObject5{
		s: service,
	}
	err = service.Export("/org/deepin/dde/lib/Object5", srvObj5)
	if err != nil {
		t.Error("Unexpected error exporting srvObj5:", err)
	}

	err = service.RequestName("org.deepin.dde.lib.Object5")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	service.SetAutoQuitHandler(100*time.Millisecond, func() bool {
		return true
	})

	go func() {
		service.Wait()
		ch <- struct{}{}
	}()

	go func() {
		clientObj5 := service.conn.Object("org.deepin.dde.lib.Object5",
			"/org/deepin/dde/lib/Object5")

		for i := 0; i < 5; i++ {
			err = clientObj5.Call("org.deepin.dde.lib.Object5.Method1", 0).Err
			if err != nil {
				t.Error("Unexpected error calling srvObject5.Method1")
			}
			t.Log("call Method1")
		}

	}()

	select {
	case <-ch:
		// ok
		close(ch)
	case <-time.After(5 * time.Second):
		t.Log("Failed to announce that the service has quit")
	}
}

func TestService_GetConnPIDAndUID(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	names := service.conn.Names()
	uniqueName := names[0]

	t.Log(uniqueName)
	pid, err := service.GetConnPID(uniqueName)
	if err != nil {
		t.Error("Unexpected error calling GetConnPID:", err)
	}

	if int(pid) != os.Getpid() {
		t.Errorf("pid expected %d, got %d", os.Getpid(), pid)
	}

	uid, err := service.GetConnUID(uniqueName)
	if err != nil {
		t.Error("Unexpected error calling GetConnUID:", err)
	}
	if int(uid) != os.Getuid() {
		t.Errorf("uid expected %d, got %d", os.Getuid(), uid)
	}
}

type srvObject6 struct {
	Prop1          string
	prop1ReadCount int
}

func (*srvObject6) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject6) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object6"
}

func TestService_SetReadCallback(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	srvObj6 := &srvObject6{
		Prop1: "apple",
	}

	serverObject6, err := service.NewServerObject("/org/deepin/dde/lib/Object6", srvObj6)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	err = serverObject6.SetReadCallback(srvObj6, "Prop1",
		func(read *PropertyRead) *dbus.Error {
			srvObj6.prop1ReadCount++
			return nil
		})

	if err != nil {
		t.Error("Unexpected error set read callabck for Prop1:", err)
	}

	err = serverObject6.Export()
	if err != nil {
		t.Error("Unexpected error exporting srvObj6:", err)
	}

	err = service.RequestName("org.deepin.dde.lib.Object6")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj6 := service.conn.Object("org.deepin.dde.lib.Object6",
		"/org/deepin/dde/lib/Object6")

	var prop1Value interface{}
	err = clientObj6.Call(orgFreedesktopDBus+".Properties.Get", 0,
		"org.deepin.dde.lib.Object6", "Prop1").Store(&prop1Value)
	if err != nil {
		t.Error("Unexpected error getting Prop1 value")
	}

	if prop1ValStr, ok := prop1Value.(string); !ok {
		t.Error("prop1Value is not string")
	} else {
		if prop1ValStr != "apple" {
			t.Errorf("prop1ValStr expected %q, got %q", "abcdef", prop1ValStr)
		}
	}

	if srvObj6.prop1ReadCount != 1 {
		t.Errorf("prop1ReadCount expected 1, got %d", srvObj6.prop1ReadCount)
	}

	_ = serverObject6.SetReadCallback(srvObj6, "Prop1",
		func(read *PropertyRead) *dbus.Error {
			srvObj6.prop1ReadCount++
			return dbus.MakeFailedError(errors.New("xxx err msg"))
		})

	err = clientObj6.Call(orgFreedesktopDBus+".Properties.Get", 0,
		"org.deepin.dde.lib.Object6", "Prop1").Store(&prop1Value)
	if err == nil {
		t.Error("Expected error due to read callback return error")
	}
	if srvObj6.prop1ReadCount != 2 {
		t.Errorf("prop1ReadCount expected 2, got %d", srvObj6.prop1ReadCount)
	}
}

type srvObject7 struct {
	Prop1           string `prop:"access:rw"`
	prop1WriteCount int
}

func (*srvObject7) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject7) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object7"
}

func TestService_SetWriteCallback(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	srvObj7 := &srvObject7{
		Prop1: "apple",
	}

	serverObject7, err := service.NewServerObject("/org/deepin/dde/lib/Object7", srvObj7)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	err = serverObject7.SetWriteCallback(srvObj7, "Prop1",
		func(write *PropertyWrite) *dbus.Error {
			srvObj7.prop1WriteCount++
			return nil
		})

	if err != nil {
		t.Error("Unexpected error set write callabck for Prop1:", err)
	}

	err = serverObject7.Export()
	if err != nil {
		t.Error("Unexpected error exporting srvObj7:", err)
	}

	err = service.RequestName("org.deepin.dde.lib.Object7")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj7 := service.conn.Object("org.deepin.dde.lib.Object7",
		"/org/deepin/dde/lib/Object7")
	err = clientObj7.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object7", "Prop1", dbus.MakeVariant("orange")).Err
	if err != nil {
		t.Error("Unexpected error setting prop1 value:", err)
	}

	if srvObj7.prop1WriteCount != 1 {
		t.Errorf("prop1WriteCount expected 1, got %d", srvObj7.prop1WriteCount)
	}

	if srvObj7.Prop1 != "orange" {
		t.Errorf("prop1 expected %q, got %q", "orange", srvObj7.Prop1)
	}

	err = serverObject7.SetWriteCallback(srvObj7, "Prop1",
		func(write *PropertyWrite) *dbus.Error {
			srvObj7.prop1WriteCount++
			return dbus.MakeFailedError(errors.New("xxx err msg"))
		})

	if err != nil {
		t.Error("Unexpected error set write callabck for Prop1:", err)
	}

	err = clientObj7.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object7", "Prop1", dbus.MakeVariant("banana")).Err
	if err == nil {
		t.Error("Expected error due to write callback return error:")
	}

	if srvObj7.prop1WriteCount != 2 {
		t.Errorf("prop1WriteCount expected 1, got %d", srvObj7.prop1WriteCount)
	}

	if srvObj7.Prop1 != "orange" {
		t.Errorf("prop1 expected %q, got %q", "orange", srvObj7.Prop1)
	}
}

type srvObject8 struct {
	Prop1 string `prop:"access:rw"`

	changedCountA int
	changedCountB int
}

func (*srvObject8) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject8) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object8"
}

func TestService_ConnectChanged(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	srvObj8 := &srvObject8{
		Prop1: "apple",
	}

	serverObject8, err := service.NewServerObject("/org/deepin/dde/lib/Object8", srvObj8)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	err = serverObject8.ConnectChanged(srvObj8, "Prop1",
		func(change *PropertyChanged) {
			srvObj8.changedCountA++
		})

	if err != nil {
		t.Error("Unexpected error set write callback for Prop1:", err)
	}

	err = serverObject8.ConnectChanged(srvObj8, "Prop1",
		func(change *PropertyChanged) {
			srvObj8.changedCountB += 2
		})

	if err != nil {
		t.Error("Unexpected error set write callabck for Prop1:", err)
	}

	err = serverObject8.Export()
	if err != nil {
		t.Error("Unexpected error exporting srvObj8:", err)
	}

	err = service.RequestName("org.deepin.dde.lib.Object8")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj8 := service.conn.Object("org.deepin.dde.lib.Object8",
		"/org/deepin/dde/lib/Object8")
	err = clientObj8.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object8", "Prop1", dbus.MakeVariant("banana")).Err
	if err != nil {
		t.Error("Unexpected error setting prop1 value:", err)
	}

	if srvObj8.changedCountA != 1 {
		t.Errorf("changedCountA expected 1, got %d", srvObj8.changedCountA)
	}

	if srvObj8.changedCountB != 2 {
		t.Errorf("changedCountB expected 2, got %d", srvObj8.changedCountB)
	}

	if srvObj8.Prop1 != "banana" {
		t.Errorf("prop1 expected %q, got %q", "123abc", srvObj8.Prop1)
	}

	// Set the same value to Prop1 again
	err = clientObj8.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object8", "Prop1", dbus.MakeVariant("banana")).Err
	if err != nil {
		t.Error("Unexpected error setting prop1 value:", err)
	}
	if srvObj8.changedCountA != 1 {
		t.Errorf("changedCountA expected 1, got %d", srvObj8.changedCountA)
	}

	if srvObj8.changedCountB != 2 {
		t.Errorf("changedCountB expected 2, got %d", srvObj8.changedCountB)
	}
}

type srvObject9 struct {
	Prop0 int32 `prop:"-"`
	Prop1 int32 `prop:"access:rw"`
	Prop2 int32 `prop:"access:r"`
	Prop3 int32 `prop:"access:w"`

	Prop4 int32 `prop:"access:rw,emit:true"`
	Prop5 int32 `prop:"access:rw,emit:invalidates"`
	Prop6 int32 `prop:"access:rw,emit:false"`
}

func (*srvObject9) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject9) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object9"
}

func TestService_PropTag(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	srvObj9 := &srvObject9{
		Prop0: 11,
		Prop1: 1,
		Prop2: 2,
		Prop3: 3,
	}
	const srvObj9Path = "/org/deepin/dde/lib/Object9"

	err = service.Export(srvObj9Path, srvObj9)
	if err != nil {
		t.Error("Unexpected error exporting srvObj9:", err)
	}

	err = service.RequestName("org.deepin.dde.lib.Object9")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj9 := service.conn.Object("org.deepin.dde.lib.Object9",
		"/org/deepin/dde/lib/Object9")

	// prop0 ignored
	_, err = clientObj9.GetProperty("org.deepin.dde.lib.Object9.Prop0")
	if err == nil {
		t.Error("Expected error due to Prop0 should be ignored")
	} else {
		if !strings.Contains(err.Error(), "PropertyNotFound") {
			t.Error("Expected error is PropertyNotFound")
		}
	}

	// prop1 access rw - readwrite
	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object9", "Prop1", dbus.MakeVariant(int32(11))).Err
	if err != nil {
		t.Error("Unexpected error setting Prop1 value:", err)
	}

	if srvObj9.Prop1 != 11 {
		t.Errorf("prop1 expected 11 got %d", srvObj9.Prop1)
	}

	var prop1Value uint32
	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Get", 0,
		"org.deepin.dde.lib.Object9", "Prop1").Store(&prop1Value)
	if err != nil {
		t.Error("Unexpected error getting Prop1 value:", err)
	}

	if prop1Value != 11 {
		t.Errorf("prop1Value expected 11, got %d", prop1Value)
	}

	// prop2 access r - read only
	var prop2Value uint32
	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Get", 0,
		"org.deepin.dde.lib.Object9", "Prop2").Store(&prop2Value)
	if err != nil {
		t.Error("Unexpected error getting Prop2 value:", err)
	}

	if prop2Value != 2 {
		t.Errorf("prop2Value expected 2, got %d", prop2Value)
	}

	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object9", "Prop2", dbus.MakeVariant(int32(12))).Err
	if err == nil {
		t.Error("Expected error due to Prop2 read only")
	}

	// prop3 access w - write only
	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Get", 0,
		"org.deepin.dde.lib.Object9", "Prop3").Err
	if err == nil {
		t.Error("Expected error due to Prop2 write only")
	}

	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object9", "Prop3", dbus.MakeVariant(int32(13))).Err
	if err != nil {
		t.Error("Unexpected error setting Prop3 value:", err)
	}

	if srvObj9.Prop3 != 13 {
		t.Errorf("prop3 expected 13 got %d", srvObj9.Prop3)
	}

	// rule for watch PropertiesChanged signal
	rulePC := NewMatchRuleBuilder().ExtPropertiesChanged(srvObj9Path, srvObj9.GetInterfaceName()).Build()
	_ = rulePC.AddTo(service.conn)
	// Test Prop4 emit: true

	chP4 := make(chan int)
	go processSignal(service.conn, func(sig *dbus.Signal) bool {
		if sig.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" &&
			sig.Path == "/org/deepin/dde/lib/Object9" {
			chP4 <- 1

			interfaceName := sig.Body[0].(string)
			expectedInterface := "org.deepin.dde.lib.Object9"
			if interfaceName != expectedInterface {
				t.Errorf("interfaceName expected %q got %q", expectedInterface,
					interfaceName)
			}

			propVarMap := sig.Body[1].(map[string]dbus.Variant)
			variant, ok := propVarMap["Prop4"]
			if !ok {
				t.Error("propVarMap expected has key Prop4")
			}
			val := variant.Value().(int32)
			if val != 14 {
				t.Errorf("val expected 14 got %d", variant.Value())
			}
			return false
		}
		return true
	})

	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object9", "Prop4", dbus.MakeVariant(int32(14))).Err
	if err != nil {
		t.Error("Unexpected error setting Prop3 value:", err)
	}

	if srvObj9.Prop4 != 14 {
		t.Errorf("prop4 expected 4 got %d", srvObj9.Prop4)
	}

	select {
	case <-chP4:
		close(chP4)
	case <-time.After(5 * time.Second):
		t.Log("Failed to announce that the Prop4 PropertyChanged signal emitted")
	}

	// Test Prop5 emit: invalidates
	chP5 := make(chan int)
	go processSignal(service.conn, func(sig *dbus.Signal) bool {
		if sig.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" &&
			sig.Path == "/org/deepin/dde/lib/Object9" {
			chP5 <- 1

			interfaceName := sig.Body[0].(string)
			expectedInterface := "org.deepin.dde.lib.Object9"
			if interfaceName != expectedInterface {
				t.Errorf("interfaceName expected %q got %q", expectedInterface,
					interfaceName)
			}

			propVarMap := sig.Body[1].(map[string]dbus.Variant)
			if len(propVarMap) != 0 {
				t.Errorf("len(propVarMap) expected 0 got %d", len(propVarMap))
			}

			invalidatesProps := sig.Body[2].([]string)
			expectedProps := []string{"Prop5"}
			if !reflect.DeepEqual(invalidatesProps, expectedProps) {
				t.Errorf("invalidatesProps expected %#v got %#v",
					expectedProps, invalidatesProps)
			}
			return false
		}
		return true
	})

	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object9", "Prop5", dbus.MakeVariant(int32(15))).Err
	if err != nil {
		t.Error("Unexpected error setting Prop3 value:", err)
	}

	if srvObj9.Prop5 != 15 {
		t.Errorf("prop5 expected 15 got %d", srvObj9.Prop5)
	}

	select {
	case <-chP5:
		close(chP5)
	case <-time.After(5 * time.Second):
		t.Log("Failed to announce that the Prop5 PropertyChanged signal emitted")
	}

	// Test Prop6 emit: false
	chP6 := make(chan int)
	go processSignal(service.conn, func(sig *dbus.Signal) bool {
		if sig.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" &&
			sig.Path == "/org/deepin/dde/lib/Object9" {
			chP6 <- 1
			return false
		}
		return true
	})

	err = clientObj9.Call(orgFreedesktopDBus+".Properties.Set", 0,
		"org.deepin.dde.lib.Object9", "Prop6", dbus.MakeVariant(int32(16))).Err
	if err != nil {
		t.Error("Unexpected error setting Prop3 value:", err)
	}

	if srvObj9.Prop6 != 16 {
		t.Errorf("prop6 expected 16 got %d", srvObj9.Prop6)
	}

	select {
	case <-chP6:
		t.Error("Failed to announce that the Prop6 PropertyChanged signal was not emitted")
	case <-time.After(1 * time.Second):
		// ok
	}

	// Test End
	_ = rulePC.RemoveFrom(service.conn)
}

type rect struct {
	X, Y uint32
	W, H int32
}

type twoRect struct {
	A, B rect
}

type srvObject10 struct {
	Prop1             rect `prop:"access:rw"`
	prop1ChangedCount int

	Prop2             []rect `prop:"access:rw"`
	prop2ChangedCount int

	Prop3             map[string]rect `prop:"access:rw"`
	prop3ChangedCount int

	Prop4             twoRect `prop:"access:rw"`
	prop4ChangedCount int
}

func (*srvObject10) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject10) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object10"
}

func TestService_StructProp(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	srvObj10 := &srvObject10{
		Prop1: rect{1, 2, 3, 4},
	}

	serverObject10, err := service.NewServerObject("/org/deepin/dde/lib/Object10", srvObj10)

	_ = serverObject10.ConnectChanged(srvObj10, "Prop1", func(change *PropertyChanged) {
		srvObj10.prop1ChangedCount++
	})

	_ = serverObject10.ConnectChanged(srvObj10, "Prop2", func(change *PropertyChanged) {
		srvObj10.prop2ChangedCount++
	})

	_ = serverObject10.ConnectChanged(srvObj10, "Prop3", func(change *PropertyChanged) {
		srvObj10.prop3ChangedCount++
	})

	_ = serverObject10.ConnectChanged(srvObj10, "Prop4", func(change *PropertyChanged) {
		srvObj10.prop4ChangedCount++
	})

	err = serverObject10.Export()
	if err != nil {
		t.Error("Unexpected error exporting srvObj10:", err)
	}

	_ = service.RequestName("org.deepin.dde.lib.Object10")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj10 := service.conn.Object("org.deepin.dde.lib.Object10", "/org/deepin/dde/lib/Object10")

	// Test Prop1
	expectedProp1 := rect{2, 4, 6, 8}
	testProp1 := func() {
		err = clientObj10.Call(orgFreedesktopDBus+".Properties.Set", 0,
			"org.deepin.dde.lib.Object10", "Prop1", dbus.MakeVariant(expectedProp1)).Err

		if err != nil {
			t.Error("Unexpected error setting Prop1:", err)
		}
		if srvObj10.Prop1 != expectedProp1 {
			t.Errorf("prop1 expected %#v got %#v", expectedProp1, srvObj10.Prop1)
		}
		if srvObj10.prop1ChangedCount != 1 {
			t.Errorf("prop1ChangedCount expected 1 got %d", srvObj10.prop1ChangedCount)
		}

	}
	testProp1()
	testProp1() // Set the same value to Prop1 again

	// Test Prop2
	expectedProp2 := []rect{
		{1, 3, 5, 7},
		{2, 4, 6, 8},
	}

	testProp2 := func() {
		err = clientObj10.Call(orgFreedesktopDBus+".Properties.Set", 0,
			"org.deepin.dde.lib.Object10", "Prop2", dbus.MakeVariant(expectedProp2)).Err

		if err != nil {
			t.Error("Unexpected error setting Prop2:", err)
		}
		if !reflect.DeepEqual(srvObj10.Prop2, expectedProp2) {
			t.Errorf("prop2 expected %#v got %#v", expectedProp2, srvObj10.Prop2)
		}
		if srvObj10.prop2ChangedCount != 1 {
			t.Errorf("prop2ChangedCount expected 1 got %d", srvObj10.prop2ChangedCount)
		}
	}

	testProp2()
	testProp2()

	// Test Prop3
	expectedProp3 := map[string]rect{
		"a": {1, 3, 5, 7},
		"b": {2, 4, 6, 8},
	}

	testProp3 := func() {
		err = clientObj10.Call(orgFreedesktopDBus+".Properties.Set", 0,
			"org.deepin.dde.lib.Object10", "Prop3", dbus.MakeVariant(expectedProp3)).Err

		if err != nil {
			t.Error("Unexpected error setting Prop3", err)
		}
		if !reflect.DeepEqual(srvObj10.Prop3, expectedProp3) {
			t.Errorf("prop3 expected %#v got %#v", expectedProp3, srvObj10.Prop3)
		}
		if srvObj10.prop3ChangedCount != 1 {
			t.Errorf("prop3ChangedCount expected 1 got %d", srvObj10.prop3ChangedCount)
		}
	}

	testProp3()
	testProp3()

	// Test Prop4
	expectedProp4 := twoRect{
		A: rect{1, 3, 5, 7},
		B: rect{2, 4, 6, 8},
	}
	testProp4 := func() {
		err = clientObj10.Call(orgFreedesktopDBus+".Properties.Set", 0,
			"org.deepin.dde.lib.Object10", "Prop4", dbus.MakeVariant(expectedProp4)).Err

		if err != nil {
			t.Error("Unexpected error setting Prop4:", err)
		}
		if !reflect.DeepEqual(srvObj10.Prop4, expectedProp4) {
			t.Errorf("prop4 expected %#v got %#v", expectedProp4, srvObj10.Prop4)
		}
		if srvObj10.prop4ChangedCount != 1 {
			t.Errorf("prop4ChangedCount expected 1 got %d", srvObj10.prop4ChangedCount)
		}
	}
	testProp4()
	testProp4()
}

type srvObject11 struct {
	PropsMu sync.RWMutex

	Prop1 uint32
	Prop2 uint64
	Prop3 string

	Prop4   map[string]uint32
	Prop4Mu sync.RWMutex

	Prop5 *customProperty
}

func (*srvObject11) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject11) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object11"
}

type customProperty struct {
	value int32
}

func (*customProperty) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	return false, nil
}

func (cp *customProperty) GetValue() (val interface{}, err *dbus.Error) {
	return cp.value, nil
}

func (*customProperty) SetNotifyChangedFunc(func(val interface{})) {
}

func (cp *customProperty) GetType() reflect.Type {
	return reflect.TypeOf(cp.value)
}

func TestService_DumpProperties(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	srvObj11 := &srvObject11{
		Prop5: &customProperty{5},
	}
	err = service.Export("/org/deepin/dde/lib/Object11", srvObj11)
	if err != nil {
		t.Error("Unexpected error exporting srvObj11:", err)
	}
	propsInfo, err := service.DumpProperties(srvObj11)
	if err != nil {
		t.Error("Unexpected error dumping properties info:", err)
	}
	t.Log(propsInfo)
	t.Logf("PropsMu: %p", &srvObj11.PropsMu)
	t.Logf("Prop4Mu: %p", &srvObj11.Prop4Mu)
}

type integers struct {
	A int
	B uint
}

type srvObject13 struct {
	Prop1             int `prop:"access:rw"`
	prop1ChangedCount int

	Prop2             uint `prop:"access:rw"`
	prop2ChangedCount int

	Prop3             integers `prop:"access:rw"`
	prop3ChangedCount int
}

func (*srvObject13) GetExportedMethods() ExportedMethods {
	return nil
}

func (*srvObject13) GetInterfaceName() string {
	return "org.deepin.dde.lib.Object13"
}

func TestService_IntUintProp(t *testing.T) {
	if !isSessionBusExists() {
		t.Skip()
		return
	}

	service, err := NewSessionService()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	srvObj13 := &srvObject13{
		Prop1: 1,
		Prop2: 2,
		Prop3: integers{3, 4},
	}

	serverObject13, err := service.NewServerObject("/org/deepin/dde/lib/Object13", srvObj13)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	_ = serverObject13.ConnectChanged(srvObj13, "Prop1", func(change *PropertyChanged) {
		srvObj13.prop1ChangedCount++
	})

	_ = serverObject13.ConnectChanged(srvObj13, "Prop2", func(change *PropertyChanged) {
		srvObj13.prop2ChangedCount++
	})

	_ = serverObject13.ConnectChanged(srvObj13, "Prop3", func(change *PropertyChanged) {
		srvObj13.prop3ChangedCount++
	})

	err = serverObject13.Export()
	if err != nil {
		t.Error("Unexpected error exporting srvObj13:", err)
	}

	_ = service.RequestName("org.deepin.dde.lib.Object13")
	if err != nil {
		t.Error("Unexpected error calling RequestName:", err)
	}

	clientObj13 := service.conn.Object("org.deepin.dde.lib.Object13", "/org/deepin/dde/lib/Object13")

	// Test Prop1
	testProp1 := func() {
		err = clientObj13.Call(orgFreedesktopDBus+".Properties.Set", 0,
			"org.deepin.dde.lib.Object13", "Prop1", dbus.MakeVariant(int32(11))).Err
		if err != nil {
			t.Error("Unexpected error setting Prop1:", err)
		}

		if srvObj13.Prop1 != 11 {
			t.Errorf("prop1 expected 11 got %d", srvObj13.Prop1)
		}

		if srvObj13.prop1ChangedCount != 1 {
			t.Errorf("prop1ChangedCount expected 1 got %d", srvObj13.prop1ChangedCount)
		}
	}
	testProp1()
	testProp1()

	// Test Prop2
	testProp2 := func() {
		err = clientObj13.Call(orgFreedesktopDBus+".Properties.Set", 0,
			"org.deepin.dde.lib.Object13", "Prop2", dbus.MakeVariant(uint32(12))).Err
		if err != nil {
			t.Error("Unexpected error setting Prop2:", err)
		}

		if srvObj13.Prop2 != 12 {
			t.Errorf("prop2 expected 12 got %d", srvObj13.Prop2)
		}

		if srvObj13.prop2ChangedCount != 1 {
			t.Errorf("prop2ChangedCount expected 1 got %d", srvObj13.prop2ChangedCount)
		}
	}
	testProp2()
	testProp2()

	// Test Prop3
	expectedProp3 := integers{6, 7}
	testProp3 := func() {
		err = clientObj13.Call(orgFreedesktopDBus+".Properties.Set", 0,
			"org.deepin.dde.lib.Object13", "Prop3", dbus.MakeVariant(struct {
				A int32
				B uint32
			}{6, 7})).Err
		if err != nil {
			t.Error("Unexpected error setting Prop3:", err)
		}

		if srvObj13.Prop3 != expectedProp3 {
			t.Errorf("prop3 expected %#v got %#v", expectedProp3, srvObj13.Prop3)
		}

		if srvObj13.prop3ChangedCount != 1 {
			t.Errorf("prop3ChangedCount expected 1 got %d", srvObj13.prop3ChangedCount)
		}
	}
	testProp3()
	testProp3()
}
