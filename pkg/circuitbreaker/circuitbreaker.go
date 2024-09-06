package circuitbreaker

import (
	"sync"
	"time"
)

type CircuitBreaker interface {
	Exec(func() error) error
}

func NewCircuitBreaker(opts *Options) CircuitBreaker {
	if opts == nil {
		opts = NewOptions()
	}

	return &circuitBreaker{
		failureThresholdCount: opts.failureThresholdCount,
		resetTimeout:          opts.resetTimeout,
	}
}

type circuitBreaker struct {
	failureCount          uint
	failureThresholdCount uint
	open                  bool
	resetTimeout          time.Duration
	lastFailureTime       time.Time
	mutex                 sync.Mutex
}

func (cb *circuitBreaker) Exec(f func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.checkResetTimeout()

	if cb.open {
		return ErrCircuitBreakerOpen
	}

	err := f()
	if err != nil {
		cb.recordFailure()
		return err
	}

	cb.failureCount = 0
	return nil
}

func (cb *circuitBreaker) checkResetTimeout() {
	if time.Since(cb.lastFailureTime) > cb.resetTimeout {
		cb.failureCount = 0
		cb.open = false
	}
}

func (cb *circuitBreaker) recordFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()
	if cb.failureCount >= cb.failureThresholdCount {
		cb.open = true
	}
}
