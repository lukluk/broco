package scheme

type Stats struct {
	successCount float32
	errorCount float32
}

func (m *Stats) IncSuccessCount() {
	m.successCount ++
}

func (m *Stats) IncErrorCount() {
	m.errorCount ++
}

func (m *Stats) SuccessCount() float32 {
	return m.successCount
}

func (m *Stats) ErrorCount() float32 {
	return m.errorCount
}

func (m *Stats) SuccessRateInPercent() int {
	if m.successCount + m.errorCount == 0 {
		return 0
	}
	return int((m.successCount / (m.successCount + m.errorCount)) * 100)
}

func (m *Stats) ErrorRateInPercent() int {
	if m.successCount + m.errorCount == 0 {
		return 0
	}
	return int((m.errorCount / (m.successCount + m.errorCount)) * 100)
}

func (m *Stats) Reset() {
	m.errorCount = 0
	m.successCount = 0
}