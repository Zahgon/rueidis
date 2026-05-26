package rueidis

import (
	"context"
	"sync"
)

// LuaOption is a functional option for configuring Lua script behavior.
type LuaOption func(*Lua)

// WithLoadSHA1 allows enabling loading of SHA-1 from Redis via SCRIPT LOAD instead of calculating
// it on the client side. When enabled, the SHA-1 hash is not calculated client-side (important
// for FIPS compliance). Instead, on first execution, SCRIPT LOAD is called to obtain the SHA-1
// from Redis, which is then used for EVALSHA commands in subsequent executions.
func WithLoadSHA1(enabled bool) LuaOption { _ = "STUB: not implemented"; return *new(LuaOption) }

// NewLuaScript creates a Lua instance whose Lua.Exec uses EVALSHA and EVAL.
// By default, SHA-1 is calculated client-side. Use WithLoadSHA1(true) option to load SHA-1 from Redis instead.
func NewLuaScript(script string, opts ...LuaOption) *Lua { _ = "STUB: not implemented"; return nil }

// NewLuaScriptReadOnly creates a Lua instance whose Lua.Exec uses EVALSHA_RO and EVAL_RO.
// By default, SHA-1 is calculated client-side. Use WithLoadSHA1(true) option to load SHA-1 from Redis instead.
func NewLuaScriptReadOnly(script string, opts ...LuaOption) *Lua {
	_ = "STUB: not implemented"
	return nil
}

// NewLuaScriptNoSha creates a Lua instance whose Lua.Exec uses EVAL only (never EVALSHA).
// No SHA-1 is calculated or loaded. The script is sent to the server every time. Use this when you want
// to avoid SHA-1 entirely (e.g., to fully avoid hash collision concerns).
func NewLuaScriptNoSha(script string) *Lua { _ = "STUB: not implemented"; return nil }

// NewLuaScriptReadOnlyNoSha creates a Lua instance whose Lua.Exec uses EVAL_RO only (never EVALSHA_RO).
// No SHA-1 is calculated or loaded. The script is sent to the server every time. Use this when you want
// to avoid SHA-1 entirely (e.g., to fully avoid hash collision concerns).
func NewLuaScriptReadOnlyNoSha(script string) *Lua { _ = "STUB: not implemented"; return nil }

// NewLuaScriptRetryable creates a retryable Lua instance whose Lua.Exec uses EVALSHA and EVAL.
// By default, SHA-1 is calculated client-side. Use WithLoadSHA1(true) option to load SHA-1 from Redis instead.
func NewLuaScriptRetryable(script string, opts ...LuaOption) *Lua {
	_ = "STUB: not implemented"
	return nil
}

// NewLuaScriptNoShaRetryable creates a retryable Lua instance whose Lua.Exec uses EVAL only (never EVALSHA).
// No SHA-1 is calculated or loaded. The script is sent to the server every time. Use this when you want
// to avoid SHA-1 entirely (e.g., to fully avoid hash collision concerns).
func NewLuaScriptNoShaRetryable(script string) *Lua { _ = "STUB: not implemented"; return nil }

func newLuaScript(script string, readonly bool, noSha1, retryable bool, opts ...LuaOption) *Lua {
	_ = "STUB: not implemented"
	return nil
}

// It's important to avoid calling sha1 methods where not needed since Go will panic in FIPS mode.

// Lua represents a redis lua script. It should be created from the NewLuaScript() or NewLuaScriptReadOnly().
type Lua struct {
	script    string
	sha1      string
	maxp      int
	sha1Mu    sync.RWMutex
	readonly  bool
	noSha1    bool
	loadSha1  bool
	retryable bool
}

// Exec the script to the given Client.
// It will first try with the EVALSHA/EVALSHA_RO and then EVAL/EVAL_RO if the first try failed.
// If Lua is initialized with disabled SHA1, it will use EVAL/EVAL_RO without the EVALSHA/EVALSHA_RO attempt.
// If Lua is initialized with SHA-1 loading, it will call SCRIPT LOAD once to obtain the SHA-1 from Redis.
// Cross-slot keys are prohibited if the Client is a cluster client.
func (s *Lua) Exec(ctx context.Context, c Client, keys, args []string) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

// Determine which SHA-1 to use.

// Check if SHA-1 is already loaded.

// the double check

// NoSha constructors: always use EVAL, never EVALSHA.
// Regular constructors: use EVALSHA if SHA-1 is available, fall back to EVAL on NOSCRIPT error.

// LuaExec is a single execution unit of Lua.ExecMulti.
type LuaExec struct {
	Keys []string
	Args []string
}

// ExecMulti exec the script multiple times by the provided LuaExec to the given Client.
// For regular constructors, it will SCRIPT LOAD to all redis nodes and then use EVALSHA/EVALSHA_RO.
// For NoSha constructors, it will use EVAL/EVAL_RO only without any script loading.
// Cross-slot keys within the single LuaExec are prohibited if the Client is a cluster client.
func (s *Lua) ExecMulti(ctx context.Context, c Client, multi ...LuaExec) (resp []RedisResult) {
	_ = "STUB: not implemented"
	return nil

	// For regular constructors (not NoSha), load the script to all nodes.
}

// Store the first successful SHA-1 result for cases when sha1 loading on.

// Set SHA-1 from Redis if sha1 loading is enabled.

// NoSha constructors: always use EVAL.
// Regular constructors: use EVALSHA if SHA-1 is available.

func (s *Lua) mayRetryable(cmd Completed) Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}
