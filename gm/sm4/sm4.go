/*
 * Copyright (C) 2015 ~ 2022 Deepin Technology Co., Ltd.
 *
 * Author:     liaohanqin <liaohanqin@uniontech.com>
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

package sm4

// #include <openssl/sm4.h>
// #cgo pkg-config: openssl
import "C"
import (
	"crypto/cipher"
	"fmt"
	"unsafe"
)

// The SM4 block size in bytes.
const (
	BlockSize = 16
	KeySize   = 16
)

// A cipher is an instance of SM4 encryption using a particular key.
type sm4Cipher struct {
	key C.SM4_KEY
}

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the SM4 key,
func NewCipher(key []byte) (cipher.Block, error) {
	k := len(key)
	switch k {
	default:
		return nil, fmt.Errorf("sm4: invalid key size %d", k)
	case KeySize:
		break
	}
	ret := &sm4Cipher{}
	C.SM4_set_key((*C.uint8_t)(unsafe.Pointer(&key[0])), &ret.key)
	return ret, nil
}

func (c *sm4Cipher) BlockSize() int { return BlockSize }

func (c *sm4Cipher) Encrypt(dst, src []byte) {
	C.SM4_encrypt((*C.uint8_t)(unsafe.Pointer(&src[0])), (*C.uint8_t)(unsafe.Pointer(&dst[0])), &c.key)
}

func (c *sm4Cipher) Decrypt(dst, src []byte) {
	C.SM4_decrypt((*C.uint8_t)(unsafe.Pointer(&src[0])), (*C.uint8_t)(unsafe.Pointer(&dst[0])), &c.key)
}
