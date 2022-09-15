// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package appinfo

import (
	"bytes"
	"io"

	"github.com/linuxdeepin/go-x11-client"
)

type StartupNotifyMessage struct {
	Type      string
	KeyValues map[string]string
}

func (msg *StartupNotifyMessage) fillBuffer() *bytes.Buffer {
	var buf bytes.Buffer

	buf.WriteString(msg.Type)
	buf.WriteByte(':')

	for k, v := range msg.KeyValues {
		buf.WriteString(k)
		buf.WriteString(`="`)

		vlen := len(v)
		for i := 0; i < vlen; i++ {
			b := v[i]
			switch b {
			case 0, ' ', '"', '\\':
				buf.WriteByte('\\')
			}
			buf.WriteByte(b)
		}

		buf.WriteString(`" `)
	}

	buf.WriteByte(0)
	return &buf
}

func (msg *StartupNotifyMessage) Broadcast(conn *x.Conn) error {
	atomMsgType, err := conn.GetAtom("_NET_STARTUP_INFO")
	if err != nil {
		return err
	}
	atomMsgTypeBegin, err := conn.GetAtom("_NET_STARTUP_INFO_BEGIN")
	if err != nil {
		return err
	}
	return broadcastXMessage(conn, atomMsgType, atomMsgTypeBegin, msg.fillBuffer())
}

func broadcastXMessage(conn *x.Conn, atomMsgType, atomMsgTypeBegin x.Atom, msgReader io.Reader) error {
	// create window
	xid, err := conn.AllocID()
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.FreeID(xid)
	}()

	win := x.Window(xid)
	root := conn.GetDefaultScreen().Root
	err = x.CreateWindowChecked(conn, 0, win, root, 0, 0, 1, 1,
		0, x.WindowClassInputOnly, x.CopyFromParent,
		x.CWOverrideRedirect|x.CWEventMask,
		[]uint32{1, x.EventMaskPropertyChange | x.EventMaskStructureNotify}).Check(conn)
	if err != nil {
		return err
	}

	// send x message
	ev := x.ClientMessageEvent{
		Format: 8,
		Window: win,
		Type:   atomMsgTypeBegin,
	}

	var buf [20]byte
	var readDone bool

	for !readDone {
		n, err := msgReader.Read(buf[:])
		if err != nil {
			// EOF
			readDone = true
		}
		if n == 0 {
			break
		}
		ev.Data = x.ClientMessageData{}
		ev.Data.SetData8(&buf)
		w := x.NewWriter()
		x.WriteClientMessageEvent(w, &ev)
		x.SendEvent(conn, false, root,
			x.EventMaskPropertyChange, w.Bytes())
		ev.Type = atomMsgType
	}

	x.DestroyWindow(conn, win)
	conn.Flush()
	return nil
}
