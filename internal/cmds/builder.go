package cmds

import (
	"sync"
)

const ErrBuiltTwice = "a command should not be built twice"
const ErrUnfinished = "a command should be finished by calling Build() or Cache()"

var pool = &sync.Pool{New: func() any {
	return &CommandSlice{s: make([]string, 0, 2), l: -1}
}}

// CommandSlice is the command container managed by the sync.Pool
type CommandSlice struct {
	s []string
	l int32
	r int32
}

func (cs *CommandSlice) Build() { _ = "STUB: not implemented"; return }

func (cs *CommandSlice) Verify() { _ = "STUB: not implemented"; return }

func newCommandSlice(s []string) *CommandSlice { _ = "STUB: not implemented"; return nil }

// NewBuilder creates a Builder and initializes the internal sync.Pool
func NewBuilder(initSlot uint16) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Builder builds commands by reusing CommandSlice from the sync.Pool
type Builder struct {
	ks uint16
}

func get() *CommandSlice { _ = "STUB: not implemented"; return nil }

// PutCompletedForce recycles the Completed regardless of the c.cs.r
func PutCompletedForce(c Completed) {
	_ = "STUB: not implemented"

	// PutCacheableForce recycles the Cacheable regardless of the c.cs.r
	return
}

func PutCacheableForce(c Cacheable) {
	_ = "STUB: not implemented"

	// PutCompleted recycles the Completed
	return
}

func PutCompleted(c Completed) { _ = "STUB: not implemented"; return }

// PutCacheable recycles the Cacheable
func PutCacheable(c Cacheable) { _ = "STUB: not implemented"; return }

// Arbitrary allows user to build an arbitrary redis command with Builder.Arbitrary
type Arbitrary Completed

// Arbitrary allows user to build an arbitrary redis command by following Arbitrary.Keys and Arbitrary.Args
func (b Builder) Arbitrary(token ...string) (c Arbitrary) {
	_ = "STUB: not implemented"
	return *new(Arbitrary)
}

// Keys calculate which key slot the command belongs to.
// Users must use Keys to construct the key part of the command; otherwise,
// the command will not be sent to correct redis node.
func (c Arbitrary) Keys(keys ...string) Arbitrary {
	_ = "STUB: not implemented"
	return *new(Arbitrary)
}

// Args is used to construct non-key parts of the command.
func (c Arbitrary) Args(args ...string) Arbitrary {
	_ = "STUB: not implemented"
	return *new(Arbitrary)
}

// Build is used to complete constructing a command
func (c Arbitrary) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// Blocking is used to complete constructing a command and mark it as blocking command.
// Blocking command will occupy a connection from a separated connection pool.
func (c Arbitrary) Blocking() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// ReadOnly is used to complete constructing a command and mark it as readonly command.
// ReadOnly will be retried under network issues.
func (c Arbitrary) ReadOnly() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// MultiGet is used to complete constructing a command and mark it as mtGetTag command.
func (c Arbitrary) MultiGet() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// IsZero is used to test if Arbitrary is initialized
func (c Arbitrary) IsZero() bool { _ = "STUB: not implemented"; return false }

var (
	arbitraryNoCommand = "Arbitrary should be provided with redis command"
	arbitrarySubscribe = "Arbitrary does not support SUBSCRIBE/UNSUBSCRIBE"
	arbitraryMultiGet  = "Arbitrary.MultiGet is only valid for MGET and JSON.MGET"
)
