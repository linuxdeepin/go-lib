// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sm3

import (
	"encoding/binary"
	"hash"
	"math/bits"
)

const (
	DigestSize  = 32
	SizeBitSize = 5
	BlockSize   = 64
)

// 国标规定的常数T
// T0 = 0x79CC4519 T16 = 0x7A879D8A
// Tn = bits.RotateLeft32(T0, n) (n = [0, 15))
// Tn = bits.RotateLeft32(T16, n) (n = [16, 64))
// 提前生成好提高性能
var T = []uint32{
	0x79CC4519, 0xF3988A32, 0xE7311465, 0xCE6228CB, 0x9CC45197, 0x3988A32F, 0x7311465E, 0xE6228CBC,
	0xCC451979, 0x988A32F3, 0x311465E7, 0x6228CBCE, 0xC451979C, 0x88A32F39, 0x11465E73, 0x228CBCE6,
	0x9D8A7A87, 0x3B14F50F, 0x7629EA1E, 0xEC53D43C, 0xD8A7A879, 0xB14F50F3, 0x629EA1E7, 0xC53D43CE,
	0x8A7A879D, 0x14F50F3B, 0x29EA1E76, 0x53D43CEC, 0xA7A879D8, 0x4F50F3B1, 0x9EA1E762, 0x3D43CEC5,
	0x7A879D8A, 0xF50F3B14, 0xEA1E7629, 0xD43CEC53, 0xA879D8A7, 0x50F3B14F, 0xA1E7629E, 0x43CEC53D,
	0x879D8A7A, 0x0F3B14F5, 0x1E7629EA, 0x3CEC53D4, 0x79D8A7A8, 0xF3B14F50, 0xE7629EA1, 0xCEC53D43,
	0x9D8A7A87, 0x3B14F50F, 0x7629EA1E, 0xEC53D43C, 0xD8A7A879, 0xB14F50F3, 0x629EA1E7, 0xC53D43CE,
	0x8A7A879D, 0x14F50F3B, 0x29EA1E76, 0x53D43CEC, 0xA7A879D8, 0x4F50F3B1, 0x9EA1E762, 0x3D43CEC5}

type SM3 struct {
	digest [DigestSize / 4]uint32 // digest represents the partial evaluation of v
	length uint64                 // length of the message
	buff   []byte                 // unhandle message
}

func ff0(x, y, z uint32) uint32 { return x ^ y ^ z }
func ff1(x, y, z uint32) uint32 { return (x & y) | (x & z) | (y & z) }
func gg0(x, y, z uint32) uint32 { return x ^ y ^ z }
func gg1(x, y, z uint32) uint32 { return (x & y) | ((^x) & z) }

func p0(x uint32) uint32 {
	r9 := bits.RotateLeft32(x, 9)
	r17 := bits.RotateLeft32(x, 17)
	return x ^ r9 ^ r17
}

func p1(x uint32) uint32 {
	r15 := bits.RotateLeft32(x, 15)
	r23 := bits.RotateLeft32(x, 23)
	return x ^ r15 ^ r23
}

func (s *SM3) pad() []byte {
	msg := s.buff
	msg = append(msg, 0x80) // Append '1'
	for len(msg)%BlockSize != 56 {
		msg = append(msg, 0x00)
	}
	// append message length
	msg = append(msg, uint8(s.length>>56&0xff))
	msg = append(msg, uint8(s.length>>48&0xff))
	msg = append(msg, uint8(s.length>>40&0xff))
	msg = append(msg, uint8(s.length>>32&0xff))
	msg = append(msg, uint8(s.length>>24&0xff))
	msg = append(msg, uint8(s.length>>16&0xff))
	msg = append(msg, uint8(s.length>>8&0xff))
	msg = append(msg, uint8(s.length>>0&0xff))

	if len(msg)%BlockSize != 0 {
		panic("------SM3 Pad: error msgLen =")
	}

	return msg
}

