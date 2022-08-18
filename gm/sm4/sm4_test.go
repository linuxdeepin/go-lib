package sm4

import (
	"bytes"
	"fmt"
	"testing"
)

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	if blockSize == 0 {
		return []byte("")
	}
	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	padding := int(origData[length-1])
	return origData[:(length - padding)]
}

func TestSm4_ecb(t *testing.T) {
	key := []byte("0000000000000000")
	in := []byte("ssssssssssssssss")

	cipher, _ := NewCipher(key)
	src := pkcs5Padding(in, BlockSize)
	block := make([]byte, len(src))
	cipher.Encrypt(block, src)
	cipher.Decrypt(block, block)

	plainText := pkcs5UnPadding(src)
	fmt.Println(string(plainText))
	if !bytes.Equal(in, plainText) {
		t.Error("decrypt result not equal expected")
		return
	}
}
