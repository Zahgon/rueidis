// Copyright (c) 2013 The github.com/go-redis/redis Authors.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package rueidiscompat

import (
	"reflect"
)

// decoderFunc represents decoding functions for default built-in types.
type decoderFunc func(reflect.Value, string) error

// Scanner is the interface implemented by themselves,
// which will override the decoding behavior of decoderFunc.
type Scanner interface {
	ScanRedis(s string) error
}

var (
	// List of built-in decoders indexed by their numeric constant values (e.g.: reflect.Bool = 1).
	decoders = []decoderFunc{
		reflect.Bool:          decodeBool,
		reflect.Int:           decodeInt,
		reflect.Int8:          decodeInt8,
		reflect.Int16:         decodeInt16,
		reflect.Int32:         decodeInt32,
		reflect.Int64:         decodeInt64,
		reflect.Uint:          decodeUint,
		reflect.Uint8:         decodeUint8,
		reflect.Uint16:        decodeUint16,
		reflect.Uint32:        decodeUint32,
		reflect.Uint64:        decodeUint64,
		reflect.Float32:       decodeFloat32,
		reflect.Float64:       decodeFloat64,
		reflect.Complex64:     decodeUnsupported,
		reflect.Complex128:    decodeUnsupported,
		reflect.Array:         decodeUnsupported,
		reflect.Chan:          decodeUnsupported,
		reflect.Func:          decodeUnsupported,
		reflect.Interface:     decodeUnsupported,
		reflect.Map:           decodeUnsupported,
		reflect.Ptr:           decodeUnsupported,
		reflect.Slice:         decodeSlice,
		reflect.String:        decodeString,
		reflect.Struct:        decodeUnsupported,
		reflect.UnsafePointer: decodeUnsupported,
	}

	// Global map of struct field specs that is populated once for every new
	// struct type that is scanned. This caches the field types and the corresponding
	// decoder functions to avoid iterating through struct fields on subsequent scans.
	globalStructMap = newStructMap()
)

func Struct(dst interface{}) (StructValue, error) {
	_ = "STUB: not implemented"
	return *

	// The destination to scan into should be a struct pointer.
	new(StructValue), nil
}

// Scan scans the results from a key-value Redis map result set to a destination struct.
// The Redis keys are matched to the struct's field with the `redis` tag.
// NOTE: vals' element's underlying type should be string
func Scan(dst interface{}, keys []string, vals []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Iterate through the (key, value) sequence.

func decodeBool(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeInt8(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeInt16(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeInt32(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeInt64(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeInt(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeNumber(f reflect.Value, s string, bitSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeUint8(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeUint16(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeUint32(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeUint64(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeUint(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeUnsignedNumber(f reflect.Value, s string, bitSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeFloat32(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

// although the default is float64, but we better define it.
func decodeFloat64(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeString(f reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func decodeSlice(f reflect.Value, s string) error {
	_ = "STUB: not implemented"
	// []byte slice ([]uint8).
	return nil
}

func decodeUnsupported(v reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }
