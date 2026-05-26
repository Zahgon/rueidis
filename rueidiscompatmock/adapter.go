package rueidiscompatmock

import (
	"sync"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

type ClientMock struct {
	raw *mock.Client

	mu        sync.Mutex
	queue     []*expectation
	unmatched []error
	ordered   bool
}

type expectation struct {
	matcher gomock.Matcher
	result  rueidis.RedisResult
}

func NewAdapter(m *mock.Client) *ClientMock { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) MatchExpectationsInOrder(b bool) { _ = "STUB: not implemented"; return }

func (m *ClientMock) ClearExpect() { _ = "STUB: not implemented"; return }

func (m *ClientMock) ExpectationsWereMet() error { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) wire() { _ = "STUB: not implemented"; return }

func (m *ClientMock) consume(n int, cmds []rueidis.Completed) []rueidis.RedisResult {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) matchLocked(cmd rueidis.Completed) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (m *ClientMock) push(matcher gomock.Matcher, defaultResult rueidis.RedisResult) *expectation {
	_ = "STUB: not implemented"
	return nil
}

const keepTTL = -1

func usePrecise(d time.Duration) bool { _ = "STUB: not implemented"; return false }

func formatMs(d time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

func formatSec(d time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

func str(arg any) string { _ = "STUB: not implemented"; return "" }

func (m *ClientMock) ExpectGet(key string) *ExpectedString { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectSet(key string, value any, expiration time.Duration) *ExpectedStatus {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectSetNX(key string, value any, expiration time.Duration) *ExpectedBool {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectGetSet(key string, value any) *ExpectedString {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectAppend(key, value string) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectStrLen(key string) *ExpectedInt { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectDel(keys ...string) *ExpectedInt { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectExists(keys ...string) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectType(key string) *ExpectedStatus { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectTTL(key string) *ExpectedDuration { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectExpire(key string, expiration time.Duration) *ExpectedBool {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectPing() *ExpectedStatus { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectEcho(message any) *ExpectedString { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectIncr(key string) *ExpectedInt { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectIncrBy(key string, value int64) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectDecr(key string) *ExpectedInt { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectDecrBy(key string, value int64) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectMGet(keys ...string) *ExpectedSlice {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectMSet(values ...any) *ExpectedStatus {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectHGet(key, field string) *ExpectedString {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectHSet(key string, values ...any) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectHDel(key string, fields ...string) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectHGetAll(key string) *ExpectedStringStringMap {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectLPush(key string, elements ...any) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectRPush(key string, elements ...any) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectLPop(key string) *ExpectedString { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectRPop(key string) *ExpectedString { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectLLen(key string) *ExpectedInt { _ = "STUB: not implemented"; return nil }

func (m *ClientMock) ExpectSAdd(key string, members ...any) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectSRem(key string, members ...any) *ExpectedInt {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectSMembers(key string) *ExpectedStringSlice {
	_ = "STUB: not implemented"
	return nil
}

func (m *ClientMock) ExpectEval(script string, keys []string, args ...any) *ExpectedCmd {
	_ = "STUB: not implemented"
	return nil
}

func setMatcher(key string, value any, expiration time.Duration) gomock.Matcher {
	_ = "STUB: not implemented"
	return *new(gomock.Matcher)
}

func setNXMatcher(key string, value any, expiration time.Duration) gomock.Matcher {
	_ = "STUB: not implemented"
	return *new(gomock.Matcher)
}

func delMatcher(keys ...string) gomock.Matcher {
	_ = "STUB: not implemented"
	return *new(gomock.Matcher)
}

func hsetArgsToSlice(values []any) []string { _ = "STUB: not implemented"; return nil }

func evalMatcher(script string, keys []string, args ...any) gomock.Matcher {
	_ = "STUB: not implemented"
	return *new(gomock.Matcher)
}
