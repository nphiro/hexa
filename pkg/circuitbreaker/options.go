package circuitbreaker

import "time"

type Options struct {
	failureThresholdCount uint
	resetTimeout          time.Duration
}

func NewOptions() *Options {
	return &Options{
		failureThresholdCount: 3,
		resetTimeout:          3 * time.Minute,
	}
}

func (o *Options) WithFailureThresholdCount(count uint) *Options {
	o.failureThresholdCount = count
	return o
}

func (o *Options) WithResetTimeout(timeout time.Duration) *Options {
	o.resetTimeout = timeout
	return o
}
