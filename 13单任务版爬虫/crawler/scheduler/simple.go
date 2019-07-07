package scheduler

import "imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"

// 任务调度器
type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkerChan <- r
	}()
}
