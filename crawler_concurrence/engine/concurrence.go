package engine

import "fmt"

type ConcurrenceEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (ce *ConcurrenceEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	ce.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < ce.WorkCount; i++ {
		CreateWorker(in, out)
	}

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("item : %s\n", item)
		}

		for _, request := range result.Requests {
			ce.Scheduler.Submit(request)
		}
	}
}

func CreateWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
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
