package upstream

type Upstream struct {
	Host              string            `yaml:"host"`
	CircuitBreakerKey circuitBreakerKey `yaml:"circuit_breaker_key"`
	SuccessCondition  condition         `yaml:"success_condition"`
}