package circuitbreaker

import (
	"time"
)

func (c *CircuitBreaker) trafficControl() {
	cbConf := c.config.CircuitBreaker
	for key, data := range c.circuitBreakerData.Instances() {
		timeToOpen := time.Now().Add(cbConf.Durations.ClosedDuration)
		if data.Traffic.TrafficCount() > 0 && data.Traffic.IsOnFullyOpen() && data.Stats.ErrorRateInPercent() > cbConf.Thresholds.MaxErrorPercent {
			data.Traffic.CloseUntil(timeToOpen)
		} else
		if data.Traffic.IsOnClosed() && time.Now().After(data.Traffic.OpenTime()) {
			data.Traffic.HalfOpen()
		} else
		if data.Traffic.IsOnHalfOpen() && data.Stats.SuccessRateInPercent() > cbConf.Thresholds.MinSuccessPercent {
			data.Traffic.FullyOpen()
		} else
		if data.Traffic.TrafficCount() > 0 && data.Traffic.IsOnHalfOpen() && data.Stats.ErrorRateInPercent() > cbConf.Thresholds.MaxErrorPercent {
			data.Traffic.CloseUntil(timeToOpen)
		}
		c.circuitBreakerData.ResetStats(key)
	}
}
