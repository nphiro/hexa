package circuitbreaker

import "errors"

var ErrCircuitBreakerOpen = errors.New("circuit breaker is open")
