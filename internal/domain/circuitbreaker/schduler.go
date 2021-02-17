package circuitbreaker

import (
	"github.com/jasonlvhit/gocron"
)

func (c *CircuitBreaker) RunScheduler() {
	_ = gocron.Every(uint64(c.config.CircuitBreaker.Durations.StateInterval.Seconds())).Seconds().
		Do(c.trafficControl)
	<- gocron.Start()
}

