package circuitbreaker

import (
	"time"
)

func (c *CircuitBreaker) trafficControl() {
	cbConf := c.config.CircuitBreaker
	for key, data := range c.circuitBreakerData.Instances() {
		timeToOpen := time.Now().Add(cbConf.Durations.ClosedDuration)
		errRate := data.Stats.ErrorRateInPercent(data.Traffic.TrafficCount())
		if data.Traffic.TrafficCount() > 0 && data.Traffic.IsOnFullyOpen() && errRate > cbConf.Thresholds.MaxErrorPercent {
			c.statsdClient.Incr(TrafficClosed, []string{key}, 1)
			data.Traffic.CloseUntil(timeToOpen)
		} else
		if data.Traffic.IsOnClosed() && time.Now().After(data.Traffic.OpenTime()) {
			c.statsdClient.Incr(TrafficHalfOpen, []string{key}, 1)
			data.Traffic.HalfOpen()
		} else
		if data.Traffic.IsOnHalfOpen() && errRate < cbConf.Thresholds.MinErrorPercent {
			c.statsdClient.Incr(TrafficFullyOpen, []string{key}, 1)
			data.Traffic.FullyOpen()
		} else
		if data.Traffic.TrafficCount() > 0 && data.Traffic.IsOnHalfOpen() && errRate > cbConf.Thresholds.MaxErrorPercent {
			c.statsdClient.Incr(TrafficReturnClosed, []string{key}, 1)
			data.Traffic.CloseUntil(timeToOpen)
		}
		c.circuitBreakerData.ResetStats(key)
	}
}
