package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDefaultValueStateInterval(t *testing.T) {
	cfg := NewConfig("example.yaml")
	assert.Equal(t, time.Minute, cfg.CircuitBreaker.Durations.StateInterval)
}
