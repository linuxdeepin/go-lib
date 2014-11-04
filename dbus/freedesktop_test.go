package dbus

import "testing"
import "reflect"

func TestIsStructreMatched(t *testing.T) {
	tInt := int32(0)
	ok := []interface{}{
		struct {
			Name string
			Age  int32
			X    *int32
			x    *int32
			List []string
		}{"snyh", 3, &tInt, &tInt, []string{"a"}},
		[]interface{}{
			"snyh",
			int32(26),
			int32(0),
			[]string{"a", "b"},
		},
	}
	for i := 0; i < len(ok); i = i + 2 {
		if !isStructureMatched(ok[i], ok[i+1]) {
			t.Fatal(ok[i], "Is Not Matched", ok[i+1])
		}
	}

	notOk := []interface{}{
		struct {
			Name string
			Age  int32
			X    uint32
		}{"snyh", 3, 3},
		[]interface{}{
			"snyh",
			int32(26),
			int32(0),
		},
	}
	for i := 0; i < len(notOk); i = i + 2 {
		if isStructureMatched(notOk[i], notOk[i+1]) {
			t.Fatal(notOk[i], "shouldn't Matched", notOk[i+1])
		}
	}

}

func TestIsExportedStructField(t *testing.T) {
	yes := struct {
		Yes1 string
		Yes2 *string
		Yes3 int32
	}{}
	for i := 0; i < reflect.TypeOf(yes).NumField(); i++ {
		if isExportedStructField(reflect.TypeOf(yes).Field(i)) == false {
			t.Fatal("Number of", i, "can't exported.")
		}
	}

	//TODO: isExportedStructField should check non-dbustype like int/uint
	no := struct {
		no1 string
		No2 string `dbus:"-"`
	}{}
	for i := 0; i < reflect.TypeOf(no).NumField(); i++ {
		if isExportedStructField(reflect.TypeOf(no).Field(i)) == true {
			t.Fatal("Number of", i, "can exported!!")
		}
	}

}

type testEmbedded struct {
	A string
	B int
}
type testEmebdedding struct {
	testEmbedded `dbus:"-"`
	A            string
}

func (testEmebdedding) Test() string {
	return "ABC"
}

func (testEmebdedding) GetDBusInfo() DBusInfo {
	return DBusInfo{
		Dest:       "com.deepin.test",
		ObjectPath: "/com/deepin/test",
		Interface:  "com.deepin.test",
	}
}

func TestEmbededStrcut(t *testing.T) {
	s := &testEmebdedding{}
	err := InstallOnSession(s)
	if err != nil {
		t.Skip("connect bus session failed" + err.Error())
	}
	c := detectConnByDBusObject(s)
	var ret string
	c.Object("com.deepin.test", "/com/deepin/test").Call("Test", 0).Store(&ret)
	if ret != "ABC" {
		t.Fail()
	}
	var props map[string]Variant
	c.Object("com.deepin.test", "/com/deepin/test").Call(
		"org.freedesktop.DBus.Properties.GetAll", 0, "com.deepin.test").Store(&props)
	_, ok := props["A"]
	if !ok {
		t.Fail()
	}

	var prop Variant
	err = c.Object("com.deepin.test", "/com/deepin/test").Call(
		"org.freedesktop.DBus.Properties.Get", 0, "com.deepin.test", "testEmbedded").Store(&prop)
	if err == nil {
		t.Fail()
	}

	err = c.Object("com.deepin.test", "/com/deepin/test").Call(
		"org.freedesktop.DBus.Properties.Get", 0, "com.deepin.test", "A").Store(&prop)
	if err != nil {
		t.Fail()
	}
}
