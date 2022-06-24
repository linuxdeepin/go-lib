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

package sm2

// #cgo pkg-config: openssl
// #include "dde-sm2.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type SM2Helper struct {
	context *C.sm2_context
}

func NewHelper() *SM2Helper {
	context := C.new_sm2_context()
	if context == nil {
		return nil
	}
	return &SM2Helper{
		context: context,
	}
}

func (s *SM2Helper) GenPairKey() (string, string) {
	pub := C.get_sm2_public_key(s.context)
	pri := C.get_sm2_private_key(s.context)

	return C.GoString(pub), C.GoString(pri)
}

func (s *SM2Helper) Encrypt(p []byte) ([]byte, error) {
	if len(p) == 0 {
		return nil, fmt.Errorf("plaintext size is zero")
	}
	size := C.get_ciphertext_size(s.context, C.size_t(len(p)))
	if size <= 0 {
		return nil, fmt.Errorf("get ciphertext size failed")
	}
	ret := make([]byte, size)

	n := C.encrypt(s.context, (*C.uint8_t)(unsafe.Pointer(&p[0])), C.size_t(len(p)), (*C.uint8_t)(unsafe.Pointer(&ret[0])), C.size_t(size))

	if n < 0 {
		return nil, fmt.Errorf("sm2 encrypt failed %d", n)
	}

	return ret[:n], nil
}

func (s *SM2Helper) Decrypt(c []byte) ([]byte, error) {
	if len(c) == 0 {
		return nil, fmt.Errorf("ciphertext size is zero")
	}
	size := C.get_plaintext_size((*C.uint8_t)(unsafe.Pointer(&c[0])), C.size_t(len(c)))
	if size <= 0 {
		return nil, fmt.Errorf("get plaintext size failed")
	}
	ret := make([]byte, size)

	n := C.decrypt(s.context, (*C.uint8_t)(unsafe.Pointer(&c[0])), C.size_t(len(c)), (*C.uint8_t)(unsafe.Pointer(&ret[0])), C.size_t(size))

	if n < 0 {
		return nil, fmt.Errorf("sm2 decrypt failed")
	}

	return ret[:n], nil
}

func (s *SM2Helper) Release() {
	if s.context != nil {
		C.free_sm2_context(s.context)
	}
	s.context = nil
}
