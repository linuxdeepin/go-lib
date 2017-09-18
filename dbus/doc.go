/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

/*
Package dbus implements bindings to the D-Bus message bus system.

To use the message bus API, you first need to connect to a bus (usually the
session or system bus). The acquired connection then can be used to call methods
on remote objects and emit or receive signals. Using the Export method, you can
arrange D-Bus methods calls to be directly translated to method calls on a Go
value.

Conversion Rules

For outgoing messages, Go types are automatically converted to the
corresponding D-Bus types. The following types are directly encoded as their
respective D-Bus equivalents:

     Go type     | D-Bus type
     ------------+-----------
     byte        | BYTE
     bool        | BOOLEAN
     int16       | INT16
     uint16      | UINT16
     int32       | INT32
     uint32      | UINT32
     int64       | INT64
     uint64      | UINT64
     float64     | DOUBLE
     string      | STRING
     ObjectPath  | OBJECT_PATH
     Signature   | SIGNATURE
     Variant     | VARIANT
     UnixFDIndex | UNIX_FD

Slices and arrays encode as ARRAYs of their element type.

Maps encode as DICTs, provided that their key type can be used as a key for
a DICT.

Structs other than Variant and Signature encode as a STRUCT containing their
exported fields. Fields whose tags contain `dbus:"-"` and unexported fields will
be skipped.

Pointers encode as the value they're pointed to.

Trying to encode any other type or a slice, map or struct containing an
unsupported type will result in an InvalidTypeError.

For incoming messages, the inverse of these rules are used, with the exception
of STRUCTs. Incoming STRUCTS are represented as a slice of empty interfaces
containing the struct fields in the correct order. The Store function can be
used to convert such values to Go structs.

Unix FD passing

Handling Unix file descriptors deserves special mention. To use them, you should
first check that they are supported on a connection by calling SupportsUnixFDs.
If it returns true, all method of Connection will translate messages containing
UnixFD's to messages that are accompanied by the given file descriptors with the
UnixFD values being substituted by the correct indices. Similarily, the indices
of incoming messages are automatically resolved. It shouldn't be necessary to use
UnixFDIndex.

*/
package dbus
