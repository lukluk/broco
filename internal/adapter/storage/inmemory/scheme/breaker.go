package scheme

import "time"

type Traffic struct {
	closeUntil time.Time
	halfOpen bool
	closed bool
	trafficCount int
}

func (b *Traffic) CloseUntil(time time.Time) {
	b.closed = true
	b.closeUntil = time
}

func (b *Traffic) HalfOpen() {
	b.closed = false
	b.halfOpen = true
}

func (b *Traffic) FullyOpen() {
	b.closed = false
	b.halfOpen = false
}

func (b*Traffic) OpenTime() time.Time {
	return b.closeUntil
}

func (b *Traffic) IsOnClosed() bool {
	return b.closed
}

func (b *Traffic) IsOnHalfOpen() bool {
	return b.halfOpen
}

func (b *Traffic) IsOnFullyOpen() bool {
	return !b.closed
}

func (b *Traffic) Check() bool {
	if b.closed {
		return false
	} else
	if b.halfOpen {
		return (b.trafficCount % 2) > 0
	}
	return true
}

func (b *Traffic) TrafficCount() int {
	return b.trafficCount
}

func (b *Traffic) IncTrafficCount() {
	b.trafficCount ++
}

func (b *Traffic) ResetTrafficCount() {
	b.trafficCount = 0
}

