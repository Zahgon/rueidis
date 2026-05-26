// Code generated DO NOT EDIT

package cmds

type Copy Incomplete

func (b Builder) Copy() (c Copy) { _ = "STUB: not implemented"; return *new(Copy) }

func (c Copy) Source(source string) CopySource { _ = "STUB: not implemented"; return *new(CopySource) }

type CopyDb Incomplete

func (c CopyDb) Replace() CopyReplace { _ = "STUB: not implemented"; return *new(CopyReplace) }

func (c CopyDb) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CopyDestination Incomplete

func (c CopyDestination) Db(destinationDb int64) CopyDb {
	_ = "STUB: not implemented"
	return *new(CopyDb)
}

func (c CopyDestination) Replace() CopyReplace { _ = "STUB: not implemented"; return *new(CopyReplace) }

func (c CopyDestination) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CopyReplace Incomplete

func (c CopyReplace) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CopySource Incomplete

func (c CopySource) Destination(destination string) CopyDestination {
	_ = "STUB: not implemented"
	return *new(CopyDestination)
}

type Del Incomplete

func (b Builder) Del() (c Del) { _ = "STUB: not implemented"; return *new(Del) }

func (c Del) Key(key ...string) DelKey { _ = "STUB: not implemented"; return *new(DelKey) }

type DelKey Incomplete

func (c DelKey) Key(key ...string) DelKey { _ = "STUB: not implemented"; return *new(DelKey) }

