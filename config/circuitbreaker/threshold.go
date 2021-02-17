package circuitbreaker

type Thresholds struct {
	MaxErrorPercent   int `yaml:"max_error_percentage" env-default:"70"`
	MinSuccessPercent int `yaml:"min_success_percentage" env-default:"90"`
}