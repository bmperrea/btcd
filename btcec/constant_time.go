// Copyright (c) 2017 The btcsuite developers
// Copyright (c) 2017 Brent Perreault
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcec

// constant_time.go provides constant time implementations of useful
// mathematical operations. In addition, these functions return integers,
// using 0 or 1 to represent false or true respectively, which is useful
// for writing logic in terms of bitwise operators

// References
// These functions are based on the sample implementation in
// golang.org/src/crypto/subtle/constant_time.go
// Here we have refactored these functions for uint32 arithmetic and
// to avoid extra shifts and casts

// Note - these assume that the values are less than 2^31 and use the
// most significant bit of uint32 to indicate underflow in modular arithmetic.

// These are intended for use internal to btcec.

// lessThanUint32 returns 1 if x < y and 0 otherwise.
func lessThanUint32(x, y uint32) uint32 {
	diff := x - y
	return diff >> 31
}

// isZeroUint32 returns 1 if x == 0 and 0 otherwise.
func isZeroUint32(x uint32) uint32 {
	return ((x - 1) ^ x) >> 31
}

// notZeroUint32 returns 1 if x != 0 and 0 otherwise.
//
// Trick: there is only one value which has a high bit of 0,
// both before and after negation: uint32(0).
func notZeroUint32(x uint32) uint32 {
	return ((-x) | x) >> 31
}
