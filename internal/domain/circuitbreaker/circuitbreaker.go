package circuitbreaker

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/lukluk/link-proxy/config"
	"github.com/lukluk/link-proxy/internal/adapter/storage/inmemory"
)
type ICircuitBreaker interface {
	RunScheduler()
}

type CircuitBreaker struct {
	config             config.Config
	circuitBreakerData *inmemory.CircuitBreakerData
	statsdClient *statsd.Client
}

func NewCircuitBreaker(cfg config.Config, circuitBreakerData *inmemory.CircuitBreakerData, statsdClient *statsd.Client) *CircuitBreaker {
	return &CircuitBreaker{
		config:             cfg,
		circuitBreakerData: circuitBreakerData,
		statsdClient: statsdClient,
	}
}