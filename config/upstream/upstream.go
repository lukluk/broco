package upstream

type Upstream struct {
	Host               string            `yaml:"host"`
	CircuitBreakerKey  circuitBreakerKey `yaml:"circuit_breaker_key"`
	AddErrorConditions []condition       `yaml:"add_error_conditions"`
}