package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)

	e.Scheduler.Run()

	// 根据指定的数量创建 worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//将任务源 request 送入调度器中
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		// 通过结果 chan 中取出数据
		result := <-out
		// 处理结果
		for _, item := range result.Items {
			fmt.Printf("ConcurrentEngine got item : %v\n", item)
		}

		// 将新任务源 request 加入到调度器中处理
		for _, requestSource := range result.Requests {
			e.Scheduler.Submit(requestSource)
		}
	}
}

/*
	开启一个 goroutine 去执行 worker 任务
*/
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler  i'm ready
			ready.WorkerReady(in)

			// 从 chan 中取出任务
			request := <-in
			// 交给 worker 去执行解析
			result, e := worker(request)
			if e != nil {
				continue
			}
			// 将结果发送到解析结果 chan 中
			out <- result

		}
	}()
}
