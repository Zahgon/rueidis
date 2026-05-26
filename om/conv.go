package om

import (
	"reflect"
	"strconv"
	"unsafe"

	"github.com/redis/rueidis"
)

func newHashConvFactory(t reflect.Type, schema schema) *hashConvFactory {
	_ = "STUB: not implemented"
	return nil
}

type hashConvFactory struct {
	fields map[string]fieldConv
}

type fieldConv struct {
	conv converter
	idx  int
}

func (f hashConvFactory) NewConverter(entity reflect.Value) hashConv {
	_ = "STUB: not implemented"
	return *new(hashConv)
}

type hashConv struct {
	factory hashConvFactory
	entity  reflect.Value
}

func (r hashConv) ToHash() (fields map[string]string) { _ = "STUB: not implemented"; return nil }

func (r hashConv) FromHash(fields map[string]string) error { _ = "STUB: not implemented"; return nil }

type converter struct {
	ValueToString func(value reflect.Value) (string, bool)
	StringToValue func(value string) (reflect.Value, error)
}

var converters = struct {
	val   map[reflect.Kind]converter
	ptr   map[reflect.Kind]converter
	slice map[reflect.Kind]converter
}{
	ptr: map[reflect.Kind]converter{
		reflect.Int64: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.IsNil() {
					return "", false
				}
				return strconv.FormatInt(value.Elem().Int(), 10), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				v, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return reflect.Value{}, err
				}
				return reflect.ValueOf(&v), nil
			},
		},
		reflect.String: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.IsNil() {
					return "", false
				}
				return value.Elem().String(), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(&value), nil
			},
		},
		reflect.Bool: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.IsNil() {
					return "", false
				}
				if value.Elem().Bool() {
					return "t", true
				}
				return "f", true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				b := value == "t"
				return reflect.ValueOf(&b), nil
			},
		},
		reflect.Struct: {
			ValueToString: nil,
			StringToValue: nil,
		},
	},
	val: map[reflect.Kind]converter{
		reflect.Int64: {
			ValueToString: func(value reflect.Value) (string, bool) {
				return strconv.FormatInt(value.Int(), 10), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				v, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return reflect.Value{}, err
				}
				return reflect.ValueOf(v), nil
			},
		},
		reflect.String: {
			ValueToString: func(value reflect.Value) (string, bool) {
				return value.String(), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(value), nil
			},
		},
		reflect.Bool: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.Bool() {
					return "t", true
				}
				return "f", true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				b := value == "t"
				return reflect.ValueOf(b), nil
			},
		},
		reflect.Struct: {
			ValueToString: nil,
			StringToValue: nil,
		},
	},
	slice: map[reflect.Kind]converter{
		reflect.Uint8: {
			ValueToString: func(value reflect.Value) (string, bool) {
				return rueidis.BinaryString(value.Bytes()), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				buf := unsafe.Slice(unsafe.StringData(value), len(value))
				return reflect.ValueOf(buf), nil
			},
		},
		reflect.Float32: {
			ValueToString: func(value reflect.Value) (string, bool) {
				vs, ok := value.Interface().([]float32)
				return rueidis.VectorString32(vs), ok
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(rueidis.ToVector32(value)), nil
			},
		},
		reflect.Float64: {
			ValueToString: func(value reflect.Value) (string, bool) {
				vs, ok := value.Interface().([]float64)
				return rueidis.VectorString64(vs), ok
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(rueidis.ToVector64(value)), nil
			},
		},
		reflect.Struct: {
			ValueToString: nil,
			StringToValue: nil,
		},
	},
}
