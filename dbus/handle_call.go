package dbus

import "fmt"
import "reflect"
import "strings"
import "unicode"

var autoHandler = map[string]func(*Conn, *Message) error{
	"org.freedesktop.DBus.Peer":           handlePeer,
	"org.freedesktop.DBus.Introspectable": handleIntrospect,
	//"org.freedesktop.DBus.LifeManager":    nil,
	//"org.freedesktop.DBus.Properties":     nil,
}

func handlePeer(conn *Conn, msg *Message) error {
	name := msg.Headers[FieldMember].value.(string)
	path := msg.Headers[FieldPath].value.(ObjectPath)
	ifcName, _ := msg.Headers[FieldInterface].value.(string)
	sender := msg.Headers[FieldSender].value.(string)
	serial := msg.serial
	switch name {
	case "Ping":
		conn.sendReply(sender, serial)
	case "GetMachineId":
		conn.sendReply(sender, serial, conn.uuid)
	}
	return NewUnknowMethod(path, ifcName, name)
}

func handleIntrospectionPartialPathRequest(possible_path []string, partial_path string) string {
	var xml string = `<node>`
	valid_field := make(map[string]bool)
	for _, path := range possible_path {
		begin := strings.Index(path, partial_path)
		if begin != -1 {
			path = path[begin+len(partial_path):]
			if len(path) == 0 {
				continue
			}
			if path[0] == '/' {
				path = path[1:]
			}
			end := strings.Index(path, "/")
			if end != -1 {
				path = path[:end]
			}
			valid_field[path] = true
		}
	}
	for k, _ := range valid_field {
		xml += `	<node name="` + k + `"/>`
	}
	xml += `</node>`
	return xml
}

func handleIntrospect(conn *Conn, msg *Message) error {
	path := msg.Headers[FieldPath].value.(ObjectPath)
	if _, ok := conn.handlers[path]; ok {
		return nil
	}
	name := msg.Headers[FieldMember].value.(string)
	if name != "Introspect" {
		return nil
	}

	sender := msg.Headers[FieldSender].value.(string)
	serial := msg.serial

	paths := make([]string, 0)
	for key, _ := range conn.handlers {
		paths = append(paths, string(key))
	}
	conn.sendReply(sender, serial, handleIntrospectionPartialPathRequest(paths, string(path)))
	return nil
}

func (conn *Conn) parseParams(methodType reflect.Type, msg *Message) ([]reflect.Value, error) {
	flags := detectExportMethodFlags(methodType)
	needDMessage := flags&UserMethodFlagNeedDMessage != 0

	vs := msg.Body
	pointers := make([]interface{}, len(vs))
	decode := make([]interface{}, 0)
	for i := 0; i < len(vs); i++ {
		var tp reflect.Type
		if needDMessage {
			tp = methodType.In(i + 1)
		} else {
			tp = methodType.In(i)
		}
		val := reflect.New(tp)
		pointers[i] = val.Interface()
		decode = append(decode, pointers[i])
	}
	if len(decode) != len(vs) {
		return nil, NewInvalidArg(fmt.Sprintf("Need %d paramters but get %d", len(decode), len(vs)))
	}
	if err := Store(vs, decode...); err != nil {
		return nil, NewInvalidArg(err.Error())
	}
	params := make([]reflect.Value, len(pointers))
	for i := 0; i < len(pointers); i++ {
		params[i] = reflect.ValueOf(pointers[i]).Elem()
	}
	if needDMessage {
		params = append([]reflect.Value{reflect.ValueOf(DMessage{msg, conn})}, params...)
	}
	return params, nil
}

func (conn *Conn) callUserMethod(msg *Message) ([]reflect.Value, error) {
	name := msg.Headers[FieldMember].value.(string)
	path := msg.Headers[FieldPath].value.(ObjectPath)
	ifcName, _ := msg.Headers[FieldInterface].value.(string)

	conn.handlersLck.RLock()
	ifcs, ok := conn.handlers[path]
	conn.handlersLck.RUnlock()

	if !ok {
		return nil, NewNoObjectError(path)
	}

	var userMethod reflect.Value
	if ifc, ok := ifcs[ifcName]; ok {
		userMethod = reflect.ValueOf(ifc).MethodByName(name)
	} else {
		for _, ifc := range ifcs {
			userMethod = reflect.ValueOf(ifc).MethodByName(name)
			if userMethod.IsValid() {
				break
			}
		}
	}
	if !userMethod.IsValid() {
		return nil, NewUnknowMethod(path, ifcName, name)
	}

	methodType := userMethod.Type()

	params, err := conn.parseParams(methodType, msg)
	if err != nil {
		return nil, err
	}

	ret := userMethod.Call(params)

	flags := detectExportMethodFlags(methodType)

	if flags&UserMethodFlagWillThrowError != 0 {
		out_n := methodType.NumOut()
		v := ret[out_n-1].Interface()
		if v != nil {
			if em, ok := v.(dbusError); ok {
				return nil, em
			} else if goErr, ok := v.(error); ok {
				return nil, NewOtherError(goErr.Error())
			}
		}
		ret = ret[:out_n-1]
	}
	for i, r := range ret {
		ret[i] = tryTranslateDBusObjectToObjectPath(conn, r)
	}
	return ret, nil
}

