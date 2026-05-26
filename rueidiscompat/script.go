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
	"context"
)

type Scripter interface {
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
	ScriptLoad(ctx context.Context, script string) *StringCmd
}

var (
	_ Scripter = (*Compat)(nil)
)

type Script struct {
	src, hash string
}

func NewScript(src string) *Script { _ = "STUB: not implemented"; return nil }

func (s *Script) Hash() string { _ = "STUB: not implemented"; return "" }

func (s *Script) Load(ctx context.Context, c Scripter) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) Exists(ctx context.Context, c Scripter) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) Eval(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) EvalRO(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) EvalSha(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) EvalShaRO(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

// Run optimistically uses EVALSHA to run the script. If the script does not exist,
// it is retried using EVAL.
func (s *Script) Run(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

// RunRO optimistically uses EVALSHA_RO to run the script. If the script does not exist,
// it is retried using EVAL_RO.
func (s *Script) RunRO(ctx context.Context, c Scripter, keys []string, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}
