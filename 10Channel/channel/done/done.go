package main

import (
	"fmt"
)

/*
	使用 channel 来双向通信等待任务结束，通过通信来共享内存的实例
	channel 默认为同步模式，需要'发送'和'接受'配对，否则会被阻塞，直到另一方准备好后被唤醒
*/

func main() {
	chanDemo()
}

type worker struct {
	in   chan int
	done chan bool
}

func doWorker(id int, c chan int, done chan bool) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
		done <- true
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)

	return w
}

/*
	从输出结果看，程序并没有异常除以 deadlock 死锁的情况，

	因为无论是 in chan int  还是 done chan bool 都是发送和接收配对。

	但是这里必须 in 发送接收完。done 就跟着发送接收通知结束继续下一个成了顺序执行了，并发的意义全无

	为了解决这种问题，我们使用系统的 WaitGroup 来避免。查看 waitGroup --> done2
*/
func chanDemo() {

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		<-worker.done
	}

}
