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
	"sync"
)

// structMap contains the map of struct fields for target structs
// indexed by the struct type.
type structMap struct {
	m sync.Map
}

func newStructMap() *structMap { _ = "STUB: not implemented"; return nil }

func (s *structMap) get(t reflect.Type) *structSpec { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------------------------------

// structSpec contains the list of all fields in a target struct.
type structSpec struct {
	m map[string]*structField
}

func (s *structSpec) set(tag string, sf *structField) { _ = "STUB: not implemented"; return }

func newStructSpec(t reflect.Type, fieldTag string) *structSpec {
	_ = "STUB: not implemented"
	return nil
}

// Added a check for Pointer here. If it's a Pointer, use the built-in decoder of the element.
// This works, because in Scan() #129-131
// if isPtr && v.IsNil() {
//     v.Set(reflect.New(v.Type().Elem()))
// }
// A new value is set

// Use the built-in decoder.

//------------------------------------------------------------------------------

// structField represents a single field in a target struct.
type structField struct {
	fn    decoderFunc
	index int
}

//------------------------------------------------------------------------------

type StructValue struct {
	spec  *structSpec
	value reflect.Value
}

func (s StructValue) Scan(key string, value string) error { _ = "STUB: not implemented"; return nil }
