// Code generated DO NOT EDIT

package cmds

type Sadd Incomplete

func (b Builder) Sadd() (c Sadd) { _ = "STUB: not implemented"; return *new(Sadd) }

func (c Sadd) Key(key string) SaddKey { _ = "STUB: not implemented"; return *new(SaddKey) }

type SaddKey Incomplete

func (c SaddKey) Member(member ...string) SaddMember {
	_ = "STUB: not implemented"
	return *new(SaddMember)
}

type SaddMember Incomplete

func (c SaddMember) Member(member ...string) SaddMember {
	_ = "STUB: not implemented"
	return *new(SaddMember)
}

func (c SaddMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Scard Incomplete

func (b Builder) Scard() (c Scard) { _ = "STUB: not implemented"; return *new(Scard) }

func (c Scard) Key(key string) ScardKey { _ = "STUB: not implemented"; return *new(ScardKey) }

type ScardKey Incomplete

func (c ScardKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ScardKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Sdiff Incomplete

func (b Builder) Sdiff() (c Sdiff) { _ = "STUB: not implemented"; return *new(Sdiff) }

func (c Sdiff) Key(key ...string) SdiffKey { _ = "STUB: not implemented"; return *new(SdiffKey) }

type SdiffKey Incomplete

func (c SdiffKey) Key(key ...string) SdiffKey { _ = "STUB: not implemented"; return *new(SdiffKey) }

func (c SdiffKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sdiffstore Incomplete

func (b Builder) Sdiffstore() (c Sdiffstore) { _ = "STUB: not implemented"; return *new(Sdiffstore) }

func (c Sdiffstore) Destination(destination string) SdiffstoreDestination {
	_ = "STUB: not implemented"
	return *new(SdiffstoreDestination)
}

type SdiffstoreDestination Incomplete

func (c SdiffstoreDestination) Key(key ...string) SdiffstoreKey {
	_ = "STUB: not implemented"
	return *new(SdiffstoreKey)
}

type SdiffstoreKey Incomplete

func (c SdiffstoreKey) Key(key ...string) SdiffstoreKey {
	_ = "STUB: not implemented"
	return *new(SdiffstoreKey)
}

func (c SdiffstoreKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sinter Incomplete

func (b Builder) Sinter() (c Sinter) { _ = "STUB: not implemented"; return *new(Sinter) }

func (c Sinter) Key(key ...string) SinterKey { _ = "STUB: not implemented"; return *new(SinterKey) }

type SinterKey Incomplete

func (c SinterKey) Key(key ...string) SinterKey { _ = "STUB: not implemented"; return *new(SinterKey) }

func (c SinterKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sintercard Incomplete

func (b Builder) Sintercard() (c Sintercard) { _ = "STUB: not implemented"; return *new(Sintercard) }

func (c Sintercard) Numkeys(numkeys int64) SintercardNumkeys {
	_ = "STUB: not implemented"
	return *new(SintercardNumkeys)
}

type SintercardKey Incomplete

func (c SintercardKey) Key(key ...string) SintercardKey {
	_ = "STUB: not implemented"
	return *new(SintercardKey)
}

func (c SintercardKey) Limit(limit int64) SintercardLimit {
	_ = "STUB: not implemented"
	return *new(SintercardLimit)
}

func (c SintercardKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SintercardLimit Incomplete

func (c SintercardLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SintercardNumkeys Incomplete

func (c SintercardNumkeys) Key(key ...string) SintercardKey {
	_ = "STUB: not implemented"
	return *new(SintercardKey)
}

type Sinterstore Incomplete

func (b Builder) Sinterstore() (c Sinterstore) { _ = "STUB: not implemented"; return *new(Sinterstore) }

func (c Sinterstore) Destination(destination string) SinterstoreDestination {
	_ = "STUB: not implemented"
	return *new(SinterstoreDestination)
}

type SinterstoreDestination Incomplete

func (c SinterstoreDestination) Key(key ...string) SinterstoreKey {
	_ = "STUB: not implemented"
	return *new(SinterstoreKey)
}

type SinterstoreKey Incomplete

func (c SinterstoreKey) Key(key ...string) SinterstoreKey {
	_ = "STUB: not implemented"
	return *new(SinterstoreKey)
}

func (c SinterstoreKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sismember Incomplete

func (b Builder) Sismember() (c Sismember) { _ = "STUB: not implemented"; return *new(Sismember) }

func (c Sismember) Key(key string) SismemberKey {
	_ = "STUB: not implemented"
	return *new(SismemberKey)
}

type SismemberKey Incomplete

func (c SismemberKey) Member(member string) SismemberMember {
	_ = "STUB: not implemented"
	return *new(SismemberMember)
}

type SismemberMember Incomplete

func (c SismemberMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SismemberMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Smembers Incomplete

func (b Builder) Smembers() (c Smembers) { _ = "STUB: not implemented"; return *new(Smembers) }

func (c Smembers) Key(key string) SmembersKey { _ = "STUB: not implemented"; return *new(SmembersKey) }

type SmembersKey Incomplete

func (c SmembersKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SmembersKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Smismember Incomplete

func (b Builder) Smismember() (c Smismember) { _ = "STUB: not implemented"; return *new(Smismember) }

func (c Smismember) Key(key string) SmismemberKey {
	_ = "STUB: not implemented"
	return *new(SmismemberKey)
}

type SmismemberKey Incomplete

func (c SmismemberKey) Member(member ...string) SmismemberMember {
	_ = "STUB: not implemented"
	return *new(SmismemberMember)
}

type SmismemberMember Incomplete

func (c SmismemberMember) Member(member ...string) SmismemberMember {
	_ = "STUB: not implemented"
	return *new(SmismemberMember)
}

func (c SmismemberMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SmismemberMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Smove Incomplete

func (b Builder) Smove() (c Smove) { _ = "STUB: not implemented"; return *new(Smove) }

func (c Smove) Source(source string) SmoveSource {
	_ = "STUB: not implemented"
	return *new(SmoveSource)
}

type SmoveDestination Incomplete

func (c SmoveDestination) Member(member string) SmoveMember {
	_ = "STUB: not implemented"
	return *new(SmoveMember)
}

type SmoveMember Incomplete

func (c SmoveMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SmoveSource Incomplete

func (c SmoveSource) Destination(destination string) SmoveDestination {
	_ = "STUB: not implemented"
	return *new(SmoveDestination)
}

type Spop Incomplete

func (b Builder) Spop() (c Spop) { _ = "STUB: not implemented"; return *new(Spop) }

func (c Spop) Key(key string) SpopKey { _ = "STUB: not implemented"; return *new(SpopKey) }

type SpopCount Incomplete

func (c SpopCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SpopKey Incomplete

func (c SpopKey) Count(count int64) SpopCount { _ = "STUB: not implemented"; return *new(SpopCount) }

func (c SpopKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Srandmember Incomplete

func (b Builder) Srandmember() (c Srandmember) { _ = "STUB: not implemented"; return *new(Srandmember) }

func (c Srandmember) Key(key string) SrandmemberKey {
	_ = "STUB: not implemented"
	return *new(SrandmemberKey)
}

type SrandmemberCount Incomplete

func (c SrandmemberCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SrandmemberKey Incomplete

func (c SrandmemberKey) Count(count int64) SrandmemberCount {
	_ = "STUB: not implemented"
	return *new(SrandmemberCount)
}

func (c SrandmemberKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Srem Incomplete

func (b Builder) Srem() (c Srem) { _ = "STUB: not implemented"; return *new(Srem) }

func (c Srem) Key(key string) SremKey { _ = "STUB: not implemented"; return *new(SremKey) }

type SremKey Incomplete

func (c SremKey) Member(member ...string) SremMember {
	_ = "STUB: not implemented"
	return *new(SremMember)
}

type SremMember Incomplete

func (c SremMember) Member(member ...string) SremMember {
	_ = "STUB: not implemented"
	return *new(SremMember)
}

func (c SremMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sscan Incomplete

func (b Builder) Sscan() (c Sscan) { _ = "STUB: not implemented"; return *new(Sscan) }

func (c Sscan) Key(key string) SscanKey { _ = "STUB: not implemented"; return *new(SscanKey) }

type SscanCount Incomplete

func (c SscanCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SscanCursor Incomplete

func (c SscanCursor) Match(pattern string) SscanMatch {
	_ = "STUB: not implemented"
	return *new(SscanMatch)
}

func (c SscanCursor) Count(count int64) SscanCount {
	_ = "STUB: not implemented"
	return *new(SscanCount)
}

func (c SscanCursor) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SscanKey Incomplete

func (c SscanKey) Cursor(cursor uint64) SscanCursor {
	_ = "STUB: not implemented"
	return *new(SscanCursor)
}

type SscanMatch Incomplete

func (c SscanMatch) Count(count int64) SscanCount {
	_ = "STUB: not implemented"
	return *new(SscanCount)
}

func (c SscanMatch) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sunion Incomplete

func (b Builder) Sunion() (c Sunion) { _ = "STUB: not implemented"; return *new(Sunion) }

func (c Sunion) Key(key ...string) SunionKey { _ = "STUB: not implemented"; return *new(SunionKey) }

type SunionKey Incomplete

func (c SunionKey) Key(key ...string) SunionKey { _ = "STUB: not implemented"; return *new(SunionKey) }

func (c SunionKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sunionstore Incomplete

func (b Builder) Sunionstore() (c Sunionstore) { _ = "STUB: not implemented"; return *new(Sunionstore) }

func (c Sunionstore) Destination(destination string) SunionstoreDestination {
	_ = "STUB: not implemented"
	return *new(SunionstoreDestination)
}

type SunionstoreDestination Incomplete

func (c SunionstoreDestination) Key(key ...string) SunionstoreKey {
	_ = "STUB: not implemented"
	return *new(SunionstoreKey)
}

type SunionstoreKey Incomplete

func (c SunionstoreKey) Key(key ...string) SunionstoreKey {
	_ = "STUB: not implemented"
	return *new(SunionstoreKey)
}

func (c SunionstoreKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
