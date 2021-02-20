package upstream

import "time"

type Upstream struct {
	Host               string            `yaml:"host"`
	Timeout			   time.Duration	 `yaml:"timeout" env-default:"15s"`
	CircuitBreakerKey  circuitBreakerKey `yaml:"circuit_breaker_key"`
	AddErrorConditions []condition       `yaml:"add_error_conditions"`
}