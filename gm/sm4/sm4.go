// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sm4

/*
#cgo pkg-config: openssl
#include <openssl/evp.h>

static void openssl_evp_sm4_cipher(const unsigned char *key,
                                   unsigned char *out,
                                   unsigned char *in, int inl,
                                   int enc) {
    int ret = 0;
    EVP_CIPHER_CTX *ctx = EVP_CIPHER_CTX_new();
    if (ctx == NULL) {
        return;
    }
    ret = EVP_CipherInit(ctx, EVP_sm4_ecb(), key, NULL, enc);
    if (1 != ret) {
        printf("EVP_CipherInit fail... ret = %d \n", ret);
        EVP_CIPHER_CTX_free(ctx);
        return;
    }
    ret = EVP_Cipher(ctx, out, in, inl);
    if (1 != ret) {
        printf("EVP_Cipher fail.. ret = %d \n", ret);
    }

    EVP_CIPHER_CTX_free(ctx);
}
*/
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
	key []byte
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

	ret.key = make([]byte, k)
	copy(ret.key, key)
	return ret, nil
}

func (c *sm4Cipher) BlockSize() int { return BlockSize }

func (c *sm4Cipher) Encrypt(dst, src []byte) {
	C.openssl_evp_sm4_cipher((*C.uint8_t)(unsafe.Pointer(&c.key[0])),
		(*C.uint8_t)(unsafe.Pointer(&dst[0])),
		(*C.uint8_t)(unsafe.Pointer(&src[0])), C.int(len(src)),
		1)
}

func (c *sm4Cipher) Decrypt(dst, src []byte) {
	C.openssl_evp_sm4_cipher((*C.uint8_t)(unsafe.Pointer(&c.key[0])),
		(*C.uint8_t)(unsafe.Pointer(&dst[0])),
		(*C.uint8_t)(unsafe.Pointer(&src[0])), C.int(len(src)),
		0)
}
