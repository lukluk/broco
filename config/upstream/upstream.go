package upstream

type Upstream struct {
	Host                 string            `yaml:"host"`
	CircuitBreakerKey    circuitBreakerKey `yaml:"circuit_breaker_key"`
	ExtraErrorConditions []condition       `yaml:"extra_error_condition"`
}