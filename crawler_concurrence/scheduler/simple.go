package scheduler

import "go_study/crawler_concurrence/engine"

type SimpleSchedule struct {
	workChan chan engine.Request
}

func (s *SimpleSchedule) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	s.workChan <- request
}
