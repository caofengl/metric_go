package metric

import (
	"fmt"
	"sync"
	"time"
)

type Metric struct {
	responseTimes map[string][]float32
	timestamps    map[string][]uint32
	lock          sync.Mutex
}

func (m *Metric) Init() {
	m.responseTimes = make(map[string][]float32)
	m.timestamps = make(map[string][]uint32)
	m.lock = sync.Mutex{}
}

func (m *Metric) RecordResponseTime(key string, responseTime float32) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.responseTimes[key] = append(m.responseTimes[key], responseTime)
}

func (m *Metric) RecordTimestamp(key string, timestamp uint32) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.timestamps[key] = append(m.timestamps[key], timestamp)
}

func (m *Metric) StartRepeateReport() {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		status := make(map[string]map[string]float32)

		m.lock.Lock()
		for name, times := range m.responseTimes {
			if status[name] == nil {
				status[name] = map[string]float32{}
			}
			status[name]["max"] = m.max(times)
			status[name]["avg"] = m.avg(times)
		}
		m.lock.Unlock()

		m.lock.Lock()
		for name, times := range m.timestamps {
			if status[name] == nil {
				status[name] = map[string]float32{}
			}
			status[name]["count"] = float32(len(times))
		}
		m.lock.Unlock()

		fmt.Println(status)
	}
}

func (m *Metric) max(responseTimes []float32) float32 {
	var maxTime float32
	for _, v := range responseTimes {
		if v > maxTime {
			maxTime = v
		}
	}
	return maxTime
}

func (m *Metric) avg(responseTimes []float32) float32 {
	var totalTime float32
	for _, v := range responseTimes {
		totalTime += v
	}
	return totalTime / float32(len(responseTimes))
}
