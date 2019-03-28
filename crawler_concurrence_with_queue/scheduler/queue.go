package scheduler

import (
	"go_study/crawler_concurrence_with_queue/engine"
)

type QueueScheduler struct {
	RequestQueue chan engine.Request
	WorkerChan   chan chan engine.Request
}

func (qs *QueueScheduler) WorkerReady(work chan engine.Request) {
	qs.WorkerChan <- work
}

func (qs *QueueScheduler) Submit(r engine.Request) {
	qs.RequestQueue <- r
}

func (qs *QueueScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (qs *QueueScheduler) Run() {
	qs.WorkerChan = make(chan chan engine.Request)
	qs.RequestQueue = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case request := <-qs.RequestQueue:
				requestQ = append(requestQ, request)
			case work := <-qs.WorkerChan:
				workerQ = append(workerQ, work)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]

			}
		}
	}()
}
