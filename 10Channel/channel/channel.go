package main

import (
	"fmt"
	"time"
)

/*
	Go 中的并发模型理论，取自于 CSP "Communicating Sequential Processes" 衍生出来的

	Go 语言的创作者对于并发所说的的一句话，也是 CSP 在 go 中的体现：
		"Don't communicate by sharing memory; share memory by communicating"
		"不要通过共享内存来通信；而是通过通信来共享内存"

		内存对应了 goroutine 中协程实体，通信对应了 channel 通道

*/
func main() {
	//chanDemo()
	//chanDemo2()
	//bufferChannel()
	channelClose()
}

func chanDemo() {
	//定义一个传输 int 类型的 channel （管道）
	//var c chan int 		// c == nil

	c := make(chan int)
	go worker(0, c)

	/*
		使用  '<-' 来表示 channel 的数据流向

			向 channel 加入数据：   c <- data
			从 channel 取出数据：	  data <- c

		还可以用当 chan 作为参数或者方返回值的时候，也可以用 '<-' 规定  chan 是用作 "发送" 还是 "取出" 数据

			规定 channel 用作发送数据(send-only)： 	 chan <-  type
			规定 channel 用过取出数据(received-only)：  <- chan  type

	*/
	c <- 1
	c <- 2

}

func chanDemo2() {

	var channels [10]chan<- int
	// 创建 10 个 channel
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	// 遍历向每个 channel 中加入数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	for {
		//n := <-c
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

/*
	chan 也可以作为返回值 ,
	并且可以修饰 chan 让调用者知道 chan 是用于 send-only 还是 received-only
*/
func createWorker(id int) chan<- int {
	c := make(chan int)
	//go func() {
	//	for {
	//		fmt.Printf("Create Worker %d received %c\n", id, <-c)
	//	}
	//}()

	go worker(id, c)

	return c
}

/*
	必要的场景下通过 bufferChannel 可以提高效率
*/
func bufferChannel() {
	// 通过内函数 make 创建带有缓冲的 chan
	c := make(chan int, 3)
	go worker(66, c)
	c <- 'a'
	c <- 'b'
	c <- 'c' //因为有缓冲，在没有 goroutine 的时候，如果没有达到缓冲数量是不会出现 deadlock! 的 error
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go workerCanClose(88, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func workerCanClose(id int, c chan int) {
	// 方式一：判断当 channel close 以后，取出来不成功就停止
	//for {
	//	n,ok := <-c
	//	fmt.Println("n:",n," ok:",ok)
	//	if !ok{
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}

	// 方式二：使用 range 接收。当 channel close 以后。取不出 n 循环结束也被动停止
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}
