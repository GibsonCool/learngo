package main

import (
	"fmt"
	"sync"
)

func main() {
	chanDemo()
}

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, w worker) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-w.in)
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//借助函数式编程，将减少 goroutine 数量的操作抽离到这里函数中统一实现
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)

	return w
}

/*
	WaitGroup:
		在 go 语言中，用于线程同步，单从字面意思理解，wait 等待的意思，
		group 组、团队的意思，WaitGroup 就是指等待一组，等待一个系列执行完成后才会继续向下执行。

	用途：
		它能够一直等到所有的 goroutine 执行完成，并且阻塞主线程的执行，直到所有的 goroutine 执行完成。


	WaitGroup总共有三个方法：

		Add:添加或者减少等待 goroutine 的数量

		Done:内部调用的 wg.Add(-1)

		Wait:执行阻塞，直到所有的 WaitGroup 数量变成0
*/
func chanDemo() {
	wg := sync.WaitGroup{}
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	//直接一次性把 chan 的数量加入等待组
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//或者这种每次操作一个加入一个也可以
		//wg.Add(1)
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	//阻塞主线程执行，等待 wg 中所有的 goroutine 执行完在继续往下执行
	wg.Wait()
}
