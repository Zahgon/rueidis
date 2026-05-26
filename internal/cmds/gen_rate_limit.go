// Code generated DO NOT EDIT

package cmds

type Gcra Incomplete

func (b Builder) Gcra() (c Gcra) { _ = "STUB: not implemented"; return *new(Gcra) }

func (c Gcra) Key(key string) GcraKey { _ = "STUB: not implemented"; return *new(GcraKey) }

type GcraCount Incomplete

func (c GcraCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GcraKey Incomplete

func (c GcraKey) MaxBurst(maxBurst int64) GcraMaxBurst {
	_ = "STUB: not implemented"
	return *new(GcraMaxBurst)
}

type GcraMaxBurst Incomplete

func (c GcraMaxBurst) TokensPerPeriod(tokensPerPeriod int64) GcraTokensPerPeriod {
	_ = "STUB: not implemented"
	return *new(GcraTokensPerPeriod)
}

type GcraPeriod Incomplete

func (c GcraPeriod) Count(count int64) GcraCount { _ = "STUB: not implemented"; return *new(GcraCount) }

func (c GcraPeriod) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GcraTokensPerPeriod Incomplete

func (c GcraTokensPerPeriod) Period(period float64) GcraPeriod {
	_ = "STUB: not implemented"
	return *new(GcraPeriod)
}
