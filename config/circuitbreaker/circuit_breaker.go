package circuitbreaker

type CircuitBreaker struct {
	Thresholds      Thresholds `yaml:"thresholds"`
	Durations       Duration   `yaml:"durations"`
	Fallback		Fallback   `yaml:"fallback"`
}