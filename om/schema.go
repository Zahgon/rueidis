package om

import (
	"reflect"
)

const ignoreField = "-"

type schema struct {
	key     *field
	ver     *field
	ext     *field
	fields  map[string]*field
	verless bool
}

type field struct {
	typ   reflect.Type
	name  string
	idx   int
	isKey bool
	isVer bool
	isExt bool
}

func newSchema(t reflect.Type) schema { _ = "STUB: not implemented"; return *new(schema) }

// ver is no longer required

func parse(f reflect.StructField) (field field) { _ = "STUB: not implemented"; return *new(field) }

func key(prefix, id string) (key string) { _ = "STUB: not implemented"; return "" }