func (c DelKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Dump Incomplete

func (b Builder) Dump() (c Dump) { _ = "STUB: not implemented"; return *new(Dump) }

func (c Dump) Key(key string) DumpKey { _ = "STUB: not implemented"; return *new(DumpKey) }

type DumpKey Incomplete

func (c DumpKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Exists Incomplete

func (b Builder) Exists() (c Exists) { _ = "STUB: not implemented"; return *new(Exists) }

func (c Exists) Key(key ...string) ExistsKey { _ = "STUB: not implemented"; return *new(ExistsKey) }

type ExistsKey Incomplete

func (c ExistsKey) Key(key ...string) ExistsKey { _ = "STUB: not implemented"; return *new(ExistsKey) }

func (c ExistsKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Expire Incomplete

func (b Builder) Expire() (c Expire) { _ = "STUB: not implemented"; return *new(Expire) }

func (c Expire) Key(key string) ExpireKey { _ = "STUB: not implemented"; return *new(ExpireKey) }

type ExpireConditionGt Incomplete

func (c ExpireConditionGt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireConditionLt Incomplete

func (c ExpireConditionLt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireConditionNx Incomplete

func (c ExpireConditionNx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireConditionXx Incomplete

func (c ExpireConditionXx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireKey Incomplete

func (c ExpireKey) Seconds(seconds int64) ExpireSeconds {
	_ = "STUB: not implemented"
	return *new(ExpireSeconds)
}

type ExpireSeconds Incomplete

func (c ExpireSeconds) Nx() ExpireConditionNx {
	_ = "STUB: not implemented"
	return *new(ExpireConditionNx)
}

func (c ExpireSeconds) Xx() ExpireConditionXx {
	_ = "STUB: not implemented"
	return *new(ExpireConditionXx)
}

func (c ExpireSeconds) Gt() ExpireConditionGt {
	_ = "STUB: not implemented"
	return *new(ExpireConditionGt)
}

func (c ExpireSeconds) Lt() ExpireConditionLt {
	_ = "STUB: not implemented"
	return *new(ExpireConditionLt)
}

func (c ExpireSeconds) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Expireat Incomplete

func (b Builder) Expireat() (c Expireat) { _ = "STUB: not implemented"; return *new(Expireat) }

func (c Expireat) Key(key string) ExpireatKey { _ = "STUB: not implemented"; return *new(ExpireatKey) }

type ExpireatConditionGt Incomplete

func (c ExpireatConditionGt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireatConditionLt Incomplete

func (c ExpireatConditionLt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireatConditionNx Incomplete

func (c ExpireatConditionNx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireatConditionXx Incomplete

func (c ExpireatConditionXx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ExpireatKey Incomplete

func (c ExpireatKey) Timestamp(timestamp int64) ExpireatTimestamp {
	_ = "STUB: not implemented"
	return *new(ExpireatTimestamp)
}

type ExpireatTimestamp Incomplete

func (c ExpireatTimestamp) Nx() ExpireatConditionNx {
	_ = "STUB: not implemented"
	return *new(ExpireatConditionNx)
}

func (c ExpireatTimestamp) Xx() ExpireatConditionXx {
	_ = "STUB: not implemented"
	return *new(ExpireatConditionXx)
}

func (c ExpireatTimestamp) Gt() ExpireatConditionGt {
	_ = "STUB: not implemented"
	return *new(ExpireatConditionGt)
}

func (c ExpireatTimestamp) Lt() ExpireatConditionLt {
	_ = "STUB: not implemented"
	return *new(ExpireatConditionLt)
}

func (c ExpireatTimestamp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Expiretime Incomplete

func (b Builder) Expiretime() (c Expiretime) { _ = "STUB: not implemented"; return *new(Expiretime) }

func (c Expiretime) Key(key string) ExpiretimeKey {
	_ = "STUB: not implemented"
	return *new(ExpiretimeKey)
}

type ExpiretimeKey Incomplete

func (c ExpiretimeKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ExpiretimeKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Keys Incomplete

func (b Builder) Keys() (c Keys) { _ = "STUB: not implemented"; return *new(Keys) }

func (c Keys) Pattern(pattern string) KeysPattern {
	_ = "STUB: not implemented"
	return *new(KeysPattern)
}

type KeysPattern Incomplete

func (c KeysPattern) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Migrate Incomplete

func (b Builder) Migrate() (c Migrate) { _ = "STUB: not implemented"; return *new(Migrate) }

func (c Migrate) Host(host string) MigrateHost { _ = "STUB: not implemented"; return *new(MigrateHost) }

type MigrateAuthAuth Incomplete

func (c MigrateAuthAuth) Auth2(username string, password string) MigrateAuthAuth2 {
	_ = "STUB: not implemented"
	return *new(MigrateAuthAuth2)
}

func (c MigrateAuthAuth) Keys(key ...string) MigrateKeys {
	_ = "STUB: not implemented"
	return *new(MigrateKeys)
}

func (c MigrateAuthAuth) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MigrateAuthAuth2 Incomplete

func (c MigrateAuthAuth2) Keys(key ...string) MigrateKeys {
	_ = "STUB: not implemented"
	return *new(MigrateKeys)
}

func (c MigrateAuthAuth2) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MigrateCopy Incomplete

func (c MigrateCopy) Replace() MigrateReplace {
	_ = "STUB: not implemented"
	return *new(MigrateReplace)
}

func (c MigrateCopy) Auth(password string) MigrateAuthAuth {
	_ = "STUB: not implemented"
	return *new(MigrateAuthAuth)
}

func (c MigrateCopy) Auth2(username string, password string) MigrateAuthAuth2 {
	_ = "STUB: not implemented"
	return *new(MigrateAuthAuth2)
}

func (c MigrateCopy) Keys(key ...string) MigrateKeys {
	_ = "STUB: not implemented"
	return *new(MigrateKeys)
}

func (c MigrateCopy) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MigrateDestinationDb Incomplete

func (c MigrateDestinationDb) Timeout(timeout int64) MigrateTimeout {
	_ = "STUB: not implemented"
	return *new(MigrateTimeout)
}

type MigrateHost Incomplete

func (c MigrateHost) Port(port int64) MigratePort {
	_ = "STUB: not implemented"
	return *new(MigratePort)
}

type MigrateKey Incomplete

func (c MigrateKey) DestinationDb(destinationDb int64) MigrateDestinationDb {
	_ = "STUB: not implemented"
	return *new(MigrateDestinationDb)
}

type MigrateKeys Incomplete

func (c MigrateKeys) Keys(key ...string) MigrateKeys {
	_ = "STUB: not implemented"
	return *new(MigrateKeys)
}

func (c MigrateKeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MigratePort Incomplete

func (c MigratePort) Key(key string) MigrateKey { _ = "STUB: not implemented"; return *new(MigrateKey) }

type MigrateReplace Incomplete

func (c MigrateReplace) Auth(password string) MigrateAuthAuth {
	_ = "STUB: not implemented"
	return *new(MigrateAuthAuth)
}

func (c MigrateReplace) Auth2(username string, password string) MigrateAuthAuth2 {
	_ = "STUB: not implemented"
	return *new(MigrateAuthAuth2)
}

func (c MigrateReplace) Keys(key ...string) MigrateKeys {
	_ = "STUB: not implemented"
	return *new(MigrateKeys)
}

func (c MigrateReplace) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MigrateTimeout Incomplete

func (c MigrateTimeout) Copy() MigrateCopy { _ = "STUB: not implemented"; return *new(MigrateCopy) }

func (c MigrateTimeout) Replace() MigrateReplace {
	_ = "STUB: not implemented"
	return *new(MigrateReplace)
}

func (c MigrateTimeout) Auth(password string) MigrateAuthAuth {
	_ = "STUB: not implemented"
	return *new(MigrateAuthAuth)
}

func (c MigrateTimeout) Auth2(username string, password string) MigrateAuthAuth2 {
	_ = "STUB: not implemented"
	return *new(MigrateAuthAuth2)
}

func (c MigrateTimeout) Keys(key ...string) MigrateKeys {
	_ = "STUB: not implemented"
	return *new(MigrateKeys)
}

func (c MigrateTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Move Incomplete

func (b Builder) Move() (c Move) { _ = "STUB: not implemented"; return *new(Move) }

func (c Move) Key(key string) MoveKey { _ = "STUB: not implemented"; return *new(MoveKey) }

type MoveDb Incomplete

func (c MoveDb) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MoveKey Incomplete

func (c MoveKey) Db(db int64) MoveDb { _ = "STUB: not implemented"; return *new(MoveDb) }

type ObjectEncoding Incomplete

func (b Builder) ObjectEncoding() (c ObjectEncoding) {
	_ = "STUB: not implemented"
	return *new(ObjectEncoding)
}

func (c ObjectEncoding) Key(key string) ObjectEncodingKey {
	_ = "STUB: not implemented"
	return *new(ObjectEncodingKey)
}

type ObjectEncodingKey Incomplete

func (c ObjectEncodingKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ObjectFreq Incomplete

func (b Builder) ObjectFreq() (c ObjectFreq) { _ = "STUB: not implemented"; return *new(ObjectFreq) }

func (c ObjectFreq) Key(key string) ObjectFreqKey {
	_ = "STUB: not implemented"
	return *new(ObjectFreqKey)
}

type ObjectFreqKey Incomplete

func (c ObjectFreqKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ObjectHelp Incomplete

func (b Builder) ObjectHelp() (c ObjectHelp) { _ = "STUB: not implemented"; return *new(ObjectHelp) }

func (c ObjectHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ObjectIdletime Incomplete

func (b Builder) ObjectIdletime() (c ObjectIdletime) {
	_ = "STUB: not implemented"
	return *new(ObjectIdletime)
}

func (c ObjectIdletime) Key(key string) ObjectIdletimeKey {
	_ = "STUB: not implemented"
	return *new(ObjectIdletimeKey)
}

type ObjectIdletimeKey Incomplete

func (c ObjectIdletimeKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ObjectRefcount Incomplete

func (b Builder) ObjectRefcount() (c ObjectRefcount) {
	_ = "STUB: not implemented"
	return *new(ObjectRefcount)
}

func (c ObjectRefcount) Key(key string) ObjectRefcountKey {
	_ = "STUB: not implemented"
	return *new(ObjectRefcountKey)
}

type ObjectRefcountKey Incomplete

func (c ObjectRefcountKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Persist Incomplete

func (b Builder) Persist() (c Persist) { _ = "STUB: not implemented"; return *new(Persist) }

func (c Persist) Key(key string) PersistKey { _ = "STUB: not implemented"; return *new(PersistKey) }

type PersistKey Incomplete

func (c PersistKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Pexpire Incomplete

func (b Builder) Pexpire() (c Pexpire) { _ = "STUB: not implemented"; return *new(Pexpire) }

func (c Pexpire) Key(key string) PexpireKey { _ = "STUB: not implemented"; return *new(PexpireKey) }

type PexpireConditionGt Incomplete

func (c PexpireConditionGt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireConditionLt Incomplete

func (c PexpireConditionLt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireConditionNx Incomplete

func (c PexpireConditionNx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireConditionXx Incomplete

func (c PexpireConditionXx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireKey Incomplete

func (c PexpireKey) Milliseconds(milliseconds int64) PexpireMilliseconds {
	_ = "STUB: not implemented"
	return *new(PexpireMilliseconds)
}

type PexpireMilliseconds Incomplete

func (c PexpireMilliseconds) Nx() PexpireConditionNx {
	_ = "STUB: not implemented"
	return *new(PexpireConditionNx)
}

func (c PexpireMilliseconds) Xx() PexpireConditionXx {
	_ = "STUB: not implemented"
	return *new(PexpireConditionXx)
}

func (c PexpireMilliseconds) Gt() PexpireConditionGt {
	_ = "STUB: not implemented"
	return *new(PexpireConditionGt)
}

func (c PexpireMilliseconds) Lt() PexpireConditionLt {
	_ = "STUB: not implemented"
	return *new(PexpireConditionLt)
}

func (c PexpireMilliseconds) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Pexpireat Incomplete

func (b Builder) Pexpireat() (c Pexpireat) { _ = "STUB: not implemented"; return *new(Pexpireat) }

func (c Pexpireat) Key(key string) PexpireatKey {
	_ = "STUB: not implemented"
	return *new(PexpireatKey)
}

type PexpireatConditionGt Incomplete

func (c PexpireatConditionGt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireatConditionLt Incomplete

func (c PexpireatConditionLt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireatConditionNx Incomplete

func (c PexpireatConditionNx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireatConditionXx Incomplete

func (c PexpireatConditionXx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PexpireatKey Incomplete

func (c PexpireatKey) MillisecondsTimestamp(millisecondsTimestamp int64) PexpireatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(PexpireatMillisecondsTimestamp)
}

type PexpireatMillisecondsTimestamp Incomplete

func (c PexpireatMillisecondsTimestamp) Nx() PexpireatConditionNx {
	_ = "STUB: not implemented"
	return *new(PexpireatConditionNx)
}

func (c PexpireatMillisecondsTimestamp) Xx() PexpireatConditionXx {
	_ = "STUB: not implemented"
	return *new(PexpireatConditionXx)
}

func (c PexpireatMillisecondsTimestamp) Gt() PexpireatConditionGt {
	_ = "STUB: not implemented"
	return *new(PexpireatConditionGt)
}

func (c PexpireatMillisecondsTimestamp) Lt() PexpireatConditionLt {
	_ = "STUB: not implemented"
	return *new(PexpireatConditionLt)
}

func (c PexpireatMillisecondsTimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type Pexpiretime Incomplete

func (b Builder) Pexpiretime() (c Pexpiretime) { _ = "STUB: not implemented"; return *new(Pexpiretime) }

func (c Pexpiretime) Key(key string) PexpiretimeKey {
	_ = "STUB: not implemented"
	return *new(PexpiretimeKey)
}

type PexpiretimeKey Incomplete

func (c PexpiretimeKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c PexpiretimeKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Pttl Incomplete

func (b Builder) Pttl() (c Pttl) { _ = "STUB: not implemented"; return *new(Pttl) }

func (c Pttl) Key(key string) PttlKey { _ = "STUB: not implemented"; return *new(PttlKey) }

type PttlKey Incomplete

func (c PttlKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c PttlKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Randomkey Incomplete

func (b Builder) Randomkey() (c Randomkey) { _ = "STUB: not implemented"; return *new(Randomkey) }

func (c Randomkey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Rename Incomplete

func (b Builder) Rename() (c Rename) { _ = "STUB: not implemented"; return *new(Rename) }

func (c Rename) Key(key string) RenameKey { _ = "STUB: not implemented"; return *new(RenameKey) }

type RenameKey Incomplete

func (c RenameKey) Newkey(newkey string) RenameNewkey {
	_ = "STUB: not implemented"
	return *new(RenameNewkey)
}

type RenameNewkey Incomplete

func (c RenameNewkey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Renamenx Incomplete

func (b Builder) Renamenx() (c Renamenx) { _ = "STUB: not implemented"; return *new(Renamenx) }

func (c Renamenx) Key(key string) RenamenxKey { _ = "STUB: not implemented"; return *new(RenamenxKey) }

type RenamenxKey Incomplete

func (c RenamenxKey) Newkey(newkey string) RenamenxNewkey {
	_ = "STUB: not implemented"
	return *new(RenamenxNewkey)
}

type RenamenxNewkey Incomplete

func (c RenamenxNewkey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Restore Incomplete

func (b Builder) Restore() (c Restore) { _ = "STUB: not implemented"; return *new(Restore) }

func (c Restore) Key(key string) RestoreKey { _ = "STUB: not implemented"; return *new(RestoreKey) }

type RestoreAbsttl Incomplete

func (c RestoreAbsttl) Idletime(seconds int64) RestoreIdletime {
	_ = "STUB: not implemented"
	return *new(RestoreIdletime)
}

func (c RestoreAbsttl) Freq(frequency int64) RestoreFreq {
	_ = "STUB: not implemented"
	return *new(RestoreFreq)
}

func (c RestoreAbsttl) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RestoreFreq Incomplete

func (c RestoreFreq) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RestoreIdletime Incomplete

func (c RestoreIdletime) Freq(frequency int64) RestoreFreq {
	_ = "STUB: not implemented"
	return *new(RestoreFreq)
}

func (c RestoreIdletime) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RestoreKey Incomplete

func (c RestoreKey) Ttl(ttl int64) RestoreTtl { _ = "STUB: not implemented"; return *new(RestoreTtl) }

type RestoreReplace Incomplete

func (c RestoreReplace) Absttl() RestoreAbsttl {
	_ = "STUB: not implemented"
	return *new(RestoreAbsttl)
}

func (c RestoreReplace) Idletime(seconds int64) RestoreIdletime {
	_ = "STUB: not implemented"
	return *new(RestoreIdletime)
}

func (c RestoreReplace) Freq(frequency int64) RestoreFreq {
	_ = "STUB: not implemented"
	return *new(RestoreFreq)
}

func (c RestoreReplace) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RestoreSerializedValue Incomplete

func (c RestoreSerializedValue) Replace() RestoreReplace {
	_ = "STUB: not implemented"
	return *new(RestoreReplace)
}

func (c RestoreSerializedValue) Absttl() RestoreAbsttl {
	_ = "STUB: not implemented"
	return *new(RestoreAbsttl)
}

func (c RestoreSerializedValue) Idletime(seconds int64) RestoreIdletime {
	_ = "STUB: not implemented"
	return *new(RestoreIdletime)
}

func (c RestoreSerializedValue) Freq(frequency int64) RestoreFreq {
	_ = "STUB: not implemented"
	return *new(RestoreFreq)
}

func (c RestoreSerializedValue) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type RestoreTtl Incomplete

func (c RestoreTtl) SerializedValue(serializedValue string) RestoreSerializedValue {
	_ = "STUB: not implemented"
	return *new(RestoreSerializedValue)
}

type Scan Incomplete

func (b Builder) Scan() (c Scan) { _ = "STUB: not implemented"; return *new(Scan) }

func (c Scan) Cursor(cursor uint64) ScanCursor { _ = "STUB: not implemented"; return *new(ScanCursor) }

type ScanCount Incomplete

func (c ScanCount) Type(typ string) ScanType { _ = "STUB: not implemented"; return *new(ScanType) }

func (c ScanCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScanCursor Incomplete

func (c ScanCursor) Match(pattern string) ScanMatch {
	_ = "STUB: not implemented"
	return *new(ScanMatch)
}

func (c ScanCursor) Count(count int64) ScanCount { _ = "STUB: not implemented"; return *new(ScanCount) }

func (c ScanCursor) Type(typ string) ScanType { _ = "STUB: not implemented"; return *new(ScanType) }

func (c ScanCursor) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScanMatch Incomplete

func (c ScanMatch) Count(count int64) ScanCount { _ = "STUB: not implemented"; return *new(ScanCount) }

func (c ScanMatch) Type(typ string) ScanType { _ = "STUB: not implemented"; return *new(ScanType) }

func (c ScanMatch) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScanType Incomplete

func (c ScanType) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sort Incomplete

func (b Builder) Sort() (c Sort) { _ = "STUB: not implemented"; return *new(Sort) }

func (c Sort) Key(key string) SortKey { _ = "STUB: not implemented"; return *new(SortKey) }

type SortBy Incomplete

func (c SortBy) Limit(offset int64, count int64) SortLimit {
	_ = "STUB: not implemented"
	return *new(SortLimit)
}

func (c SortBy) Get() SortGet { _ = "STUB: not implemented"; return *new(SortGet) }

func (c SortBy) Asc() SortOrderAsc { _ = "STUB: not implemented"; return *new(SortOrderAsc) }

func (c SortBy) Desc() SortOrderDesc { _ = "STUB: not implemented"; return *new(SortOrderDesc) }

func (c SortBy) Alpha() SortSortingAlpha { _ = "STUB: not implemented"; return *new(SortSortingAlpha) }

func (c SortBy) Store(destination string) SortStore {
	_ = "STUB: not implemented"
	return *new(SortStore)
}

func (c SortBy) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SortGet Incomplete

func (c SortGet) Get(pattern string) SortGet { _ = "STUB: not implemented"; return *new(SortGet) }

func (c SortGet) Asc() SortOrderAsc { _ = "STUB: not implemented"; return *new(SortOrderAsc) }

func (c SortGet) Desc() SortOrderDesc { _ = "STUB: not implemented"; return *new(SortOrderDesc) }

func (c SortGet) Alpha() SortSortingAlpha { _ = "STUB: not implemented"; return *new(SortSortingAlpha) }

func (c SortGet) Store(destination string) SortStore {
	_ = "STUB: not implemented"
	return *new(SortStore)
}

func (c SortGet) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SortKey Incomplete

func (c SortKey) By(pattern string) SortBy { _ = "STUB: not implemented"; return *new(SortBy) }

func (c SortKey) Limit(offset int64, count int64) SortLimit {
	_ = "STUB: not implemented"
	return *new(SortLimit)
}

func (c SortKey) Get() SortGet { _ = "STUB: not implemented"; return *new(SortGet) }

func (c SortKey) Asc() SortOrderAsc { _ = "STUB: not implemented"; return *new(SortOrderAsc) }

func (c SortKey) Desc() SortOrderDesc { _ = "STUB: not implemented"; return *new(SortOrderDesc) }

func (c SortKey) Alpha() SortSortingAlpha { _ = "STUB: not implemented"; return *new(SortSortingAlpha) }

func (c SortKey) Store(destination string) SortStore {
	_ = "STUB: not implemented"
	return *new(SortStore)
}

func (c SortKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SortLimit Incomplete

func (c SortLimit) Get() SortGet { _ = "STUB: not implemented"; return *new(SortGet) }

func (c SortLimit) Asc() SortOrderAsc { _ = "STUB: not implemented"; return *new(SortOrderAsc) }

func (c SortLimit) Desc() SortOrderDesc { _ = "STUB: not implemented"; return *new(SortOrderDesc) }

func (c SortLimit) Alpha() SortSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortSortingAlpha)
}

func (c SortLimit) Store(destination string) SortStore {
	_ = "STUB: not implemented"
	return *new(SortStore)
}

func (c SortLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SortOrderAsc Incomplete

func (c SortOrderAsc) Alpha() SortSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortSortingAlpha)
}

func (c SortOrderAsc) Store(destination string) SortStore {
	_ = "STUB: not implemented"
	return *new(SortStore)
}

func (c SortOrderAsc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SortOrderDesc Incomplete

func (c SortOrderDesc) Alpha() SortSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortSortingAlpha)
}

func (c SortOrderDesc) Store(destination string) SortStore {
	_ = "STUB: not implemented"
	return *new(SortStore)
}

func (c SortOrderDesc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SortRo Incomplete

func (b Builder) SortRo() (c SortRo) { _ = "STUB: not implemented"; return *new(SortRo) }

func (c SortRo) Key(key string) SortRoKey { _ = "STUB: not implemented"; return *new(SortRoKey) }

type SortRoBy Incomplete

func (c SortRoBy) Limit(offset int64, count int64) SortRoLimit {
	_ = "STUB: not implemented"
	return *new(SortRoLimit)
}

func (c SortRoBy) Get() SortRoGet { _ = "STUB: not implemented"; return *new(SortRoGet) }

func (c SortRoBy) Asc() SortRoOrderAsc { _ = "STUB: not implemented"; return *new(SortRoOrderAsc) }

func (c SortRoBy) Desc() SortRoOrderDesc { _ = "STUB: not implemented"; return *new(SortRoOrderDesc) }

func (c SortRoBy) Alpha() SortRoSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortRoSortingAlpha)
}

func (c SortRoBy) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SortRoBy) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type SortRoGet Incomplete

func (c SortRoGet) Get(pattern string) SortRoGet { _ = "STUB: not implemented"; return *new(SortRoGet) }

func (c SortRoGet) Asc() SortRoOrderAsc { _ = "STUB: not implemented"; return *new(SortRoOrderAsc) }

func (c SortRoGet) Desc() SortRoOrderDesc { _ = "STUB: not implemented"; return *new(SortRoOrderDesc) }

func (c SortRoGet) Alpha() SortRoSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortRoSortingAlpha)
}

func (c SortRoGet) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SortRoGet) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type SortRoKey Incomplete

func (c SortRoKey) By(pattern string) SortRoBy { _ = "STUB: not implemented"; return *new(SortRoBy) }

func (c SortRoKey) Limit(offset int64, count int64) SortRoLimit {
	_ = "STUB: not implemented"
	return *new(SortRoLimit)
}

func (c SortRoKey) Get() SortRoGet { _ = "STUB: not implemented"; return *new(SortRoGet) }

func (c SortRoKey) Asc() SortRoOrderAsc { _ = "STUB: not implemented"; return *new(SortRoOrderAsc) }

func (c SortRoKey) Desc() SortRoOrderDesc { _ = "STUB: not implemented"; return *new(SortRoOrderDesc) }

func (c SortRoKey) Alpha() SortRoSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortRoSortingAlpha)
}

func (c SortRoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SortRoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type SortRoLimit Incomplete

func (c SortRoLimit) Get() SortRoGet { _ = "STUB: not implemented"; return *new(SortRoGet) }

func (c SortRoLimit) Asc() SortRoOrderAsc { _ = "STUB: not implemented"; return *new(SortRoOrderAsc) }

func (c SortRoLimit) Desc() SortRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(SortRoOrderDesc)
}

func (c SortRoLimit) Alpha() SortRoSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortRoSortingAlpha)
}

func (c SortRoLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SortRoLimit) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type SortRoOrderAsc Incomplete

func (c SortRoOrderAsc) Alpha() SortRoSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortRoSortingAlpha)
}

func (c SortRoOrderAsc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SortRoOrderAsc) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type SortRoOrderDesc Incomplete

func (c SortRoOrderDesc) Alpha() SortRoSortingAlpha {
	_ = "STUB: not implemented"
	return *new(SortRoSortingAlpha)
}

func (c SortRoOrderDesc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SortRoOrderDesc) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type SortRoSortingAlpha Incomplete

func (c SortRoSortingAlpha) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c SortRoSortingAlpha) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type SortSortingAlpha Incomplete

func (c SortSortingAlpha) Store(destination string) SortStore {
	_ = "STUB: not implemented"
	return *new(SortStore)
}

func (c SortSortingAlpha) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SortStore Incomplete

func (c SortStore) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Touch Incomplete

func (b Builder) Touch() (c Touch) { _ = "STUB: not implemented"; return *new(Touch) }

func (c Touch) Key(key ...string) TouchKey { _ = "STUB: not implemented"; return *new(TouchKey) }

type TouchKey Incomplete

func (c TouchKey) Key(key ...string) TouchKey { _ = "STUB: not implemented"; return *new(TouchKey) }

func (c TouchKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Ttl Incomplete

func (b Builder) Ttl() (c Ttl) { _ = "STUB: not implemented"; return *new(Ttl) }

func (c Ttl) Key(key string) TtlKey { _ = "STUB: not implemented"; return *new(TtlKey) }

type TtlKey Incomplete

func (c TtlKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c TtlKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Type Incomplete

func (b Builder) Type() (c Type) { _ = "STUB: not implemented"; return *new(Type) }

func (c Type) Key(key string) TypeKey { _ = "STUB: not implemented"; return *new(TypeKey) }

type TypeKey Incomplete

func (c TypeKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c TypeKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Unlink Incomplete

func (b Builder) Unlink() (c Unlink) { _ = "STUB: not implemented"; return *new(Unlink) }

func (c Unlink) Key(key ...string) UnlinkKey { _ = "STUB: not implemented"; return *new(UnlinkKey) }

type UnlinkKey Incomplete

func (c UnlinkKey) Key(key ...string) UnlinkKey { _ = "STUB: not implemented"; return *new(UnlinkKey) }

func (c UnlinkKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Wait Incomplete

func (b Builder) Wait() (c Wait) { _ = "STUB: not implemented"; return *new(Wait) }

func (c Wait) Numreplicas(numreplicas int64) WaitNumreplicas {
	_ = "STUB: not implemented"
	return *new(WaitNumreplicas)
}

type WaitNumreplicas Incomplete

func (c WaitNumreplicas) Timeout(timeout int64) WaitTimeout {
	_ = "STUB: not implemented"
	return *new(WaitTimeout)
}

type WaitTimeout Incomplete

func (c WaitTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Waitaof Incomplete

func (b Builder) Waitaof() (c Waitaof) { _ = "STUB: not implemented"; return *new(Waitaof) }

func (c Waitaof) Numlocal(numlocal int64) WaitaofNumlocal {
	_ = "STUB: not implemented"
	return *new(WaitaofNumlocal)
}

type WaitaofNumlocal Incomplete

func (c WaitaofNumlocal) Numreplicas(numreplicas int64) WaitaofNumreplicas {
	_ = "STUB: not implemented"
	return *new(WaitaofNumreplicas)
}

type WaitaofNumreplicas Incomplete

func (c WaitaofNumreplicas) Timeout(timeout int64) WaitaofTimeout {
	_ = "STUB: not implemented"
	return *new(WaitaofTimeout)
}

type WaitaofTimeout Incomplete

func (c WaitaofTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
