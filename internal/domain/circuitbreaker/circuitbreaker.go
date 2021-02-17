package circuitbreaker

import (
	"github.com/lukluk/link-proxy/config"
	"github.com/lukluk/link-proxy/internal/adapter/storage/inmemory"
)
type ICircuitBreaker interface {
	RunScheduler()
}

type CircuitBreaker struct {
	config             config.Config
	circuitBreakerData *inmemory.CircuitBreakerData
}

func NewCircuitBreaker(cfg config.Config, circuitBreakerData *inmemory.CircuitBreakerData) *CircuitBreaker {
	return &CircuitBreaker{
		config:             cfg,
		circuitBreakerData: circuitBreakerData,
	}
}