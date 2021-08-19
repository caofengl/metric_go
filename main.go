package main

import (
	"time"

	"github.com/caofengl/metric_go/metric"
	"github.com/caofengl/metric_go/user"
)

func main() {

	metricIns := metric.Metric{}
	metricIns.Init()

	user := user.UserController{}
	user.Init(&metricIns)

	ticker := time.NewTicker(1 * time.Millisecond)

	for range ticker.C {
		user.Register()
		user.Login()
	}
}
