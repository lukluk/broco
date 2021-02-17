package inmemory

import (
	"github.com/lukluk/link-proxy/internal/adapter/storage/inmemory/scheme"
	"time"
)

func (i *CircuitBreakerData) Instances() map[string]*scheme.Instance {
	return i.instances
}

func (i *CircuitBreakerData) Get(key string) *scheme.Instance {
	if val, ok := i.instances[key]; ok {
		return val
	}
	i.instances[key] = &scheme.Instance{
		Traffic: scheme.Traffic{},
		Stats:   scheme.Stats{},
	}
	return i.instances[key]
}
func (i *CircuitBreakerData) IncSuccessCount(key string) {
	i.instances[key].Stats.IncSuccessCount()
}

func (i *CircuitBreakerData) IncErrorCount(key string) {
	i.instances[key].Stats.ErrorCount()
}

func (i *CircuitBreakerData) CloseUntil(key string, until time.Time) {
	i.instances[key].Traffic.CloseUntil(until)
}

func (i *CircuitBreakerData) HalfOpen(key string) {
	i.instances[key].Traffic.HalfOpen()
}

func (i *CircuitBreakerData) FullyOpen(key string) {
	i.instances[key].Traffic.FullyOpen()
}

func (i *CircuitBreakerData) ResetStats(key string) {
	i.instances[key].Traffic.ResetTrafficCount()
	i.instances[key].Stats.Reset()
}