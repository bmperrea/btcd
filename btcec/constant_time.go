// Copyright (c) 2013-2014 The btcsuite developers
// Copyright (c) 2013-2014 Dave Collins
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcec

// constant_time.go provides constant time implementations of useful mathematical operations. In addition, these
// functions return integers, using 0 or 1 to represent false or true respectively,
// which is useful for writing logic in terms of bitwise operators

// References
// These functions are based on the sample implementation in golang.org/src/crypto/subtle/constant_time.go
// Here we have refactored these functions for uint32 arithmetic and to avoid extra shifts and casts

// ConstantTimeSelect returns x if v is 1 and y if v is 0.
// Its behavior is undefined if v takes any other value.
func SelectUint32(v, x, y uint32) uint32 { return ^(v-1)&x | (v-1)&y }

// ConstantTimeLessThan returns 1 if x < y and 0 otherwise.
// It works by checking the most significant bit, and then testing the rest of the bits by casting to int32
func LessThanUint32(x, y uint32) uint32 {
	diff := int64(x) - int64(y)
	return uint32((diff >> 63) & 1)
}

func LessOrEqUint32(x, y uint32) uint32 {
	diff := int64(x) - int64(y)
	return uint32(((diff - 1) >> 63) & 1)
}

// ConstantTimeEq returns 1 if x == y and 0 otherwise.
func EqUint32(x, y uint32) uint32 {
	diff := int64(x) - int64(y)
	return uint32((((diff - 1) ^ diff) >> 63) & 1)
}

// ConstantTimeEq returns 1 if x == y and 0 otherwise.
func NotEqUint32(x, y uint32) uint32 {
	diff := int64(x) - int64(y)
	return uint32((((-diff) | diff) >> 63) & 1)
}

// These functions contain the logic without the shift and cast to uint32, which can be left for last in a chain of logic
// The single logic bit is the msb in the int64, or the sign bit

func Shift63CastUint32(x int64) uint32 {
	return uint32((x >> 63) & 1)
}

// EqMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func EqMSBInt64(x, y int64) int64 {
	return EqZeroMSBInt64(x - y)
}

// EqMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func NotEqMSBInt64(x, y int64) int64 {
	return NotEqZeroMSBInt64(x - y)
}

// LessThanMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func LessThanMSBInt64(x, y int64) int64 {
	return x - y
}

// LessOrEqMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func LessOrEqMSBInt64(x, y int64) int64 {
	return x - y - 1
}

// EqZeroMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func EqZeroMSBInt64(x int64) int64 {
	return (x - 1) ^ x
}

// EqZeroMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func NotEqZeroMSBInt64(x int64) int64 {
	return x | (-x)
}

// LessThanZeroMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func LessThanZeroMSBInt64(x int64) int64 {
	return x
}

// LessOrEqZeroMSBInt64 returns an int64 whose most significant bit is 1 if the inputs are equal and 0 otherwise.
// The behavior of the other 63 bits is undefined.
func LessOrEqZeroMSBInt64(x int64) int64 {
	return x - 1
}
