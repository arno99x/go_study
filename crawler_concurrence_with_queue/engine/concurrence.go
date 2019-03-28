package engine

import "fmt"

type ConcurrenceEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (ce *ConcurrenceEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	ce.Scheduler.Run()
	for i := 0; i < ce.WorkCount; i++ {
		CreateWorker(out, ce.Scheduler)
	}

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}

	count := 0
	for {
		result := <-out
		for _, item := range result.Items {
			count++
			fmt.Printf("%d item : %s\n", count, item)
		}

		for _, request := range result.Requests {
			ce.Scheduler.Submit(request)
		}
	}
}

func CreateWorker(out chan ParserResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			fmt.Println("work on url : ", request.Url)
			result, err := Worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
