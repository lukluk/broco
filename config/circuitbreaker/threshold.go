package circuitbreaker

type Thresholds struct {
	MaxErrorPercent int `yaml:"max_error_percentage" env-default:"70"`
	MinErrorPercent int `yaml:"min_error_percentage" env-default:"10"`
}