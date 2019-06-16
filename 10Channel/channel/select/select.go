package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	select:

		Go 中的一个控制结构，类似于 switch 语句,不过是用于通信的,每个 case 必须是一个通信操作，要么是发送（send）要么是接收（receive）

		select 会监听 case 语句中 channel 的读写操作，当 case 中 channel 读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

		所有 channel 表达式都会被求值
		所有被发送的表达式都会被求值
		select 中的 default 子句总是可运行的。
		如果有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行。
		如果没有可运行的 case 语句，且有 default 语句，那么就会执行 default 的动作。
		如果没有可运行的 case 语句，且没有 default 语句，select 将阻塞，直到某个 case 通信可以运行

*/
func main() {
	//randSelect()
	//test1()
	test2()
}

/*
	同一个 channel 发送数据，两个不同的 case 从同一个 channel 取数据
	select 从输出结果可以看出是随机选择一个case执行其中表达式
*/
func randSelect() {

	c := make(chan int)
	go func() {
		for range time.Tick(time.Second) {
			c <- 0
		}
	}()

	for {
		select {
		case <-c:
			fmt.Println("case 1")
		case <-c:
			fmt.Println("case 2")
		}
	}
}

/*
	select 的使用
	nil chan 在 select 中的使用
	模仿一个生产消费者模型
	存在问题：
		生产速度大于消费速度时，中间的数据遗漏
*/
func test1() {
	var c1, c2 = generator(), generator()
	n := 0
	w := createWorker("test1")
	hasValue := false
	for {
		var activeWork chan<- int

		// 如果 c1 或者 c2 中有值发送过来了，则 activeWork 应该赋值
		if hasValue {
			activeWork = w
		}
		select {
		case n = <-c1:
			//fmt.Println("Received from c1:", n)
			//w <- n
			hasValue = true
		case n = <-c2:
			//fmt.Println("Received from c2:", n)
			//w <- n
			hasValue = true
		case activeWork <- n: // nil chan 在 select 中是可以正确书写执行的，借助这一特性，我们在有值的时候在给 activeWork 赋值然后可被 select 执行将值传送过去。nil的时候就不会被 select 执行
			hasValue = false
		}
	}
}

/*
	借助 slice 来记录数据，解决 test1 中遗留的问题

	time 计时器在 select 中的使用
*/
func test2() {
	var c1, c2 = generator(), generator()
	n := 0
	w := createWorker("托尔斯泰")

	// 用 slice 来保存所有生产者生成的数据。相当于中间缓冲，避免生产大于消费导致数据丢失
	var values []int

	// 倒计时使用，返回一个 chan time 类型。当到达指定时间后，回放 chan time 中发送当前 time 数据
	timeEnd := time.After(10 * time.Second)

	// 定时任务，定时每 1 秒钟往 返回的 chan time 中发送数据
	tick := time.Tick(time.Second)

	fmt.Println("开始时间：", time.Now())
	for {
		var activeWrok chan<- int
		var activeValue int

		// 如果队列中有值，则赋值 activeWork 让其有机会去被 select 选择执行取值任务。
		if len(values) > 0 {
			activeWrok = w
			// 取出队列中的首个数据
			activeValue = values[0]
		}

		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWrok <- activeValue:
			// 当执行这里，说明数据被消费了，从队列移除这个数据
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			// 由于 time.After(）会返回一个 chan 。这里就成了 <- chan time  表达式。
			// 在 select 语句的语法中，所有 channel 表达式都会被求值
			// 如果规定时间后，该 chan 返回值可执行，c1,c2 还没有被产生值达到可执行条件，
			// select 语句就选择可进入该 case . 视为响应超时
			fmt.Println("time out 连接超时")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case time := <-timeEnd: //当 timeEnd 被 send 了数据后，这个 case 就可以被 select 选择执行
			fmt.Println("时间结束：", time)
			return
		}
	}
}

func generator() chan int {
	out := make(chan int)
	go func() {
		for i := 0; ; i++ {
			// 不等长停顿时间，随机在 0-1.5 秒之间
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
		}
	}()
	return out
}

func worker(id string, c chan int) {
	for n := range c {
		//每停顿 3 秒在去收下个数据。模拟消耗数据的速度远小于生产数据的速度
		time.Sleep(time.Second)
		fmt.Printf("Worker %s received value : %d\n", id, n)
	}
}

func createWorker(id string) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
