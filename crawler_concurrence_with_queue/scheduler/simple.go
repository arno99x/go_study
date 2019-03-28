package scheduler

import "go_study/crawler_concurrence_with_queue/engine"

type SimpleSchedule struct {
	workChan chan engine.Request
}

func (s *SimpleSchedule) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	//每提交一个request都开启一个goroutine来处理，为了避免chan间的相互等待，开足火力提升效率
	//但缺点是无法控制，分发出去到goroutine的request就收不回来了，不知道发出去后的任务状态，也没办法做负载均衡类的控制
	go func() {
		s.workChan <- request
	}()

}
