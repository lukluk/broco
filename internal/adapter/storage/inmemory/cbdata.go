package inmemory

import (
	"github.com/lukluk/link-proxy/internal/adapter/storage/inmemory/scheme"
)

type CircuitBreakerData struct {
	instances                  map[string]*scheme.Instance
}

func NewCircuitBreakerData() *CircuitBreakerData {
	return &CircuitBreakerData{
		make(map[string]*scheme.Instance),
	}
}

