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

package dbus

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"os"
)

// AuthCookieSha1 returns an Auth that authenticates as the given user with the
// DBUS_COOKIE_SHA1 mechanism. The home parameter should specify the home
// directory of the user.
func AuthCookieSha1(user, home string) Auth {
	return authCookieSha1{user, home}
}

type authCookieSha1 struct {
	user, home string
}

func (a authCookieSha1) FirstData() ([]byte, []byte, AuthStatus) {
	b := make([]byte, 2*len(a.user))
	hex.Encode(b, []byte(a.user))
	return []byte("DBUS_COOKIE_SHA1"), b, AuthContinue
}

func (a authCookieSha1) HandleData(data []byte) ([]byte, AuthStatus) {
	challenge := make([]byte, len(data)/2)
	_, err := hex.Decode(challenge, data)
	if err != nil {
		return nil, AuthError
	}
	b := bytes.Split(challenge, []byte{' '})
	if len(b) != 3 {
		return nil, AuthError
	}
	context := b[0]
	id := b[1]
	svchallenge := b[2]
	cookie := a.getCookie(context, id)
	if cookie == nil {
		return nil, AuthError
	}
	clchallenge := a.generateChallenge()
	if clchallenge == nil {
		return nil, AuthError
	}
	hash := sha1.New()
	hash.Write(bytes.Join([][]byte{svchallenge, clchallenge, cookie}, []byte{':'}))
	hexhash := make([]byte, 2*hash.Size())
	hex.Encode(hexhash, hash.Sum(nil))
	data = append(clchallenge, ' ')
	data = append(data, hexhash...)
	resp := make([]byte, 2*len(data))
	hex.Encode(resp, data)
	return resp, AuthOk
}

// getCookie searches for the cookie identified by id in context and returns
// the cookie content or nil. (Since HandleData can't return a specific error,
// but only whether an error occured, this function also doesn't bother to
// return an error.)
func (a authCookieSha1) getCookie(context, id []byte) []byte {
	file, err := os.Open(a.home + "/.dbus-keyrings/" + string(context))
	if err != nil {
		return nil
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			return nil
		}
		line = line[:len(line)-1]
		b := bytes.Split(line, []byte{' '})
		if len(b) != 3 {
			return nil
		}
		if bytes.Equal(b[0], id) {
			return b[2]
		}
	}
}

// generateChallenge returns a random, hex-encoded challenge, or nil on error
// (see above).
func (a authCookieSha1) generateChallenge() []byte {
	b := make([]byte, 16)
	n, err := rand.Read(b)
	if err != nil {
		return nil
	}
	if n != 16 {
		return nil
	}
	enc := make([]byte, 32)
	hex.Encode(enc, b)
	return enc
}
