package rueidis

import (
	"bytes"
	"errors"
	"io"
	"time"
	"unsafe"
)

const messageStructSize = int(unsafe.Sizeof(RedisMessage{}))

// Nil represents a Redis Nil message
var Nil = &RedisError{typ: typeNull}

// ErrParse is a parse error that occurs when a Redis message cannot be parsed correctly.
var errParse = errors.New("rueidis: parse error")

// IsRedisNil is a handy method to check if the error is a redis nil response.
// All redis nil responses returned as an error.
func IsRedisNil(err error) bool {
	_ = "STUB: not implemented"

	// IsParseErr checks if the error is a parse error
	return false
}

func IsParseErr(err error) bool { _ = "STUB: not implemented"; return false }

// IsRedisBusyGroup checks if it is a redis BUSYGROUP message.
func IsRedisBusyGroup(err error) bool { _ = "STUB: not implemented"; return false }

// IsRedisErr is a handy method to check if the error is a redis ERR response.
func IsRedisErr(err error) (ret *RedisError, ok bool) { _ = "STUB: not implemented"; return nil, false }

// RedisError is an error response or a nil message from the redis instance
type RedisError RedisMessage

// string retrieves the contained string of the RedisError
func (m *RedisError) string() string { _ = "STUB: not implemented"; return "" }

func (r *RedisError) Error() string { _ = "STUB: not implemented"; return "" }

// IsNil checks if it is a redis nil message.
func (r *RedisError) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsMoved checks if it is a redis MOVED message and returns the moved address.
func (r *RedisError) IsMoved() (addr string, ok bool) { _ = "STUB: not implemented"; return "", false }

// IsAsk checks if it is a redis ASK message and returns ask address.
func (r *RedisError) IsAsk() (addr string, ok bool) { _ = "STUB: not implemented"; return "", false }

