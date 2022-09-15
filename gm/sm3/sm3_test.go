// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sm3

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

var testData = map[string]string{
	"abc": "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0",
	"abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd": "debe9ff92275b8a138604889c18e5a4d6fdb70e5387e5765293dcba39c0c5732"}

func TestSm3_Sum(t *testing.T) {
	for src, expected := range testData {
		testSm3Sum(t, src, expected)
	}
}

func testSm3Sum(t *testing.T, src string, expected string) {
	d := New()
	d.Write([]byte(src))
	hash := d.Sum(nil)
	hashHex := hex.EncodeToString(hash[:])
	if hashHex != expected {
		t.Errorf("result:%s , not equal expected\n", hashHex)
		return
	}
}

func TestSm3_Write(t *testing.T) {
	src1 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	src2 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	src3 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := New()
	s.Write(src1)
	s.Write(src2)
	s.Write(src3)
	digest1 := s.Sum(nil)
	fmt.Printf("1 : %s\n", hex.EncodeToString(digest1))

	s.Reset()
	s.Write(src1)
	s.Write(src2)
	s.Write(src3)
	digest2 := s.Sum(nil)
	fmt.Printf("2 : %s\n", hex.EncodeToString(digest2))

	if !bytes.Equal(digest1, digest2) {
		t.Error("")
		return
	}
}
