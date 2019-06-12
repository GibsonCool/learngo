package main

import (
	"fmt"
	"time"
)

/*
	协程 Coroutine，go 中的 Goroutine 是实际并发执行的实体，底层是使用协程实现并发

		轻量级"线程"

		非抢占式多任务处理，有协程主动交出控制权

		编译器/解释器/虚拟机层面的多任务

		多个协程可以在一个或者多个线程上运行（由调度器决定）
*/
func main() {
	var a [10]int

	//runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		/*
			任何函数只需要加上 go 就能送给调度器运行，开启协程。调度器会在合适的点进行切换

			goroutinue 可能的切换点：（只是参考，不能保证切换，不保证在其他地方不切换）
				I/O,select    channel   等待锁   函数调用（有时）  runtime.Gosched()
		*/

		go func(index int) {
			for {
				// Print 是一个 io 操作，io 操作会有等待的过程这里会自动切换，交出控制权
				//fmt.Println("hello from groutine ",index)

				// todo: 这里本机运行的结果和视频中不一样，并没有一直死循环卡着，还是正常切换结束了
				// 如果我们进行的是一个不主动交出控制权的操作，后面的 协程就拿不到执行权，一直卡在这个协程中
				a[index]++
				// 主动结束交出执行权
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("end:", a)
}
