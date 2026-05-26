// Code generated DO NOT EDIT

package cmds

type Hdel Incomplete

func (b Builder) Hdel() (c Hdel) { _ = "STUB: not implemented"; return *new(Hdel) }

func (c Hdel) Key(key string) HdelKey { _ = "STUB: not implemented"; return *new(HdelKey) }

type HdelField Incomplete

func (c HdelField) Field(field ...string) HdelField {
	_ = "STUB: not implemented"
	return *new(HdelField)
}

func (c HdelField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HdelKey Incomplete

func (c HdelKey) Field(field ...string) HdelField {
	_ = "STUB: not implemented"
	return *new(HdelField)
}

type Hexists Incomplete

func (b Builder) Hexists() (c Hexists) { _ = "STUB: not implemented"; return *new(Hexists) }

func (c Hexists) Key(key string) HexistsKey { _ = "STUB: not implemented"; return *new(HexistsKey) }

type HexistsField Incomplete

func (c HexistsField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HexistsField) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type HexistsKey Incomplete

func (c HexistsKey) Field(field string) HexistsField {
	_ = "STUB: not implemented"
	return *new(HexistsField)
}

type Hexpire Incomplete

func (b Builder) Hexpire() (c Hexpire) { _ = "STUB: not implemented"; return *new(Hexpire) }

func (c Hexpire) Key(key string) HexpireKey { _ = "STUB: not implemented"; return *new(HexpireKey) }

type HexpireConditionGt Incomplete

func (c HexpireConditionGt) Fields() HexpireFields {
	_ = "STUB: not implemented"
	return *new(HexpireFields)
}

type HexpireConditionLt Incomplete

func (c HexpireConditionLt) Fields() HexpireFields {
	_ = "STUB: not implemented"
	return *new(HexpireFields)
}

type HexpireConditionNx Incomplete

func (c HexpireConditionNx) Fields() HexpireFields {
	_ = "STUB: not implemented"
	return *new(HexpireFields)
}

type HexpireConditionXx Incomplete

func (c HexpireConditionXx) Fields() HexpireFields {
	_ = "STUB: not implemented"
	return *new(HexpireFields)
}

type HexpireField Incomplete

func (c HexpireField) Field(field ...string) HexpireField {
	_ = "STUB: not implemented"
	return *new(HexpireField)
}

func (c HexpireField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HexpireFields Incomplete

func (c HexpireFields) Numfields(numfields int64) HexpireNumfields {
	_ = "STUB: not implemented"
	return *new(HexpireNumfields)
}

type HexpireKey Incomplete

func (c HexpireKey) Seconds(seconds int64) HexpireSeconds {
	_ = "STUB: not implemented"
	return *new(HexpireSeconds)
}

type HexpireNumfields Incomplete

func (c HexpireNumfields) Field(field ...string) HexpireField {
	_ = "STUB: not implemented"
	return *new(HexpireField)
}

type HexpireSeconds Incomplete

func (c HexpireSeconds) Nx() HexpireConditionNx {
	_ = "STUB: not implemented"
	return *new(HexpireConditionNx)
}

func (c HexpireSeconds) Xx() HexpireConditionXx {
	_ = "STUB: not implemented"
	return *new(HexpireConditionXx)
}

func (c HexpireSeconds) Gt() HexpireConditionGt {
	_ = "STUB: not implemented"
	return *new(HexpireConditionGt)
}

func (c HexpireSeconds) Lt() HexpireConditionLt {
	_ = "STUB: not implemented"
	return *new(HexpireConditionLt)
}

func (c HexpireSeconds) Fields() HexpireFields {
	_ = "STUB: not implemented"
	return *new(HexpireFields)
}

type Hexpireat Incomplete

func (b Builder) Hexpireat() (c Hexpireat) { _ = "STUB: not implemented"; return *new(Hexpireat) }

func (c Hexpireat) Key(key string) HexpireatKey {
	_ = "STUB: not implemented"
	return *new(HexpireatKey)
}

type HexpireatConditionGt Incomplete

func (c HexpireatConditionGt) Fields() HexpireatFields {
	_ = "STUB: not implemented"
	return *new(HexpireatFields)
}

type HexpireatConditionLt Incomplete

func (c HexpireatConditionLt) Fields() HexpireatFields {
	_ = "STUB: not implemented"
	return *new(HexpireatFields)
}

type HexpireatConditionNx Incomplete

func (c HexpireatConditionNx) Fields() HexpireatFields {
	_ = "STUB: not implemented"
	return *new(HexpireatFields)
}

type HexpireatConditionXx Incomplete

func (c HexpireatConditionXx) Fields() HexpireatFields {
	_ = "STUB: not implemented"
	return *new(HexpireatFields)
}

type HexpireatField Incomplete

func (c HexpireatField) Field(field ...string) HexpireatField {
	_ = "STUB: not implemented"
	return *new(HexpireatField)
}

func (c HexpireatField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HexpireatFields Incomplete

func (c HexpireatFields) Numfields(numfields int64) HexpireatNumfields {
	_ = "STUB: not implemented"
	return *new(HexpireatNumfields)
}

type HexpireatKey Incomplete

func (c HexpireatKey) UnixTimeSeconds(unixTimeSeconds int64) HexpireatUnixTimeSeconds {
	_ = "STUB: not implemented"
	return *new(HexpireatUnixTimeSeconds)
}

type HexpireatNumfields Incomplete

func (c HexpireatNumfields) Field(field ...string) HexpireatField {
	_ = "STUB: not implemented"
	return *new(HexpireatField)
}

type HexpireatUnixTimeSeconds Incomplete

func (c HexpireatUnixTimeSeconds) Nx() HexpireatConditionNx {
	_ = "STUB: not implemented"
	return *new(HexpireatConditionNx)
}

func (c HexpireatUnixTimeSeconds) Xx() HexpireatConditionXx {
	_ = "STUB: not implemented"
	return *new(HexpireatConditionXx)
}

func (c HexpireatUnixTimeSeconds) Gt() HexpireatConditionGt {
	_ = "STUB: not implemented"
	return *new(HexpireatConditionGt)
}

func (c HexpireatUnixTimeSeconds) Lt() HexpireatConditionLt {
	_ = "STUB: not implemented"
	return *new(HexpireatConditionLt)
}

func (c HexpireatUnixTimeSeconds) Fields() HexpireatFields {
	_ = "STUB: not implemented"
	return *new(HexpireatFields)
}

type Hexpiretime Incomplete

func (b Builder) Hexpiretime() (c Hexpiretime) { _ = "STUB: not implemented"; return *new(Hexpiretime) }

func (c Hexpiretime) Key(key string) HexpiretimeKey {
	_ = "STUB: not implemented"
	return *new(HexpiretimeKey)
}

type HexpiretimeField Incomplete

func (c HexpiretimeField) Field(field ...string) HexpiretimeField {
	_ = "STUB: not implemented"
	return *new(HexpiretimeField)
}

func (c HexpiretimeField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HexpiretimeFields Incomplete

func (c HexpiretimeFields) Numfields(numfields int64) HexpiretimeNumfields {
	_ = "STUB: not implemented"
	return *new(HexpiretimeNumfields)
}

type HexpiretimeKey Incomplete

func (c HexpiretimeKey) Fields() HexpiretimeFields {
	_ = "STUB: not implemented"
	return *new(HexpiretimeFields)
}

type HexpiretimeNumfields Incomplete

func (c HexpiretimeNumfields) Field(field ...string) HexpiretimeField {
	_ = "STUB: not implemented"
	return *new(HexpiretimeField)
}

type Hget Incomplete

func (b Builder) Hget() (c Hget) { _ = "STUB: not implemented"; return *new(Hget) }

func (c Hget) Key(key string) HgetKey { _ = "STUB: not implemented"; return *new(HgetKey) }

type HgetField Incomplete

func (c HgetField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HgetField) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type HgetKey Incomplete

func (c HgetKey) Field(field string) HgetField { _ = "STUB: not implemented"; return *new(HgetField) }

type Hgetall Incomplete

func (b Builder) Hgetall() (c Hgetall) { _ = "STUB: not implemented"; return *new(Hgetall) }

func (c Hgetall) Key(key string) HgetallKey { _ = "STUB: not implemented"; return *new(HgetallKey) }

type HgetallKey Incomplete

func (c HgetallKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HgetallKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Hgetdel Incomplete

func (b Builder) Hgetdel() (c Hgetdel) { _ = "STUB: not implemented"; return *new(Hgetdel) }

func (c Hgetdel) Key(key string) HgetdelKey { _ = "STUB: not implemented"; return *new(HgetdelKey) }

type HgetdelField Incomplete

func (c HgetdelField) Field(field ...string) HgetdelField {
	_ = "STUB: not implemented"
	return *new(HgetdelField)
}

func (c HgetdelField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HgetdelFields Incomplete

func (c HgetdelFields) Numfields(numfields int64) HgetdelNumfields {
	_ = "STUB: not implemented"
	return *new(HgetdelNumfields)
}

type HgetdelKey Incomplete

func (c HgetdelKey) Fields() HgetdelFields { _ = "STUB: not implemented"; return *new(HgetdelFields) }

type HgetdelNumfields Incomplete

func (c HgetdelNumfields) Field(field ...string) HgetdelField {
	_ = "STUB: not implemented"
	return *new(HgetdelField)
}

type Hgetex Incomplete

func (b Builder) Hgetex() (c Hgetex) { _ = "STUB: not implemented"; return *new(Hgetex) }

func (c Hgetex) Key(key string) HgetexKey { _ = "STUB: not implemented"; return *new(HgetexKey) }

type HgetexExpirationEx Incomplete

func (c HgetexExpirationEx) Fields() HgetexFields {
	_ = "STUB: not implemented"
	return *new(HgetexFields)
}

type HgetexExpirationExat Incomplete

func (c HgetexExpirationExat) Fields() HgetexFields {
	_ = "STUB: not implemented"
	return *new(HgetexFields)
}

type HgetexExpirationPersist Incomplete

func (c HgetexExpirationPersist) Fields() HgetexFields {
	_ = "STUB: not implemented"
	return *new(HgetexFields)
}

type HgetexExpirationPx Incomplete

func (c HgetexExpirationPx) Fields() HgetexFields {
	_ = "STUB: not implemented"
	return *new(HgetexFields)
}

type HgetexExpirationPxat Incomplete

func (c HgetexExpirationPxat) Fields() HgetexFields {
	_ = "STUB: not implemented"
	return *new(HgetexFields)
}

type HgetexField Incomplete

func (c HgetexField) Field(field ...string) HgetexField {
	_ = "STUB: not implemented"
	return *new(HgetexField)
}

func (c HgetexField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HgetexFields Incomplete

func (c HgetexFields) Numfields(numfields int64) HgetexNumfields {
	_ = "STUB: not implemented"
	return *new(HgetexNumfields)
}

type HgetexKey Incomplete

func (c HgetexKey) Ex(ex int64) HgetexExpirationEx {
	_ = "STUB: not implemented"
	return *new(HgetexExpirationEx)
}

func (c HgetexKey) Px(px int64) HgetexExpirationPx {
	_ = "STUB: not implemented"
	return *new(HgetexExpirationPx)
}

func (c HgetexKey) Exat(exat int64) HgetexExpirationExat {
	_ = "STUB: not implemented"
	return *new(HgetexExpirationExat)
}

func (c HgetexKey) Pxat(pxat int64) HgetexExpirationPxat {
	_ = "STUB: not implemented"
	return *new(HgetexExpirationPxat)
}

func (c HgetexKey) Persist() HgetexExpirationPersist {
	_ = "STUB: not implemented"
	return *new(HgetexExpirationPersist)
}

func (c HgetexKey) Fields() HgetexFields { _ = "STUB: not implemented"; return *new(HgetexFields) }

type HgetexNumfields Incomplete

func (c HgetexNumfields) Field(field ...string) HgetexField {
	_ = "STUB: not implemented"
	return *new(HgetexField)
}

type Hincrby Incomplete

func (b Builder) Hincrby() (c Hincrby) { _ = "STUB: not implemented"; return *new(Hincrby) }

func (c Hincrby) Key(key string) HincrbyKey { _ = "STUB: not implemented"; return *new(HincrbyKey) }

type HincrbyField Incomplete

func (c HincrbyField) Increment(increment int64) HincrbyIncrement {
	_ = "STUB: not implemented"
	return *new(HincrbyIncrement)
}

type HincrbyIncrement Incomplete

func (c HincrbyIncrement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HincrbyKey Incomplete

func (c HincrbyKey) Field(field string) HincrbyField {
	_ = "STUB: not implemented"
	return *new(HincrbyField)
}

type Hincrbyfloat Incomplete

func (b Builder) Hincrbyfloat() (c Hincrbyfloat) {
	_ = "STUB: not implemented"
	return *new(Hincrbyfloat)
}

func (c Hincrbyfloat) Key(key string) HincrbyfloatKey {
	_ = "STUB: not implemented"
	return *new(HincrbyfloatKey)
}

type HincrbyfloatField Incomplete

func (c HincrbyfloatField) Increment(increment float64) HincrbyfloatIncrement {
	_ = "STUB: not implemented"
	return *new(HincrbyfloatIncrement)
}

type HincrbyfloatIncrement Incomplete

func (c HincrbyfloatIncrement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HincrbyfloatKey Incomplete

func (c HincrbyfloatKey) Field(field string) HincrbyfloatField {
	_ = "STUB: not implemented"
	return *new(HincrbyfloatField)
}

type Hkeys Incomplete

func (b Builder) Hkeys() (c Hkeys) { _ = "STUB: not implemented"; return *new(Hkeys) }

func (c Hkeys) Key(key string) HkeysKey { _ = "STUB: not implemented"; return *new(HkeysKey) }

type HkeysKey Incomplete

func (c HkeysKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HkeysKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Hlen Incomplete

func (b Builder) Hlen() (c Hlen) { _ = "STUB: not implemented"; return *new(Hlen) }

func (c Hlen) Key(key string) HlenKey { _ = "STUB: not implemented"; return *new(HlenKey) }

type HlenKey Incomplete

func (c HlenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HlenKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Hmget Incomplete

func (b Builder) Hmget() (c Hmget) { _ = "STUB: not implemented"; return *new(Hmget) }

func (c Hmget) Key(key string) HmgetKey { _ = "STUB: not implemented"; return *new(HmgetKey) }

type HmgetField Incomplete

func (c HmgetField) Field(field ...string) HmgetField {
	_ = "STUB: not implemented"
	return *new(HmgetField)
}

func (c HmgetField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HmgetField) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type HmgetKey Incomplete

func (c HmgetKey) Field(field ...string) HmgetField {
	_ = "STUB: not implemented"
	return *new(HmgetField)
}

type Hmset Incomplete

func (b Builder) Hmset() (c Hmset) { _ = "STUB: not implemented"; return *new(Hmset) }

func (c Hmset) Key(key string) HmsetKey { _ = "STUB: not implemented"; return *new(HmsetKey) }

type HmsetFieldValue Incomplete

func (c HmsetFieldValue) FieldValue(field string, value string) HmsetFieldValue {
	_ = "STUB: not implemented"
	return *new(HmsetFieldValue)
}

func (c HmsetFieldValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HmsetKey Incomplete

func (c HmsetKey) FieldValue() HmsetFieldValue {
	_ = "STUB: not implemented"
	return *new(HmsetFieldValue)
}

type Hpersist Incomplete

func (b Builder) Hpersist() (c Hpersist) { _ = "STUB: not implemented"; return *new(Hpersist) }

func (c Hpersist) Key(key string) HpersistKey { _ = "STUB: not implemented"; return *new(HpersistKey) }

type HpersistField Incomplete

func (c HpersistField) Field(field ...string) HpersistField {
	_ = "STUB: not implemented"
	return *new(HpersistField)
}

func (c HpersistField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HpersistFields Incomplete

func (c HpersistFields) Numfields(numfields int64) HpersistNumfields {
	_ = "STUB: not implemented"
	return *new(HpersistNumfields)
}

type HpersistKey Incomplete

func (c HpersistKey) Fields() HpersistFields {
	_ = "STUB: not implemented"
	return *new(HpersistFields)
}

type HpersistNumfields Incomplete

func (c HpersistNumfields) Field(field ...string) HpersistField {
	_ = "STUB: not implemented"
	return *new(HpersistField)
}

type Hpexpire Incomplete

func (b Builder) Hpexpire() (c Hpexpire) { _ = "STUB: not implemented"; return *new(Hpexpire) }

func (c Hpexpire) Key(key string) HpexpireKey { _ = "STUB: not implemented"; return *new(HpexpireKey) }

type HpexpireConditionGt Incomplete

func (c HpexpireConditionGt) Fields() HpexpireFields {
	_ = "STUB: not implemented"
	return *new(HpexpireFields)
}

type HpexpireConditionLt Incomplete

func (c HpexpireConditionLt) Fields() HpexpireFields {
	_ = "STUB: not implemented"
	return *new(HpexpireFields)
}

type HpexpireConditionNx Incomplete

func (c HpexpireConditionNx) Fields() HpexpireFields {
	_ = "STUB: not implemented"
	return *new(HpexpireFields)
}

type HpexpireConditionXx Incomplete

func (c HpexpireConditionXx) Fields() HpexpireFields {
	_ = "STUB: not implemented"
	return *new(HpexpireFields)
}

type HpexpireField Incomplete

func (c HpexpireField) Field(field ...string) HpexpireField {
	_ = "STUB: not implemented"
	return *new(HpexpireField)
}

func (c HpexpireField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HpexpireFields Incomplete

func (c HpexpireFields) Numfields(numfields int64) HpexpireNumfields {
	_ = "STUB: not implemented"
	return *new(HpexpireNumfields)
}

type HpexpireKey Incomplete

func (c HpexpireKey) Milliseconds(milliseconds int64) HpexpireMilliseconds {
	_ = "STUB: not implemented"
	return *new(HpexpireMilliseconds)
}

type HpexpireMilliseconds Incomplete

func (c HpexpireMilliseconds) Nx() HpexpireConditionNx {
	_ = "STUB: not implemented"
	return *new(HpexpireConditionNx)
}

func (c HpexpireMilliseconds) Xx() HpexpireConditionXx {
	_ = "STUB: not implemented"
	return *new(HpexpireConditionXx)
}

func (c HpexpireMilliseconds) Gt() HpexpireConditionGt {
	_ = "STUB: not implemented"
	return *new(HpexpireConditionGt)
}

func (c HpexpireMilliseconds) Lt() HpexpireConditionLt {
	_ = "STUB: not implemented"
	return *new(HpexpireConditionLt)
}

func (c HpexpireMilliseconds) Fields() HpexpireFields {
	_ = "STUB: not implemented"
	return *new(HpexpireFields)
}

type HpexpireNumfields Incomplete

func (c HpexpireNumfields) Field(field ...string) HpexpireField {
	_ = "STUB: not implemented"
	return *new(HpexpireField)
}

type Hpexpireat Incomplete

func (b Builder) Hpexpireat() (c Hpexpireat) { _ = "STUB: not implemented"; return *new(Hpexpireat) }

func (c Hpexpireat) Key(key string) HpexpireatKey {
	_ = "STUB: not implemented"
	return *new(HpexpireatKey)
}

type HpexpireatConditionGt Incomplete

func (c HpexpireatConditionGt) Fields() HpexpireatFields {
	_ = "STUB: not implemented"
	return *new(HpexpireatFields)
}

type HpexpireatConditionLt Incomplete

func (c HpexpireatConditionLt) Fields() HpexpireatFields {
	_ = "STUB: not implemented"
	return *new(HpexpireatFields)
}

type HpexpireatConditionNx Incomplete

func (c HpexpireatConditionNx) Fields() HpexpireatFields {
	_ = "STUB: not implemented"
	return *new(HpexpireatFields)
}

type HpexpireatConditionXx Incomplete

func (c HpexpireatConditionXx) Fields() HpexpireatFields {
	_ = "STUB: not implemented"
	return *new(HpexpireatFields)
}

type HpexpireatField Incomplete

func (c HpexpireatField) Field(field ...string) HpexpireatField {
	_ = "STUB: not implemented"
	return *new(HpexpireatField)
}

func (c HpexpireatField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HpexpireatFields Incomplete

func (c HpexpireatFields) Numfields(numfields int64) HpexpireatNumfields {
	_ = "STUB: not implemented"
	return *new(HpexpireatNumfields)
}

type HpexpireatKey Incomplete

func (c HpexpireatKey) UnixTimeMilliseconds(unixTimeMilliseconds int64) HpexpireatUnixTimeMilliseconds {
	_ = "STUB: not implemented"
	return *new(HpexpireatUnixTimeMilliseconds)
}

type HpexpireatNumfields Incomplete

func (c HpexpireatNumfields) Field(field ...string) HpexpireatField {
	_ = "STUB: not implemented"
	return *new(HpexpireatField)
}

type HpexpireatUnixTimeMilliseconds Incomplete

func (c HpexpireatUnixTimeMilliseconds) Nx() HpexpireatConditionNx {
	_ = "STUB: not implemented"
	return *new(HpexpireatConditionNx)
}

func (c HpexpireatUnixTimeMilliseconds) Xx() HpexpireatConditionXx {
	_ = "STUB: not implemented"
	return *new(HpexpireatConditionXx)
}

func (c HpexpireatUnixTimeMilliseconds) Gt() HpexpireatConditionGt {
	_ = "STUB: not implemented"
	return *new(HpexpireatConditionGt)
}

func (c HpexpireatUnixTimeMilliseconds) Lt() HpexpireatConditionLt {
	_ = "STUB: not implemented"
	return *new(HpexpireatConditionLt)
}

func (c HpexpireatUnixTimeMilliseconds) Fields() HpexpireatFields {
	_ = "STUB: not implemented"
	return *new(HpexpireatFields)
}

type Hpexpiretime Incomplete

func (b Builder) Hpexpiretime() (c Hpexpiretime) {
	_ = "STUB: not implemented"
	return *new(Hpexpiretime)
}

func (c Hpexpiretime) Key(key string) HpexpiretimeKey {
	_ = "STUB: not implemented"
	return *new(HpexpiretimeKey)
}

type HpexpiretimeField Incomplete

func (c HpexpiretimeField) Field(field ...string) HpexpiretimeField {
	_ = "STUB: not implemented"
	return *new(HpexpiretimeField)
}

func (c HpexpiretimeField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HpexpiretimeFields Incomplete

func (c HpexpiretimeFields) Numfields(numfields int64) HpexpiretimeNumfields {
	_ = "STUB: not implemented"
	return *new(HpexpiretimeNumfields)
}

type HpexpiretimeKey Incomplete

func (c HpexpiretimeKey) Fields() HpexpiretimeFields {
	_ = "STUB: not implemented"
	return *new(HpexpiretimeFields)
}

type HpexpiretimeNumfields Incomplete

func (c HpexpiretimeNumfields) Field(field ...string) HpexpiretimeField {
	_ = "STUB: not implemented"
	return *new(HpexpiretimeField)
}

type Hpttl Incomplete

func (b Builder) Hpttl() (c Hpttl) { _ = "STUB: not implemented"; return *new(Hpttl) }

func (c Hpttl) Key(key string) HpttlKey { _ = "STUB: not implemented"; return *new(HpttlKey) }

type HpttlField Incomplete

func (c HpttlField) Field(field ...string) HpttlField {
	_ = "STUB: not implemented"
	return *new(HpttlField)
}

func (c HpttlField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HpttlFields Incomplete

func (c HpttlFields) Numfields(numfields int64) HpttlNumfields {
	_ = "STUB: not implemented"
	return *new(HpttlNumfields)
}

type HpttlKey Incomplete

func (c HpttlKey) Fields() HpttlFields { _ = "STUB: not implemented"; return *new(HpttlFields) }

type HpttlNumfields Incomplete

func (c HpttlNumfields) Field(field ...string) HpttlField {
	_ = "STUB: not implemented"
	return *new(HpttlField)
}

type Hrandfield Incomplete

func (b Builder) Hrandfield() (c Hrandfield) { _ = "STUB: not implemented"; return *new(Hrandfield) }

func (c Hrandfield) Key(key string) HrandfieldKey {
	_ = "STUB: not implemented"
	return *new(HrandfieldKey)
}

type HrandfieldKey Incomplete

func (c HrandfieldKey) Count(count int64) HrandfieldOptionsCount {
	_ = "STUB: not implemented"
	return *new(HrandfieldOptionsCount)
}

func (c HrandfieldKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HrandfieldOptionsCount Incomplete

func (c HrandfieldOptionsCount) Withvalues() HrandfieldOptionsWithvalues {
	_ = "STUB: not implemented"
	return *new(HrandfieldOptionsWithvalues)
}

func (c HrandfieldOptionsCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type HrandfieldOptionsWithvalues Incomplete

func (c HrandfieldOptionsWithvalues) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type Hscan Incomplete

func (b Builder) Hscan() (c Hscan) { _ = "STUB: not implemented"; return *new(Hscan) }

func (c Hscan) Key(key string) HscanKey { _ = "STUB: not implemented"; return *new(HscanKey) }

type HscanCount Incomplete

func (c HscanCount) Novalues() HscanNovalues { _ = "STUB: not implemented"; return *new(HscanNovalues) }

func (c HscanCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HscanCursor Incomplete

func (c HscanCursor) Match(pattern string) HscanMatch {
	_ = "STUB: not implemented"
	return *new(HscanMatch)
}

func (c HscanCursor) Count(count int64) HscanCount {
	_ = "STUB: not implemented"
	return *new(HscanCount)
}

func (c HscanCursor) Novalues() HscanNovalues {
	_ = "STUB: not implemented"
	return *new(HscanNovalues)
}

func (c HscanCursor) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HscanKey Incomplete

func (c HscanKey) Cursor(cursor uint64) HscanCursor {
	_ = "STUB: not implemented"
	return *new(HscanCursor)
}

type HscanMatch Incomplete

func (c HscanMatch) Count(count int64) HscanCount {
	_ = "STUB: not implemented"
	return *new(HscanCount)
}

func (c HscanMatch) Novalues() HscanNovalues { _ = "STUB: not implemented"; return *new(HscanNovalues) }

func (c HscanMatch) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HscanNovalues Incomplete

func (c HscanNovalues) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Hset Incomplete

func (b Builder) Hset() (c Hset) { _ = "STUB: not implemented"; return *new(Hset) }

func (c Hset) Key(key string) HsetKey { _ = "STUB: not implemented"; return *new(HsetKey) }

type HsetFieldValue Incomplete

func (c HsetFieldValue) FieldValue(field string, value string) HsetFieldValue {
	_ = "STUB: not implemented"
	return *new(HsetFieldValue)
}

func (c HsetFieldValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HsetKey Incomplete

func (c HsetKey) FieldValue() HsetFieldValue {
	_ = "STUB: not implemented"
	return *new(HsetFieldValue)
}

type Hsetex Incomplete

func (b Builder) Hsetex() (c Hsetex) { _ = "STUB: not implemented"; return *new(Hsetex) }

func (c Hsetex) Key(key string) HsetexKey { _ = "STUB: not implemented"; return *new(HsetexKey) }

type HsetexConditionFnx Incomplete

func (c HsetexConditionFnx) Ex(ex int64) HsetexExpirationEx {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationEx)
}

func (c HsetexConditionFnx) Px(px int64) HsetexExpirationPx {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationPx)
}

func (c HsetexConditionFnx) Exat(exat int64) HsetexExpirationExat {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationExat)
}

func (c HsetexConditionFnx) Pxat(pxat int64) HsetexExpirationPxat {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationPxat)
}

func (c HsetexConditionFnx) Keepttl() HsetexExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationKeepttl)
}

func (c HsetexConditionFnx) Fields() HsetexFields {
	_ = "STUB: not implemented"
	return *new(HsetexFields)
}

type HsetexConditionFxx Incomplete

func (c HsetexConditionFxx) Ex(ex int64) HsetexExpirationEx {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationEx)
}

func (c HsetexConditionFxx) Px(px int64) HsetexExpirationPx {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationPx)
}

func (c HsetexConditionFxx) Exat(exat int64) HsetexExpirationExat {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationExat)
}

func (c HsetexConditionFxx) Pxat(pxat int64) HsetexExpirationPxat {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationPxat)
}

func (c HsetexConditionFxx) Keepttl() HsetexExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationKeepttl)
}

func (c HsetexConditionFxx) Fields() HsetexFields {
	_ = "STUB: not implemented"
	return *new(HsetexFields)
}

type HsetexExpirationEx Incomplete

func (c HsetexExpirationEx) Fields() HsetexFields {
	_ = "STUB: not implemented"
	return *new(HsetexFields)
}

type HsetexExpirationExat Incomplete

func (c HsetexExpirationExat) Fields() HsetexFields {
	_ = "STUB: not implemented"
	return *new(HsetexFields)
}

type HsetexExpirationKeepttl Incomplete

func (c HsetexExpirationKeepttl) Fields() HsetexFields {
	_ = "STUB: not implemented"
	return *new(HsetexFields)
}

type HsetexExpirationPx Incomplete

func (c HsetexExpirationPx) Fields() HsetexFields {
	_ = "STUB: not implemented"
	return *new(HsetexFields)
}

type HsetexExpirationPxat Incomplete

func (c HsetexExpirationPxat) Fields() HsetexFields {
	_ = "STUB: not implemented"
	return *new(HsetexFields)
}

type HsetexFieldValue Incomplete

func (c HsetexFieldValue) FieldValue(field string, value string) HsetexFieldValue {
	_ = "STUB: not implemented"
	return *new(HsetexFieldValue)
}

func (c HsetexFieldValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HsetexFields Incomplete

func (c HsetexFields) Numfields(numfields int64) HsetexNumfields {
	_ = "STUB: not implemented"
	return *new(HsetexNumfields)
}

type HsetexKey Incomplete

func (c HsetexKey) Fnx() HsetexConditionFnx {
	_ = "STUB: not implemented"
	return *new(HsetexConditionFnx)
}

func (c HsetexKey) Fxx() HsetexConditionFxx {
	_ = "STUB: not implemented"
	return *new(HsetexConditionFxx)
}

func (c HsetexKey) Ex(ex int64) HsetexExpirationEx {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationEx)
}

func (c HsetexKey) Px(px int64) HsetexExpirationPx {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationPx)
}

func (c HsetexKey) Exat(exat int64) HsetexExpirationExat {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationExat)
}

func (c HsetexKey) Pxat(pxat int64) HsetexExpirationPxat {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationPxat)
}

func (c HsetexKey) Keepttl() HsetexExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(HsetexExpirationKeepttl)
}

func (c HsetexKey) Fields() HsetexFields { _ = "STUB: not implemented"; return *new(HsetexFields) }

type HsetexNumfields Incomplete

func (c HsetexNumfields) FieldValue() HsetexFieldValue {
	_ = "STUB: not implemented"
	return *new(HsetexFieldValue)
}

type Hsetnx Incomplete

func (b Builder) Hsetnx() (c Hsetnx) { _ = "STUB: not implemented"; return *new(Hsetnx) }

func (c Hsetnx) Key(key string) HsetnxKey { _ = "STUB: not implemented"; return *new(HsetnxKey) }

type HsetnxField Incomplete

func (c HsetnxField) Value(value string) HsetnxValue {
	_ = "STUB: not implemented"
	return *new(HsetnxValue)
}

type HsetnxKey Incomplete

func (c HsetnxKey) Field(field string) HsetnxField {
	_ = "STUB: not implemented"
	return *new(HsetnxField)
}

type HsetnxValue Incomplete

func (c HsetnxValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Hstrlen Incomplete

func (b Builder) Hstrlen() (c Hstrlen) { _ = "STUB: not implemented"; return *new(Hstrlen) }

func (c Hstrlen) Key(key string) HstrlenKey { _ = "STUB: not implemented"; return *new(HstrlenKey) }

type HstrlenField Incomplete

func (c HstrlenField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HstrlenField) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type HstrlenKey Incomplete

func (c HstrlenKey) Field(field string) HstrlenField {
	_ = "STUB: not implemented"
	return *new(HstrlenField)
}

type Httl Incomplete

func (b Builder) Httl() (c Httl) { _ = "STUB: not implemented"; return *new(Httl) }

func (c Httl) Key(key string) HttlKey { _ = "STUB: not implemented"; return *new(HttlKey) }

type HttlField Incomplete

func (c HttlField) Field(field ...string) HttlField {
	_ = "STUB: not implemented"
	return *new(HttlField)
}

func (c HttlField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HttlFields Incomplete

func (c HttlFields) Numfields(numfields int64) HttlNumfields {
	_ = "STUB: not implemented"
	return *new(HttlNumfields)
}

type HttlKey Incomplete

func (c HttlKey) Fields() HttlFields { _ = "STUB: not implemented"; return *new(HttlFields) }

type HttlNumfields Incomplete

func (c HttlNumfields) Field(field ...string) HttlField {
	_ = "STUB: not implemented"
	return *new(HttlField)
}

type Hvals Incomplete

func (b Builder) Hvals() (c Hvals) { _ = "STUB: not implemented"; return *new(Hvals) }

func (c Hvals) Key(key string) HvalsKey { _ = "STUB: not implemented"; return *new(HvalsKey) }

type HvalsKey Incomplete

func (c HvalsKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c HvalsKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }
