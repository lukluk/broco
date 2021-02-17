package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lukluk/link-proxy/config/circuitbreaker"
	"github.com/lukluk/link-proxy/config/upstream"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Upstreams      map[string]upstream.Upstream `yaml:"upstreams"`
	StatsdHost     string              `yaml:"statsd_host"`
	CircuitBreaker circuitbreaker.CircuitBreaker `yaml:"circuit_breaker"`
	Port           string	`yaml:"port" env-default:"8080"`
}

func NewConfig(configPath string) Config {
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Panic().Msgf("failed to read %s, err: %v", configPath, err)
	}
	return cfg
}