// handleCall handles the given method call (i.e. looks if it's one of the
// pre-implemented ones and searches for a corresponding handler if not).
func (conn *Conn) handleCall(msg *Message) {
	name := msg.Headers[FieldMember].value.(string)
	path := msg.Headers[FieldPath].value.(ObjectPath)
	ifcName, _ := msg.Headers[FieldInterface].value.(string)
	sender := msg.Headers[FieldSender].value.(string)
	serial := msg.serial
	defer func() {
		if err := recover(); err != nil {
			conn.sendError(newInternalError(err), sender, serial)
		}
	}()

	if len(name) == 0 || unicode.IsLower([]rune(name)[0]) {
		conn.sendError(NewUnknowMethod(path, ifcName, name), sender, serial)
		return
	}

	if handler, ok := autoHandler[ifcName]; ok {
		err := handler(conn, msg)
		if err != nil {
			conn.sendError(err, sender, serial)
			return
		}
		if ifcName == "org.freedesktop.DBus.Introspectable" && conn.handlers[path] != nil {
			//workaround, continue handle
		} else {
			return
		}
	}

	ret, err := conn.callUserMethod(msg)
	if err != nil {
		conn.sendError(err, sender, serial)
		return
	}

	if msg.Flags&FlagNoReplyExpected == 0 {
		reply := new(Message)
		reply.Type = TypeMethodReply
		reply.serial = conn.getSerial()
		reply.Headers = make(map[HeaderField]Variant)
		reply.Headers[FieldDestination] = msg.Headers[FieldSender]
		reply.Headers[FieldReplySerial] = MakeVariant(msg.serial)
		reply.Body = make([]interface{}, len(ret))
		for i := 0; i < len(ret); i++ {
			reply.Body[i] = ret[i].Interface()
		}
		if len(ret) != 0 {
			reply.Headers[FieldSignature] = MakeVariant(SignatureOf(reply.Body...))
		}
		conn.outLck.RLock()
		if !conn.closed {
			conn.out <- reply
		}
		conn.outLck.RUnlock()
	}

	_HasNewMessage = true
}

func tryTranslateDBusObjectToObjectPath(con *Conn, value reflect.Value) reflect.Value {
	if value.Type().Implements(dbusObjectInterface) {
		if !value.IsNil() {
			obj := value.Interface().(DBusObject)
			//TODO: which session to install
			InstallOnAny(con, obj)
			return reflect.ValueOf(ObjectPath(obj.GetDBusInfo().ObjectPath))
		} else {
			return reflect.ValueOf(ObjectPath("/"))
		}
	}
	switch value.Type().Kind() {
	case reflect.Array, reflect.Slice:
		n := value.Len()
		elemType := value.Type().Elem()
		if elemType.Implements(dbusObjectInterface) {
			elemType = objectPathType
		}
		new_value := reflect.MakeSlice(reflect.SliceOf(elemType), n, n)
		for i := 0; i < n; i++ {
			new_value.Index(i).Set(tryTranslateDBusObjectToObjectPath(con, value.Index(i)))
		}
		return new_value
	case reflect.Map:
		keys := value.MapKeys()
		if len(keys) == 0 {
			return value
		}
		t := tryTranslateDBusObjectToObjectPath(con, reflect.Zero(value.Type().Elem())).Type()
		if t == value.Type().Elem() {
			return value
		}
		new_value := reflect.MakeMap(reflect.MapOf(value.Type().Key(), t))
		for i := 0; i < len(keys); i++ {
			new_value.SetMapIndex(keys[i], tryTranslateDBusObjectToObjectPath(con, value.MapIndex(keys[i])))
		}
		return new_value
	}
	// Is possible dynamic create an struct or change one struct's some field's type by some trick?
	return value
}
