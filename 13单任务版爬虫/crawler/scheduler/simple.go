package scheduler

import "imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"

// 任务调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	//panic("implement me")
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//开启 goroutine 并发分发request 避免 in chan 和 out chan 循环等待
	go func() {
		s.workerChan <- r
	}()
}
