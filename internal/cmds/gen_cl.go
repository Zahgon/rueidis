// Code generated DO NOT EDIT

package cmds

type ClThrottle Incomplete

func (b Builder) ClThrottle() (c ClThrottle) { _ = "STUB: not implemented"; return *new(ClThrottle) }

func (c ClThrottle) Key(key string) ClThrottleKey {
	_ = "STUB: not implemented"
	return *new(ClThrottleKey)
}

type ClThrottleCountPerPeriod Incomplete

func (c ClThrottleCountPerPeriod) Period(period int64) ClThrottlePeriod {
	_ = "STUB: not implemented"
	return *new(ClThrottlePeriod)
}

type ClThrottleKey Incomplete

func (c ClThrottleKey) MaxBurst(maxBurst int64) ClThrottleMaxBurst {
	_ = "STUB: not implemented"
	return *new(ClThrottleMaxBurst)
}

type ClThrottleMaxBurst Incomplete

func (c ClThrottleMaxBurst) CountPerPeriod(countPerPeriod int64) ClThrottleCountPerPeriod {
	_ = "STUB: not implemented"
	return *new(ClThrottleCountPerPeriod)
}

type ClThrottlePeriod Incomplete

func (c ClThrottlePeriod) Quantity(quantity int64) ClThrottleQuantity {
	_ = "STUB: not implemented"
	return *new(ClThrottleQuantity)
}

func (c ClThrottlePeriod) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClThrottleQuantity Incomplete

func (c ClThrottleQuantity) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
