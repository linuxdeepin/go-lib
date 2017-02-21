package appinfo

import (
	"bytes"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xprop"
	"github.com/BurntSushi/xgbutil/xwindow"
	"io"
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

var _atomMsgType, _atomMsgTypeBegin xproto.Atom

func getAtomMsgType(xu *xgbutil.XUtil) xproto.Atom {
	if _atomMsgType != 0 {
		return _atomMsgType
	}
	_atomMsgType, _ = xprop.Atm(xu, "_NET_STARTUP_INFO")
	return _atomMsgType
}

func getAtomMsgTypeBegin(xu *xgbutil.XUtil) xproto.Atom {
	if _atomMsgTypeBegin != 0 {
		return _atomMsgTypeBegin
	}
	_atomMsgTypeBegin, _ = xprop.Atm(xu, "_NET_STARTUP_INFO_BEGIN")
	return _atomMsgTypeBegin
}

func (msg *StartupNotifyMessage) Broadcast(xu *xgbutil.XUtil) error {
	return broadcastXMessage(xu, getAtomMsgType(xu), getAtomMsgTypeBegin(xu), msg.fillBuffer())
}

func broadcastXMessage(xu *xgbutil.XUtil, atomMsgType, atomMsgTypeBegin xproto.Atom, msgReader io.Reader) error {
	// create window
	win, err := xwindow.Generate(xu)
	if err != nil {
		return err
	}
	win.Create(xu.RootWin(), // parent
		-100, -100, 1, 1, // x, y, width, height
		xproto.CwOverrideRedirect|xproto.CwEventMask, // value mask
		1, xproto.EventMaskPropertyChange|xproto.EventMaskStructureNotify) // value list

	// send x message
	ev := &xproto.ClientMessageEvent{
		Format: 8,
		Window: win.Id,
		Type:   atomMsgTypeBegin,
	}

	const bufLen = 20
	buf := make([]byte, bufLen)
	var readDone bool

	for !readDone {
		n, err := msgReader.Read(buf)
		if err != nil {
			// EOF
			readDone = true
		}
		if n == 0 {
			break
		}
		ev.Data = xproto.ClientMessageDataUnion{Data8: buf}

		err = xevent.SendRootEvent(xu, ev, xproto.EventMaskPropertyChange)
		if err != nil {
			return err
		}

		ev.Type = atomMsgType
	}

	win.Destroy()

	return nil
}
