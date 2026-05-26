//go:build go1.23

package cmds

import (
	"iter"
)

func (c HmsetFieldValue) FieldValueIter(seq iter.Seq2[string, string]) HmsetFieldValue {
	_ = "STUB: not implemented"
	return *new(HmsetFieldValue)
}

func (c HsetFieldValue) FieldValueIter(seq iter.Seq2[string, string]) HsetFieldValue {
	_ = "STUB: not implemented"
	return *new(HsetFieldValue)
}

func (c HsetexFieldValue) FieldValueIter(seq iter.Seq2[string, string]) HsetexFieldValue {
	_ = "STUB: not implemented"
	return *new(HsetexFieldValue)
}

func (c XaddFieldValue) FieldValueIter(seq iter.Seq2[string, string]) XaddFieldValue {
	_ = "STUB: not implemented"
	return *new(XaddFieldValue)
}

func (c ZaddScoreMember) ScoreMemberIter(seq iter.Seq2[string, float64]) ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}
