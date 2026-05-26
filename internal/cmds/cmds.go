package cmds

const (
	optInTag     = uint16(1 << 15)
	blockTag     = uint16(1 << 14)
	readonly     = uint16(1<<13) | retryableTag
	noRetTag     = uint16(1<<12) | readonly | pipeTag // make noRetTag can also be retried and auto pipelining
	mtGetTag     = uint16(1<<11) | readonly           // make mtGetTag can also be retried
	scrRoTag     = uint16(1<<10) | readonly           // make scrRoTag can also be retried
	unsubTag     = uint16(1<<9) | noRetTag
	pipeTag      = uint16(1 << 8) // make blocking mode request can use auto pipelining
	retryableTag = uint16(1 << 7) // make command retryable
	// InitSlot indicates that the command be sent to any redis node in cluster
	InitSlot = uint16(1 << 14)
	// NoSlot indicates that the command has no key slot specified
	NoSlot = uint16(1 << 15)
)

var (
	// OptInCmd is predefined CLIENT CACHING YES
	OptInCmd = Completed{
		cs: newCommandSlice([]string{"CLIENT", "CACHING", "YES"}),
		cf: optInTag,
	}
	// OptInNopCmd is a predefined alternative for CLIENT CACHING YES in BCAST/OPTOUT mode.
	OptInNopCmd = Completed{
		cs: newCommandSlice([]string{"ECHO", ""}),
		cf: optInTag,
	}
	// MultiCmd is predefined MULTI
	MultiCmd = Completed{
		cs: newCommandSlice([]string{"MULTI"}),
	}
	// ExecCmd is predefined EXEC
	ExecCmd = Completed{
		cs: newCommandSlice([]string{"EXEC"}),
	}
	// RoleCmd is predefined ROLE
	RoleCmd = Completed{
		cs: newCommandSlice([]string{"ROLE"}),
		cf: pipeTag,
	}

	// UnsubscribeCmd is predefined UNSUBSCRIBE
	UnsubscribeCmd = Completed{
		cs: newCommandSlice([]string{"UNSUBSCRIBE"}),
		cf: unsubTag,
	}
	// PUnsubscribeCmd is predefined PUNSUBSCRIBE
	PUnsubscribeCmd = Completed{
		cs: newCommandSlice([]string{"PUNSUBSCRIBE"}),
		cf: unsubTag,
	}
	// SUnsubscribeCmd is predefined SUNSUBSCRIBE
	SUnsubscribeCmd = Completed{
		cs: newCommandSlice([]string{"SUNSUBSCRIBE"}),
		cf: unsubTag,
	}
	// PingCmd is predefined PING
	PingCmd = Completed{
		cs: newCommandSlice([]string{"PING"}),
	}
	// SlotCmd is predefined CLUSTER SLOTS
	SlotCmd = Completed{
		cs: newCommandSlice([]string{"CLUSTER", "SLOTS"}),
		cf: pipeTag,
	}
	// ShardsCmd is predefined CLUSTER SHARDS
	ShardsCmd = Completed{
		cs: newCommandSlice([]string{"CLUSTER", "SHARDS"}),
		cf: pipeTag,
	}
	// AskingCmd is predefined CLUSTER ASKING
	AskingCmd = Completed{
		cs: newCommandSlice([]string{"ASKING"}),
	}
	// SentinelSubscribe is predefined SUBSCRIBE ASKING
	SentinelSubscribe = Completed{
		cs: newCommandSlice([]string{"SUBSCRIBE", "+sentinel", "+slave", "-sdown", "+sdown", "+switch-master", "+reboot"}),
		cf: noRetTag,
	}
	// SentinelUnSubscribe is predefined UNSUBSCRIBE ASKING
	SentinelUnSubscribe = Completed{
		cs: newCommandSlice([]string{"UNSUBSCRIBE", "+sentinel", "+slave", "-sdown", "+sdown", "+switch-master", "+reboot"}),
		cf: unsubTag,
	}
	// ClientTrackingOffCmd is predefined CLIENT TRACKING OFF
	ClientTrackingOffCmd = Completed{
		cs: newCommandSlice([]string{"CLIENT", "TRACKING", "OFF"}),
	}

	// DiscardCmd is predefined DISCARD
	DiscardCmd = Completed{
		cs: newCommandSlice([]string{"DISCARD"}),
	}
)

// ToBlock marks the command with blockTag
func ToBlock(c *Completed) {
	_ = "STUB: not implemented"

	// Incomplete represents an incomplete Redis command. It should then be completed by calling Build().
	return
}

type Incomplete struct {
	cs *CommandSlice
	cf int16 // use int16 instead of uint16 to make a difference with Completed
	ks uint16
}

// Completed represents a completed Redis command, should be created by the Build() of command builder.
type Completed struct {
	cs *CommandSlice
	cf uint16 // cmd flag
	ks uint16 // key slot
}

// Pin prevents a Completed to be recycled
func (c Completed) Pin() Completed {
	_ = "STUB: not implemented"
	return *

	// ToPipe returns a new command with pipeTag
	new(Completed)
}

func (c Completed) ToPipe() Completed {
	_ = "STUB: not implemented"
	return *

	// ToRetryable return a new command with retryableTag
	new(Completed)
}

func (c Completed) ToRetryable() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// IsEmpty checks if it is an empty command.
func (c *Completed) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// IsOptIn checks if it is client side caching opt-int command.
func (c *Completed) IsOptIn() bool { _ = "STUB: not implemented"; return false }

// IsBlock checks if it is blocking command which needs to be process by dedicated connection.
func (c *Completed) IsBlock() bool { _ = "STUB: not implemented"; return false }

// NoReply checks if it is one of the SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE or PUNSUBSCRIBE commands.
func (c *Completed) NoReply() bool { _ = "STUB: not implemented"; return false }

