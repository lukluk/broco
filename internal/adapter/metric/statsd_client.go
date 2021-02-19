package metric

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/lukluk/link-proxy/config"
)


func NewStatsdClient(cfg config.Config) *statsd.Client {
	client, _ := statsd.New(cfg.StatsdHost)
	return client
}
