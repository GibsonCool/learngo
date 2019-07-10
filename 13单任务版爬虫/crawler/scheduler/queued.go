package scheduler

import "imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"

/*
	request 队列和 worker 队列 搭配select实现调度器
*/
type QueueScheduler struct {
	// 请求任务
	requestChan chan engine.Request
	// worker 任务接收
	workerChan chan chan engine.Request
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (*QueueScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

//总控制调度,使用 select 来调度控制多个 chan 通信
func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			// 只有当两个队列都有值的时候才赋予其有抢到执行的权力
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