// IsRedirect checks if it is a redis REDIRECT message and returns redirect address.
func (r *RedisError) IsRedirect() (addr string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func fixIPv6HostPort(addr string) string { _ = "STUB: not implemented"; return "" }

// skip ipv4 and enclosed ipv6

// IsTryAgain checks if it is a redis TRYAGAIN message and returns ask address.
func (r *RedisError) IsTryAgain() bool { _ = "STUB: not implemented"; return false }

// IsLoading checks if it is a redis LOADING message
func (r *RedisError) IsLoading() bool { _ = "STUB: not implemented"; return false }

// IsClusterDown checks if it is a redis CLUSTERDOWN message and returns ask address.
func (r *RedisError) IsClusterDown() bool { _ = "STUB: not implemented"; return false }

// IsNoScript checks if it is a redis NOSCRIPT message.
func (r *RedisError) IsNoScript() bool { _ = "STUB: not implemented"; return false }

// IsBusyGroup checks if it is a redis BUSYGROUP message.
func (r *RedisError) IsBusyGroup() bool { _ = "STUB: not implemented"; return false }

func newResult(val RedisMessage, err error) RedisResult {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func newErrResult(err error) RedisResult { _ = "STUB: not implemented"; return *new(RedisResult) }

// RedisResult is the return struct from Client.Do or Client.DoCache
// it contains either a redis response or an underlying error (ex. network timeout).
type RedisResult struct {
	err error
	val RedisMessage
}

// NonRedisError can be used to check if there is an underlying error (ex. network timeout).
func (r RedisResult) NonRedisError() error {
	_ = "STUB: not implemented"

	// Error returns either underlying error or redis error or nil
	return nil
}

func (r RedisResult) Error() (err error) { _ = "STUB: not implemented"; return nil }

// ToMessage retrieves the RedisMessage
func (r RedisResult) ToMessage() (v RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

// ToInt64 delegates to RedisMessage.ToInt64
func (r RedisResult) ToInt64() (v int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToBool delegates to RedisMessage.ToBool
func (r RedisResult) ToBool() (v bool, err error) { _ = "STUB: not implemented"; return false, nil }

// ToFloat64 delegates to RedisMessage.ToFloat64
func (r RedisResult) ToFloat64() (v float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToString delegates to RedisMessage.ToString
func (r RedisResult) ToString() (v string, err error) { _ = "STUB: not implemented"; return "", nil }

// AsReader delegates to RedisMessage.AsReader
func (r RedisResult) AsReader() (v io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// AsBytes delegates to RedisMessage.AsBytes
func (r RedisResult) AsBytes() (v []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// DecodeJSON delegates to RedisMessage.DecodeJSON
func (r RedisResult) DecodeJSON(v any) (err error) { _ = "STUB: not implemented"; return nil }

// AsInt64 delegates to RedisMessage.AsInt64
func (r RedisResult) AsInt64() (v int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// AsUint64 delegates to RedisMessage.AsUint64
func (r RedisResult) AsUint64() (v uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

// AsBool delegates to RedisMessage.AsBool
func (r RedisResult) AsBool() (v bool, err error) { _ = "STUB: not implemented"; return false, nil }

// AsFloat64 delegates to RedisMessage.AsFloat64
func (r RedisResult) AsFloat64() (v float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToArray delegates to RedisMessage.ToArray
func (r RedisResult) ToArray() (v []RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrSlice delegates to RedisMessage.AsStrSlice
func (r RedisResult) AsStrSlice() (v []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsIntSlice delegates to RedisMessage.AsIntSlice
func (r RedisResult) AsIntSlice() (v []int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsFloatSlice delegates to RedisMessage.AsFloatSlice
func (r RedisResult) AsFloatSlice() (v []float64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsBoolSlice delegates to RedisMessage.AsBoolSlice
func (r RedisResult) AsBoolSlice() (v []bool, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRangeEntry delegates to RedisMessage.AsXRangeEntry
func (r RedisResult) AsXRangeEntry() (v XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return *new(XRangeEntry), nil
}

// AsXRange delegates to RedisMessage.AsXRange
func (r RedisResult) AsXRange() (v []XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsZScore delegates to RedisMessage.AsZScore
func (r RedisResult) AsZScore() (v ZScore, err error) {
	_ = "STUB: not implemented"
	return *new(ZScore), nil
}

// AsZScores delegates to RedisMessage.AsZScores
func (r RedisResult) AsZScores() (v []ZScore, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRead delegates to RedisMessage.AsXRead
func (r RedisResult) AsXRead() (v map[string][]XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRangeSlice delegates to RedisMessage.AsXRangeSlice
func (r RedisResult) AsXRangeSlice() (v XRangeSlice, err error) {
	_ = "STUB: not implemented"
	return *new(XRangeSlice), nil
}

// AsXRangeSlices delegates to RedisMessage.AsXRangeSlices
func (r RedisResult) AsXRangeSlices() (v []XRangeSlice, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXReadSlices delegates to RedisMessage.AsXReadSlices
func (r RedisResult) AsXReadSlices() (v map[string][]XRangeSlice, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RedisResult) AsLMPop() (v KeyValues, err error) {
	_ = "STUB: not implemented"
	return *new(KeyValues), nil
}

func (r RedisResult) AsZMPop() (v KeyZScores, err error) {
	_ = "STUB: not implemented"
	return *new(KeyZScores), nil
}

func (r RedisResult) AsFtSearch() (total int64, docs []FtSearchDoc, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (r RedisResult) AsFtAggregate() (total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (r RedisResult) AsFtAggregateCursor() (cursor, total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

func (r RedisResult) AsGeosearch() (locations []GeoLocation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsMap delegates to RedisMessage.AsMap
func (r RedisResult) AsMap() (v map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrMap delegates to RedisMessage.AsStrMap
func (r RedisResult) AsStrMap() (v map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsIntMap delegates to RedisMessage.AsIntMap
func (r RedisResult) AsIntMap() (v map[string]int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsScanEntry delegates to RedisMessage.AsScanEntry.
func (r RedisResult) AsScanEntry() (v ScanEntry, err error) {
	_ = "STUB: not implemented"
	return *new(ScanEntry), nil
}

// ToMap delegates to RedisMessage.ToMap
func (r RedisResult) ToMap() (v map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToAny delegates to RedisMessage.ToAny
func (r RedisResult) ToAny() (v any, err error) { _ = "STUB: not implemented"; return *new(any), nil }

// IsCacheHit delegates to RedisMessage.IsCacheHit
func (r RedisResult) IsCacheHit() bool { _ = "STUB: not implemented"; return false }

// CacheTTL delegates to RedisMessage.CacheTTL
func (r RedisResult) CacheTTL() int64 { _ = "STUB: not implemented"; return 0 }

// CachePTTL delegates to RedisMessage.CachePTTL
func (r RedisResult) CachePTTL() int64 { _ = "STUB: not implemented"; return 0 }

// CachePXAT delegates to RedisMessage.CachePXAT
func (r RedisResult) CachePXAT() int64 { _ = "STUB: not implemented"; return 0 }

// String returns human-readable representation of RedisResult
func (r *RedisResult) String() string { _ = "STUB: not implemented"; return "" }

type prettyRedisResult RedisResult

// MarshalJSON implements json.Marshaler interface
func (r *prettyRedisResult) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RedisMessage is a redis response message, it may be a nil response
type RedisMessage struct {
	attrs *RedisMessage
	bytes *byte
	array *RedisMessage

	// intlen is used for a simple number or
	// in conjunction with an array or bytes to store the length of array or string
	intlen int64
	typ    byte
	ttl    [7]byte
}

func (m *RedisMessage) string() string { _ = "STUB: not implemented"; return "" }

func (m *RedisMessage) values() []RedisMessage { _ = "STUB: not implemented"; return nil }

func (m *RedisMessage) setString(s string) { _ = "STUB: not implemented"; return }

func (m *RedisMessage) setValues(values []RedisMessage) { _ = "STUB: not implemented"; return }

func (m *RedisMessage) cachesize() int {
	_ = "STUB: not implemented"
	// typ (1) + length (8) TODO: can we use VarInt instead of fixed 8 bytes for length?
	return 0
}

func (m *RedisMessage) serialize(o *bytes.Buffer) {
	_ = "STUB: not implemented"
	// TODO: can we use VarInt instead of fixed 8 bytes for length?
	return
}

var ErrCacheUnmarshal = errors.New("cache unmarshal error")

func (m *RedisMessage) unmarshalView(c int64, buf []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: can we use VarInt instead of fixed 8 bytes for length?

// CacheSize returns the buffer size needed by the CacheMarshal.
func (m *RedisMessage) CacheSize() int { _ = "STUB: not implemented"; return 0 }

// 7 for ttl

// CacheMarshal writes serialized RedisMessage to the provided buffer.
// If the provided buffer is nil, CacheMarshal will allocate one.
// Note that an output format is not compatible with different client versions.
func (m *RedisMessage) CacheMarshal(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

// CacheUnmarshalView construct the RedisMessage from the buffer produced by CacheMarshal.
// Note that the buffer can't be reused after CacheUnmarshalView since it uses unsafe.String on top of the buffer.
func (m *RedisMessage) CacheUnmarshalView(buf []byte) error { _ = "STUB: not implemented"; return nil }

// IsNil check if the message is a redis nil response
func (m *RedisMessage) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsInt64 check if the message is a redis RESP3 int response
func (m *RedisMessage) IsInt64() bool { _ = "STUB: not implemented"; return false }

// IsFloat64 check if the message is a redis RESP3 double response
func (m *RedisMessage) IsFloat64() bool { _ = "STUB: not implemented"; return false }

// IsString check if the message is a redis string response
func (m *RedisMessage) IsString() bool { _ = "STUB: not implemented"; return false }

// IsBool check if the message is a redis RESP3 bool response
func (m *RedisMessage) IsBool() bool { _ = "STUB: not implemented"; return false }

// IsArray check if the message is a redis array response
func (m *RedisMessage) IsArray() bool { _ = "STUB: not implemented"; return false }

// IsMap check if the message is a redis RESP3 map response
func (m *RedisMessage) IsMap() bool { _ = "STUB: not implemented"; return false }

// Error check if the message is a redis error response, including nil response
func (m *RedisMessage) Error() error { _ = "STUB: not implemented"; return nil }

// kvrocks: https://github.com/redis/rueidis/issues/152#issuecomment-1333923750

// ToString check if the message is a redis string response and return it
func (m *RedisMessage) ToString() (val string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AsReader check if the message is a redis string response and wrap it with the strings.NewReader
func (m *RedisMessage) AsReader() (reader io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// AsBytes check if the message is a redis string response and return it as an immutable []byte
func (m *RedisMessage) AsBytes() (bs []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// DecodeJSON check if the message is a redis string response and treat it as JSON, then unmarshal it into the provided value
func (m *RedisMessage) DecodeJSON(v any) (err error) { _ = "STUB: not implemented"; return nil }

// AsInt64 check if the message is a redis string response and parse it as int64
func (m *RedisMessage) AsInt64() (val int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// AsUint64 check if the message is a redis string response and parse it as uint64
func (m *RedisMessage) AsUint64() (val uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

// AsBool checks if the message is a non-nil response and parses it as bool
func (m *RedisMessage) AsBool() (val bool, err error) { _ = "STUB: not implemented"; return false, nil }

// AsFloat64 check if the message is a redis string response and parse it as float64
func (m *RedisMessage) AsFloat64() (val float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToInt64 check if the message is a redis RESP3 int response and return it
func (m *RedisMessage) ToInt64() (val int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToBool check if the message is a redis RESP3 bool response and return it
func (m *RedisMessage) ToBool() (val bool, err error) { _ = "STUB: not implemented"; return false, nil }

// ToFloat64 check if the message is a redis RESP3 double response and return it
func (m *RedisMessage) ToFloat64() (val float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToArray check if the message is a redis array/set response and return it
func (m *RedisMessage) ToArray() ([]RedisMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrSlice check if the message is a redis array/set response and convert to []string.
// redis nil element and other non-string elements will be present as zero.
func (m *RedisMessage) AsStrSlice() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// AsIntSlice check if the message is a redis array/set response and convert to []int64.
// redis nil element and other non-integer elements will be present as zero.
func (m *RedisMessage) AsIntSlice() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

// AsFloatSlice check if the message is a redis array/set response and convert to []float64.
// redis nil element and other non-float elements will be present as zero.
func (m *RedisMessage) AsFloatSlice() ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsBoolSlice checks if the message is a redis array/set response and converts it to []bool.
// Redis nil elements and other non-boolean elements will be represented as false.
func (m *RedisMessage) AsBoolSlice() ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

// Ignore error, non-boolean values will be false

// XRangeEntry is the element type of both XRANGE and XREVRANGE command response array
type XRangeEntry struct {
	FieldValues map[string]string
	ID          string
}

// AsXRangeEntry check if the message is a redis array/set response of length 2 and convert to XRangeEntry
func (m *RedisMessage) AsXRangeEntry() (XRangeEntry, error) {
	_ = "STUB: not implemented"
	return *new(XRangeEntry), nil
}

// AsXRange check if the message is a redis array/set response and convert to []XRangeEntry
func (m *RedisMessage) AsXRange() ([]XRangeEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRead converts XREAD/XREADGRUOP response to map[string][]XRangeEntry
func (m *RedisMessage) AsXRead() (ret map[string][]XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// New slice-based structures that preserve order and duplicates
type XRangeSlice struct {
	ID          string
	FieldValues []XRangeFieldValue
}

type XRangeFieldValue struct {
	Field string
	Value string
}

// AsXRangeSlice converts a RedisMessage to XRangeSlice (preserves order and duplicates)
func (m *RedisMessage) AsXRangeSlice() (XRangeSlice, error) {
	_ = "STUB: not implemented"
	return *new(XRangeSlice), nil
}

// Handle the field-values array

// Convert pairs to slice (preserving order)

// AsXRangeSlices converts multiple XRange entries to slice format
func (m *RedisMessage) AsXRangeSlices() ([]XRangeSlice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXReadSlices converts XREAD/XREADGROUP response to use slice format
func (m *RedisMessage) AsXReadSlices() (map[string][]XRangeSlice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZScore is the element type of ZRANGE WITHSCORES, ZDIFF WITHSCORES and ZPOPMAX command response
type ZScore struct {
	Member string
	Score  float64
}

func toZScore(values []RedisMessage) (s ZScore, err error) {
	_ = "STUB: not implemented"
	return *new(ZScore), nil
}

// AsZScore converts ZPOPMAX and ZPOPMIN command with count 1 response to a single ZScore
func (m *RedisMessage) AsZScore() (s ZScore, err error) {
	_ = "STUB: not implemented"
	return *new(ZScore), nil
}

// AsZScores converts ZRANGE WITHSCORES, ZDIFF WITHSCORES and ZPOPMAX/ZPOPMIN command with count > 1 responses to []ZScore
func (m *RedisMessage) AsZScores() ([]ZScore, error) { _ = "STUB: not implemented"; return nil, nil }

// ScanEntry is the element type of both SCAN, SSCAN, HSCAN and ZSCAN command response.
type ScanEntry struct {
	Elements []string
	Cursor   uint64
}

// AsScanEntry check if the message is a redis array/set response of length 2 and convert to ScanEntry.
func (m *RedisMessage) AsScanEntry() (e ScanEntry, err error) {
	_ = "STUB: not implemented"
	return *new(ScanEntry), nil
}

// AsMap check if the message is a redis array/set response and convert to map[string]RedisMessage
func (m *RedisMessage) AsMap() (map[string]RedisMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrMap check if the message is a redis map/array/set response and convert to map[string]string.
// redis nil element and other non-string elements will be present as zero.
func (m *RedisMessage) AsStrMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsIntMap check if the message is a redis map/array/set response and convert to map[string]int64.
// redis nil element and other non-integer elements will be present as zero.
func (m *RedisMessage) AsIntMap() (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type KeyValues struct {
	Key    string
	Values []string
}

func (m *RedisMessage) AsLMPop() (kvs KeyValues, err error) {
	_ = "STUB: not implemented"
	return *new(KeyValues), nil
}

type KeyZScores struct {
	Key    string
	Values []ZScore
}

func (m *RedisMessage) AsZMPop() (kvs KeyZScores, err error) {
	_ = "STUB: not implemented"
	return *new(KeyZScores), nil
}

type FtSearchDoc struct {
	Doc   map[string]string
	Key   string
	Score float64
}

func (m *RedisMessage) AsFtSearch() (total int64, docs []FtSearchDoc, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (m *RedisMessage) AsFtAggregate() (total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (m *RedisMessage) AsFtAggregateCursor() (cursor, total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

func (m *RedisMessage) AsGeosearch() ([]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//name

//distance

//hash

//coordinates

// ToMap check if the message is a redis RESP3 map response and return it
func (m *RedisMessage) ToMap() (map[string]RedisMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToAny turns the message into go any value
func (m *RedisMessage) ToAny() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// IsCacheHit check if the message is from the client side cache
func (m *RedisMessage) IsCacheHit() bool { _ = "STUB: not implemented"; return false }

// CacheTTL returns the remaining TTL in seconds of client side cache
func (m *RedisMessage) CacheTTL() (ttl int64) { _ = "STUB: not implemented"; return 0 }

// CachePTTL returns the remaining PTTL in seconds of client side cache
func (m *RedisMessage) CachePTTL() int64 { _ = "STUB: not implemented"; return 0 }

// CachePXAT returns the remaining PXAT in seconds of client side cache
func (m *RedisMessage) CachePXAT() int64 { _ = "STUB: not implemented"; return 0 }

func (m *RedisMessage) relativePTTL(now time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func (m *RedisMessage) getExpireAt() int64 { _ = "STUB: not implemented"; return 0 }

func (m *RedisMessage) setExpireAt(pttl int64) { _ = "STUB: not implemented"; return }

func toMap(values []RedisMessage) (map[string]RedisMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *RedisMessage) approximateSize() (s int) { _ = "STUB: not implemented"; return 0 }

// String returns the human-readable representation of RedisMessage
func (m *RedisMessage) String() string { _ = "STUB: not implemented"; return "" }

type prettyRedisMessage RedisMessage

func (m *prettyRedisMessage) string() string { _ = "STUB: not implemented"; return "" }

func (m *prettyRedisMessage) values() []RedisMessage { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler interface
func (m *prettyRedisMessage) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func slicemsg(typ byte, values []RedisMessage) RedisMessage {
	_ = "STUB: not implemented"
	return *new(RedisMessage)
}

func strmsg(typ byte, value string) RedisMessage {
	_ = "STUB: not implemented"
	return *new(RedisMessage)
}