// 核心算法
func process(digest [8]uint32, msg []byte) ([8]uint32, []byte) {
	var w [68]uint32
	var w1 [64]uint32

	a, b, c, d, e, f, g, h := digest[0], digest[1], digest[2], digest[3], digest[4], digest[5], digest[6], digest[7]
	for len(msg) >= BlockSize {
		for i := 0; i < 16; i++ {
			w[i] = binary.BigEndian.Uint32(msg[4*i : 4*(i+1)])
		}
		for i := 16; i < 68; i++ {
			w[i] = p1(w[i-16]^w[i-9]^bits.RotateLeft32(w[i-3], 15)) ^ bits.RotateLeft32(w[i-13], 7) ^ w[i-6]
		}
		for i := 0; i < 64; i++ {
			w1[i] = w[i] ^ w[i+4]
		}
		A, B, C, D, E, F, G, H := a, b, c, d, e, f, g, h
		for i := 0; i < 16; i++ {
			a12 := bits.RotateLeft32(A, 12)
			s1 := a12 + E + T[i]
			SS1 := bits.RotateLeft32(s1, 7)
			SS2 := SS1 ^ a12
			TT1 := ff0(A, B, C) + D + SS2 + w1[i]
			TT2 := gg0(E, F, G) + H + SS1 + w[i]
			D = C
			C = bits.RotateLeft32(B, 9)
			B = A
			A = TT1
			H = G
			G = bits.RotateLeft32(F, 19)
			F = E
			E = p0(TT2)
		}
		for i := 16; i < 64; i++ {
			a12 := bits.RotateLeft32(A, 12)
			s1 := a12 + E + T[i]
			SS1 := bits.RotateLeft32(s1, 7)
			SS2 := SS1 ^ a12
			TT1 := ff1(A, B, C) + D + SS2 + w1[i]
			TT2 := gg1(E, F, G) + H + SS1 + w[i]
			D = C
			C = bits.RotateLeft32(B, 9)
			B = A
			A = TT1
			H = G
			G = bits.RotateLeft32(F, 19)
			F = E
			E = p0(TT2)
		}
		a ^= A
		b ^= B
		c ^= C
		d ^= D
		e ^= E
		f ^= F
		g ^= G
		h ^= H
		msg = msg[BlockSize:]
	}

	return [8]uint32{a, b, c, d, e, f, g, h}, msg
}

// New create a hash instance
func New() hash.Hash {
	s := &SM3{}
	s.Reset()

	return s
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (s *SM3) BlockSize() int { return BlockSize }

// Size returns the number of bytes Sum will return.
func (s *SM3) Size() int { return DigestSize }

// Reset clears the internal state by zeroing bytes in the state buffer.
// This can be skipped for a newly-created hash state; the default zero-allocated state is correct.
func (s *SM3) Reset() {
	// Reset digest
	s.digest[0] = 0x7380166f
	s.digest[1] = 0x4914b2b9
	s.digest[2] = 0x172442d7
	s.digest[3] = 0xda8a0600
	s.digest[4] = 0xa96f30bc
	s.digest[5] = 0x163138aa
	s.digest[6] = 0xe38dee4d
	s.digest[7] = 0xb0fb0e4e

	s.length = 0 // Reset numberic states
	s.buff = []byte{}
}

// Write hash interface
// It never returns an error.
func (s *SM3) Write(p []byte) (int, error) {
	n := len(p)
	s.length += uint64(n << 3)
	s.buff = append(s.buff, p...)
	s.digest, s.buff = process(s.digest, s.buff)

	return n, nil
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (s *SM3) Sum(in []byte) []byte {
	_, _ = s.Write(in)
	msg := s.pad()
	digest, _ := process(s.digest, msg)

	out := make([]byte, DigestSize)
	for i := range digest {
		binary.BigEndian.PutUint32(out[i*4:], digest[i])
	}

	return out
}

// Sum returns the SM3 checksum of the data.
func Sum(data []byte) []byte {
	var s SM3
	s.Reset()
	s.Write(data)
	return s.Sum(nil)
}

// Kdf key derivation function, compliance with GB/T 32918.4-2016 5.4.3.
func Kdf(z []byte, len int) ([]byte, bool) {
	limit := (len + DigestSize - 1) >> SizeBitSize
	md := New()
	var countBytes [4]byte
	var ct uint32 = 1
	k := make([]byte, len+DigestSize-1)
	for i := 0; i < limit; i++ {
		binary.BigEndian.PutUint32(countBytes[:], ct)
		md.Write(z)
		md.Write(countBytes[:])
		copy(k[i*DigestSize:], md.Sum(nil))
		ct++
		md.Reset()
	}
	for i := 0; i < len; i++ {
		if k[i] != 0 {
			return k[:len], true
		}
	}
	return k, false
}
