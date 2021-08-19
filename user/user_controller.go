package user

import "math/rand"

type MetricI interface {
	RecordResponseTime(name string, responseTimes float32)
	RecordTimestamp(name string, timestamps uint32)
	StartRepeateReport()
}

type UserController struct {
	metrics MetricI
}

func (u *UserController) Init(metrics MetricI) {
	u.metrics = metrics
	go u.metrics.StartRepeateReport()
}

func (u *UserController) Register() {
	u.metrics.RecordResponseTime("register", float32(rand.Intn(100)))
}

func (u *UserController) Login() {
	u.metrics.RecordResponseTime("login", float32(rand.Intn(100)))
}
