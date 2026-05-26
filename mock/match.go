package mock

import (
	"go.uber.org/mock/gomock"
)

func Match(cmd ...string) gomock.Matcher { _ = "STUB: not implemented"; return *new(gomock.Matcher) }

type cmdMatcher struct {
	expect []string
}

func (c *cmdMatcher) Matches(x any) bool { _ = "STUB: not implemented"; return false }

func (c *cmdMatcher) String() string { _ = "STUB: not implemented"; return "" }

func MatchFn(fn func(cmd []string) bool, description ...string) gomock.Matcher {
	_ = "STUB: not implemented"
	return *new(gomock.Matcher)
}

type fnMatcher struct {
	matcher     func(cmd []string) bool
	description []string
}

func (c *fnMatcher) Matches(x any) bool { _ = "STUB: not implemented"; return false }

func (c *fnMatcher) String() string { _ = "STUB: not implemented"; return "" }

func format(v any) string { _ = "STUB: not implemented"; return "" }

func commands(x any) any { _ = "STUB: not implemented"; return *new(any) }
