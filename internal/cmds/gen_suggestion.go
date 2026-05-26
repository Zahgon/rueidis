// Code generated DO NOT EDIT

package cmds

type FtSugadd Incomplete

func (b Builder) FtSugadd() (c FtSugadd) { _ = "STUB: not implemented"; return *new(FtSugadd) }

func (c FtSugadd) Key(key string) FtSugaddKey { _ = "STUB: not implemented"; return *new(FtSugaddKey) }

type FtSugaddIncrementScoreIncr Incomplete

func (c FtSugaddIncrementScoreIncr) Payload(payload string) FtSugaddPayload {
	_ = "STUB: not implemented"
	return *new(FtSugaddPayload)
}

func (c FtSugaddIncrementScoreIncr) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSugaddKey Incomplete

func (c FtSugaddKey) String(string string) FtSugaddString {
	_ = "STUB: not implemented"
	return *new(FtSugaddString)
}

type FtSugaddPayload Incomplete

func (c FtSugaddPayload) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSugaddScore Incomplete

func (c FtSugaddScore) Incr() FtSugaddIncrementScoreIncr {
	_ = "STUB: not implemented"
	return *new(FtSugaddIncrementScoreIncr)
}

func (c FtSugaddScore) Payload(payload string) FtSugaddPayload {
	_ = "STUB: not implemented"
	return *new(FtSugaddPayload)
}

func (c FtSugaddScore) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSugaddString Incomplete

func (c FtSugaddString) Score(score float64) FtSugaddScore {
	_ = "STUB: not implemented"
	return *new(FtSugaddScore)
}

type FtSugdel Incomplete

func (b Builder) FtSugdel() (c FtSugdel) { _ = "STUB: not implemented"; return *new(FtSugdel) }

func (c FtSugdel) Key(key string) FtSugdelKey { _ = "STUB: not implemented"; return *new(FtSugdelKey) }

type FtSugdelKey Incomplete

func (c FtSugdelKey) String(string string) FtSugdelString {
	_ = "STUB: not implemented"
	return *new(FtSugdelString)
}

type FtSugdelString Incomplete

func (c FtSugdelString) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSugget Incomplete

func (b Builder) FtSugget() (c FtSugget) { _ = "STUB: not implemented"; return *new(FtSugget) }

func (c FtSugget) Key(key string) FtSuggetKey { _ = "STUB: not implemented"; return *new(FtSuggetKey) }

type FtSuggetFuzzy Incomplete

func (c FtSuggetFuzzy) Withscores() FtSuggetWithscores {
	_ = "STUB: not implemented"
	return *new(FtSuggetWithscores)
}

func (c FtSuggetFuzzy) Withpayloads() FtSuggetWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSuggetWithpayloads)
}

func (c FtSuggetFuzzy) Max(max int64) FtSuggetMax {
	_ = "STUB: not implemented"
	return *new(FtSuggetMax)
}

func (c FtSuggetFuzzy) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSuggetKey Incomplete

func (c FtSuggetKey) Prefix(prefix string) FtSuggetPrefix {
	_ = "STUB: not implemented"
	return *new(FtSuggetPrefix)
}

type FtSuggetMax Incomplete

func (c FtSuggetMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSuggetPrefix Incomplete

func (c FtSuggetPrefix) Fuzzy() FtSuggetFuzzy {
	_ = "STUB: not implemented"
	return *new(FtSuggetFuzzy)
}

func (c FtSuggetPrefix) Withscores() FtSuggetWithscores {
	_ = "STUB: not implemented"
	return *new(FtSuggetWithscores)
}

func (c FtSuggetPrefix) Withpayloads() FtSuggetWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSuggetWithpayloads)
}

func (c FtSuggetPrefix) Max(max int64) FtSuggetMax {
	_ = "STUB: not implemented"
	return *new(FtSuggetMax)
}

func (c FtSuggetPrefix) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSuggetWithpayloads Incomplete

func (c FtSuggetWithpayloads) Max(max int64) FtSuggetMax {
	_ = "STUB: not implemented"
	return *new(FtSuggetMax)
}

func (c FtSuggetWithpayloads) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSuggetWithscores Incomplete

func (c FtSuggetWithscores) Withpayloads() FtSuggetWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSuggetWithpayloads)
}

func (c FtSuggetWithscores) Max(max int64) FtSuggetMax {
	_ = "STUB: not implemented"
	return *new(FtSuggetMax)
}

func (c FtSuggetWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSuglen Incomplete

func (b Builder) FtSuglen() (c FtSuglen) { _ = "STUB: not implemented"; return *new(FtSuglen) }

func (c FtSuglen) Key(key string) FtSuglenKey { _ = "STUB: not implemented"; return *new(FtSuglenKey) }

type FtSuglenKey Incomplete

func (c FtSuglenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
