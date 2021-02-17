package inmemory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstanceRepo_AddAndInc(t *testing.T) {
	cbData := NewCircuitBreakerData()
	test := cbData.Get("test")
	test.Stats.IncSuccessCount()
	test.Stats.IncErrorCount()
	assert.Equal(t, cbData.Instances()["test"].Stats.SuccessCount(), 1)
	assert.Equal(t, cbData.Instances()["test"].Stats.SuccessCount(), 1)
}

func TestInstanceRepo_ResetStats(t *testing.T) {
	cbData := NewCircuitBreakerData()
	test := cbData.Get("test")
	test.Stats.IncSuccessCount()
	test.Stats.IncErrorCount()
	test.Stats.Reset()
	assert.Equal(t, cbData.Instances()["test"].Stats.SuccessCount(), 0)
	assert.Equal(t, cbData.Instances()["test"].Stats.SuccessCount(), 0)
}