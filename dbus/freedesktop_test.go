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
