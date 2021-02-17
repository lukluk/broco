package scheme

type Stats struct {
	errorCount float32
}

func (m *Stats) IncErrorCount() {
	m.errorCount ++
}

func (m *Stats) ErrorCount() float32 {
	return m.errorCount
}


func (m *Stats) ErrorRateInPercent(trafficCount int) int {
	if trafficCount == 0 {
		return 0
	}
	return int((m.errorCount / float32(trafficCount)) * 100)
}

func (m *Stats) Reset() {
	m.errorCount = 0
}