package property

import "testing"
import "pkg.deepin.io/lib/glib-2.0"

func TestGVariant(t *testing.T) {
	if glib.NewVariantByte(0).GetTypeString() != "y" {
		t.Fail()
	}
	if glib.NewVariantBoolean(false).GetTypeString() != "b" {
		t.Fail()
	}
	if glib.NewVariantInt16(0).GetTypeString() != "n" {
		t.Fail()
	}
	if glib.NewVariantUint16(0).GetTypeString() != "q" {
		t.Fail()
	}
	if glib.NewVariantInt32(0).GetTypeString() != "i" {
		t.Fail()
	}
	if glib.NewVariantUint32(0).GetTypeString() != "u" {
		t.Fail()
	}
	if glib.NewVariantInt64(0).GetTypeString() != "x" {
		t.Fail()
	}
	if glib.NewVariantUint64(0).GetTypeString() != "t" {
		t.Fail()
	}
	if glib.NewVariantDouble(3).GetTypeString() != "d" {
		t.Fail()
	}
	if glib.NewVariantString("/").GetTypeString() != "s" {
		t.Fail()
	}
	if glib.NewVariantObjectPath("/").GetTypeString() != "o" {
		t.Fail()
	}
	if glib.NewVariantSignature("as").GetTypeString() != "g" {
		t.Fail()
	}
	if glib.NewVariantStrv([]string{"/"}).GetTypeString() != "as" {
		t.Fail()
	}
}
