package upstream

type circuitBreakerKey struct {
	RequestBodyJsonPaths []string `yaml:"request_body_json_paths"`
}
