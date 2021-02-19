package application

import (
	"flag"
	"github.com/lukluk/link-proxy/config"
	"github.com/lukluk/link-proxy/internal/adapter/metric"
	"github.com/lukluk/link-proxy/internal/adapter/storage/inmemory"
	"github.com/lukluk/link-proxy/internal/domain/circuitbreaker"
	"github.com/lukluk/link-proxy/internal/domain/validation"
	"github.com/lukluk/link-proxy/internal/port/handler"
	"github.com/rs/zerolog/log"
)

func StartLinkProxy()  {
	configPath := flag.String("config", "config.yaml", "config path")
	flag.Parse()
	log.Debug().Msgf("using config file : %s", *configPath)
	cfg := config.NewConfig(*configPath)
	circuitBreakerData := inmemory.NewCircuitBreakerData()
	statsdClient := metric.NewStatsdClient(cfg)
	cb := circuitbreaker.NewCircuitBreaker(cfg, circuitBreakerData, statsdClient)
	v := validation.NewValidation(cfg)
	go func() {
		cb.RunScheduler()
	}()
	handler.NewEntryPoint(cfg, circuitBreakerData, v, statsdClient).Handler()
	HTTPServer(cfg)

}