package circuitbreaker

import "time"

type Duration struct {
	StateInterval  time.Duration `yaml:"state_interval" env-default:"1m"`
	ClosedDuration time.Duration `yaml:"closed_duration" env-default:"15m"`
}