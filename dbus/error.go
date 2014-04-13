package dbus

import "reflect"

// Error represents a D-Bus message of type Error.
type dbusError struct {
	Name string
	Body []interface{}
}

var (
	goErrorType   = reflect.TypeOf((*error)(nil)).Elem()
	dbusErrorType = reflect.TypeOf((*dbusError)(nil))
)

func (e dbusError) Error() string {
	if len(e.Body) >= 1 {
		s, ok := e.Body[0].(string)
		if ok {
			return e.Name + ":" + s
		}
	}
	return e.Name
}

const (
	NoObjectError = iota
	UnknowInterfaceError
	UnknowMethodError
	UnknowPropertyError
	OtherError
)

func NewCustomError(name string, args ...interface{}) dbusError {
	return dbusError{
		name,
		args,
	}
}
func NewNoObjectError(path ObjectPath) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.NoSuchObject",
		[]interface{}{"No such object" + string(path)},
	}
}
func newError(errType int, args ...interface{}) dbusError {
	//TODO: complete this
	name := "UnknowError"
	switch errType {
	case NoObjectError:
		name = "org.freedesktop.DBus.Error.NoSuchObject"
	}
	return dbusError{
		name,
		args,
	}
}

func NewPropertyNotWritableError(name string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.NoWritable",
		[]interface{}{"Can't write this property."},
	}
}

func NewUnknowInterfaceError(ifcName string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.NoSuchInterface",
		[]interface{}{"No such interface"},
	}
}
func NewUnknowPropertyError(name string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.UnknownProperty",
		[]interface{}{"Unknown / invalid Property"},
	}
}

func NewOtherError(body interface{}) dbusError {
	return dbusError{
		"com.deepin.DBus.Error.UnknowError",
		[]interface{}{body},
	}
}
func newInternalError(body interface{}) dbusError {
	return dbusError{
		"com.deepin.DBus.Error.InternalError",
		[]interface{}{body},
	}
}
func NewUnknowMethod(path ObjectPath, ifc, name string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.UnknownMethod",
		[]interface{}{"Cant find the method of " + name},
	}
}
func NewInvalidArg(content string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.InvalidArgs",
		[]interface{}{"Invalid type / number of args" + content},
	}
}