// IsUnsub checks if it is one of the UNSUBSCRIBE, PUNSUBSCRIBE, or SUNSUBSCRIBE commands.
func (c *Completed) IsUnsub() bool { _ = "STUB: not implemented"; return false }

// IsReadOnly checks if it is readonly command and can be retried when network error.
func (c *Completed) IsReadOnly() bool { _ = "STUB: not implemented"; return false }

// IsWrite checks if it is not readonly command.
func (c *Completed) IsWrite() bool { _ = "STUB: not implemented"; return false }

// IsPipe checks if it is set pipeTag which prefers auto pipelining
func (c *Completed) IsPipe() bool { _ = "STUB: not implemented"; return false }

// IsRetryable checks if it is set retryableTag
func (c *Completed) IsRetryable() bool { _ = "STUB: not implemented"; return false }

// Commands returns the commands as []string.
// Note that the returned []string should not be modified
// and should not be read after passing into the Client interface, because it will be recycled.
func (c *Completed) Commands() []string {
	_ = "STUB: not implemented"

	// Slot returns the command key slot
	return nil
}

func (c *Completed) Slot() uint16 {
	_ = "STUB: not implemented"

	// SetSlot returns a new completed command with its key slot be overridden
	return 0
}

func (c Completed) SetSlot(key string) Completed { _ = "STUB: not implemented"; return *new(Completed) }

var Slot = slot

// Cacheable represents a completed Redis command which supports server-assisted client side caching,
// and it should be created by the Cache() of command builder.
type Cacheable Completed

// Pin prevents a Cacheable to be recycled
func (c Cacheable) Pin() Cacheable {
	_ = "STUB: not implemented"
	return *

	// Slot returns the command key slot
	new(Cacheable)
}

func (c *Cacheable) Slot() uint16 {
	_ = "STUB: not implemented"

	// Commands returns the commands as []string.
	// Note that the returned []string should not be modified
	// and should not be read after passing into the Client interface, because it will be recycled.
	return 0
}

func (c *Cacheable) Commands() []string {
	_ = "STUB: not implemented"

	// IsMGet returns if the command is MGET
	return nil
}

func (c *Cacheable) IsMGet() bool { _ = "STUB: not implemented"; return false }

// MGetCacheCmd returns the cache command of the MGET singular command
func MGetCacheCmd(c Cacheable) string { _ = "STUB: not implemented"; return "" }

// MGetCacheKey returns the cache key of the MGET singular command
func MGetCacheKey(c Cacheable, i int) string {
	_ = "STUB: not implemented"

	// CacheKey returns the cache key used by the server-assisted client side caching
	return ""
}

func CacheKey(c Cacheable) (key, command string) { _ = "STUB: not implemented"; return "", "" }

// AppendCompleted appends an arg to a Completed
func AppendCompleted(c Completed, s string) { _ = "STUB: not implemented"; return }

// CompletedCS get the underlying *CommandSlice
func CompletedCS(c Completed) *CommandSlice {
	_ = "STUB: not implemented"

	// CacheableCS get the underlying *CommandSlice
	return nil
}

func CacheableCS(c Cacheable) *CommandSlice {
	_ = "STUB: not implemented"

	// NewCompleted creates an arbitrary Completed command.
	return nil
}

func NewCompleted(ss []string) Completed { _ = "STUB: not implemented"; return *new(Completed) }

// NewBlockingCompleted creates an arbitrary blocking Completed command.
func NewBlockingCompleted(ss []string) Completed { _ = "STUB: not implemented"; return *new(Completed) }

// NewReadOnlyCompleted creates an arbitrary readonly Completed command.
func NewReadOnlyCompleted(ss []string) Completed { _ = "STUB: not implemented"; return *new(Completed) }

// NewMGetCompleted creates an arbitrary readonly Completed command.
func NewMGetCompleted(ss []string) Completed { _ = "STUB: not implemented"; return *new(Completed) }

// MGets groups keys by their slot and returns multi MGET commands
func MGets(keys []string) map[uint16]Completed { _ = "STUB: not implemented"; return nil }

// MDels groups keys by their slot and returns multi DEL commands
func MDels(keys []string) map[uint16]Completed { _ = "STUB: not implemented"; return nil }

// MSets groups keys by their slot and returns multi MSET commands
func MSets(kvs map[string]string) map[uint16]Completed { _ = "STUB: not implemented"; return nil }

// MSetNXs groups keys by their slot and returns multi MSETNX commands
func MSetNXs(kvs map[string]string) map[uint16]Completed { _ = "STUB: not implemented"; return nil }

// JsonMGets groups keys by their slot and returns multi JSON.MGET commands
func JsonMGets(keys []string, path string) map[uint16]Completed {
	_ = "STUB: not implemented"
	return nil
}

// JsonMSets groups keys by their slot and returns multi JSON.MSET commands
func JsonMSets(kvs map[string]string, path string) map[uint16]Completed {
	_ = "STUB: not implemented"
	return nil
}

func slotMCMDs(cmd string, keys []string, cf uint16) map[uint16]Completed {
	_ = "STUB: not implemented"
	return nil
}

func slotMSets(cmd string, kvs map[string]string) map[uint16]Completed {
	_ = "STUB: not implemented"
	return nil
}

// NewMultiCompleted creates multiple arbitrary Completed commands.
func NewMultiCompleted(cs [][]string) []Completed { _ = "STUB: not implemented"; return nil }

func check(prev, new uint16) uint16 { _ = "STUB: not implemented"; return 0 }

const multiKeySlotErr = "multi key command with different key slots are not allowed"
const multiKeyCacheErr = "client side caching for scripting only supports numkeys=1"
