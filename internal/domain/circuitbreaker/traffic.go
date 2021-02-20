package circuitbreaker

import (
	"time"
)

func (c *CircuitBreaker) trafficControl() {
	cbConf := c.config.CircuitBreaker
	for key, instance := range c.circuitBreakerData.Instances() {
		timeToOpen := time.Now().Add(cbConf.Durations.ClosedDuration)
		errRate := instance.Stats.ErrorRateInPercent(instance.Traffic.TrafficCount())
		if instance.Traffic.TrafficCount() > 0 && instance.Traffic.IsOnFullyOpen() && errRate > cbConf.Thresholds.MaxErrorPercent {
			c.statsdClient.Incr(TrafficClosedMetric, []string{key}, 1)
			instance.Traffic.CloseUntil(timeToOpen)
		} else
		if instance.Traffic.IsOnClosed() && time.Now().After(instance.Traffic.OpenTime()) {
			c.statsdClient.Incr(TrafficHalfOpenMetric, []string{key}, 1)
			instance.Traffic.HalfOpen()
		} else
		if instance.Traffic.IsOnHalfOpen() && errRate < cbConf.Thresholds.MinErrorPercent {
			c.statsdClient.Incr(TrafficFullyOpenMetric, []string{key}, 1)
			instance.Traffic.FullyOpen()
		} else
		if instance.Traffic.TrafficCount() > 0 && instance.Traffic.IsOnHalfOpen() && errRate > cbConf.Thresholds.MaxErrorPercent {
			c.statsdClient.Incr(TrafficReturnClosedMetric, []string{key}, 1)
			instance.Traffic.CloseUntil(timeToOpen)
		}
		c.circuitBreakerData.ResetStats(key)
	}
}
