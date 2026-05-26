package rueidiscompatmock

import (
	"time"
)

type ExpectedString struct{ exp *expectation }

func (e *ExpectedString) SetVal(v string) *ExpectedString { _ = "STUB: not implemented"; return nil }

func (e *ExpectedString) SetErr(err error) *ExpectedString { _ = "STUB: not implemented"; return nil }

func (e *ExpectedString) RedisNil() *ExpectedString { _ = "STUB: not implemented"; return nil }

type ExpectedStatus = ExpectedString

type ExpectedBool struct{ exp *expectation }

func (e *ExpectedBool) SetVal(v bool) *ExpectedBool { _ = "STUB: not implemented"; return nil }

func (e *ExpectedBool) SetErr(err error) *ExpectedBool { _ = "STUB: not implemented"; return nil }

type ExpectedInt struct{ exp *expectation }

func (e *ExpectedInt) SetVal(v int64) *ExpectedInt { _ = "STUB: not implemented"; return nil }

func (e *ExpectedInt) SetErr(err error) *ExpectedInt { _ = "STUB: not implemented"; return nil }

type ExpectedDuration struct {
	exp       *expectation
	precision time.Duration
}

func (e *ExpectedDuration) SetVal(v time.Duration) *ExpectedDuration {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedDuration) SetErr(err error) *ExpectedDuration {
	_ = "STUB: not implemented"
	return nil
}

type ExpectedStringSlice struct{ exp *expectation }

func (e *ExpectedStringSlice) SetVal(v []string) *ExpectedStringSlice {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedStringSlice) SetErr(err error) *ExpectedStringSlice {
	_ = "STUB: not implemented"
	return nil
}

type ExpectedSlice struct{ exp *expectation }

func (e *ExpectedSlice) SetVal(v []any) *ExpectedSlice { _ = "STUB: not implemented"; return nil }

func (e *ExpectedSlice) SetErr(err error) *ExpectedSlice { _ = "STUB: not implemented"; return nil }

type ExpectedStringStringMap struct{ exp *expectation }

func (e *ExpectedStringStringMap) SetVal(v map[string]string) *ExpectedStringStringMap {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedStringStringMap) SetErr(err error) *ExpectedStringStringMap {
	_ = "STUB: not implemented"
	return nil
}

type ExpectedCmd struct{ exp *expectation }

func (e *ExpectedCmd) SetVal(v string) *ExpectedCmd { _ = "STUB: not implemented"; return nil }

func (e *ExpectedCmd) SetValInt(v int64) *ExpectedCmd { _ = "STUB: not implemented"; return nil }

func (e *ExpectedCmd) SetErr(err error) *ExpectedCmd { _ = "STUB: not implemented"; return nil }

func (e *ExpectedCmd) RedisNil() *ExpectedCmd { _ = "STUB: not implemented"; return nil }
