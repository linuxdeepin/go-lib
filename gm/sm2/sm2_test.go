// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sm2

import (
	"reflect"
	"testing"
)

func Test_encryptDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		plainText string
	}{
		// TODO: Add test cases.
		{"simple", "12345"},
		{"less than 32", "encryption standard"},
		{"equals 32", "encryption standard encryption "},
		{"long than 32", "encryption standard encryption standard"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := NewHelper()
			if helper == nil {
				t.Fatalf("new sm2 helper failed")
			}
			defer helper.Release()
			ciphertext, err := helper.Encrypt([]byte(tt.plainText))
			if err != nil {
				t.Fatalf("encrypt failed %v", err)
			}
			plaintext, err := helper.Decrypt(ciphertext)
			if err != nil {
				t.Fatalf("decrypt failed %v", err)
			}
			if !reflect.DeepEqual(string(plaintext), tt.plainText) {
				t.Errorf("Decrypt() = %v, want %v", string(plaintext), tt.plainText)
			}
		})
	}
}
