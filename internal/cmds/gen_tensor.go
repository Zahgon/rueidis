// Code generated DO NOT EDIT

package cmds

type AiTensorget Incomplete

func (b Builder) AiTensorget() (c AiTensorget) { _ = "STUB: not implemented"; return *new(AiTensorget) }

func (c AiTensorget) Key(key string) AiTensorgetKey {
	_ = "STUB: not implemented"
	return *new(AiTensorgetKey)
}

type AiTensorgetFormatBlob Incomplete

func (c AiTensorgetFormatBlob) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiTensorgetFormatBlob) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiTensorgetFormatValues Incomplete

func (c AiTensorgetFormatValues) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c AiTensorgetFormatValues) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type AiTensorgetKey Incomplete

func (c AiTensorgetKey) Meta() AiTensorgetMeta {
	_ = "STUB: not implemented"
	return *new(AiTensorgetMeta)
}

type AiTensorgetMeta Incomplete

func (c AiTensorgetMeta) Blob() AiTensorgetFormatBlob {
	_ = "STUB: not implemented"
	return *new(AiTensorgetFormatBlob)
}

func (c AiTensorgetMeta) Values() AiTensorgetFormatValues {
	_ = "STUB: not implemented"
	return *new(AiTensorgetFormatValues)
}

func (c AiTensorgetMeta) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiTensorgetMeta) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiTensorset Incomplete

func (b Builder) AiTensorset() (c AiTensorset) { _ = "STUB: not implemented"; return *new(AiTensorset) }

func (c AiTensorset) Key(key string) AiTensorsetKey {
	_ = "STUB: not implemented"
	return *new(AiTensorsetKey)
}

type AiTensorsetBlob Incomplete

func (c AiTensorsetBlob) Values(value ...string) AiTensorsetValues {
	_ = "STUB: not implemented"
	return *new(AiTensorsetValues)
}

func (c AiTensorsetBlob) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiTensorsetKey Incomplete

func (c AiTensorsetKey) Float() AiTensorsetTypeFloat {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeFloat)
}

func (c AiTensorsetKey) Double() AiTensorsetTypeDouble {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeDouble)
}

func (c AiTensorsetKey) Int8() AiTensorsetTypeInt8 {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeInt8)
}

func (c AiTensorsetKey) Int16() AiTensorsetTypeInt16 {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeInt16)
}

func (c AiTensorsetKey) Int32() AiTensorsetTypeInt32 {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeInt32)
}

func (c AiTensorsetKey) Int64() AiTensorsetTypeInt64 {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeInt64)
}

func (c AiTensorsetKey) Uint8() AiTensorsetTypeUint8 {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeUint8)
}

func (c AiTensorsetKey) Uint16() AiTensorsetTypeUint16 {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeUint16)
}

func (c AiTensorsetKey) String() AiTensorsetTypeString {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeString)
}

func (c AiTensorsetKey) Bool() AiTensorsetTypeBool {
	_ = "STUB: not implemented"
	return *new(AiTensorsetTypeBool)
}

type AiTensorsetShape Incomplete

func (c AiTensorsetShape) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

func (c AiTensorsetShape) Blob(blob string) AiTensorsetBlob {
	_ = "STUB: not implemented"
	return *new(AiTensorsetBlob)
}

func (c AiTensorsetShape) Values(value ...string) AiTensorsetValues {
	_ = "STUB: not implemented"
	return *new(AiTensorsetValues)
}

func (c AiTensorsetShape) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiTensorsetTypeBool Incomplete

func (c AiTensorsetTypeBool) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeDouble Incomplete

func (c AiTensorsetTypeDouble) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeFloat Incomplete

func (c AiTensorsetTypeFloat) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeInt16 Incomplete

func (c AiTensorsetTypeInt16) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeInt32 Incomplete

func (c AiTensorsetTypeInt32) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeInt64 Incomplete

func (c AiTensorsetTypeInt64) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeInt8 Incomplete

func (c AiTensorsetTypeInt8) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeString Incomplete

func (c AiTensorsetTypeString) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeUint16 Incomplete

func (c AiTensorsetTypeUint16) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetTypeUint8 Incomplete

func (c AiTensorsetTypeUint8) Shape(shape ...int64) AiTensorsetShape {
	_ = "STUB: not implemented"
	return *new(AiTensorsetShape)
}

type AiTensorsetValues Incomplete

func (c AiTensorsetValues) Values(value ...string) AiTensorsetValues {
	_ = "STUB: not implemented"
	return *new(AiTensorsetValues)
}

func (c AiTensorsetValues) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
