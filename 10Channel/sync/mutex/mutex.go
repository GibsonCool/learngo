package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	go 中也有支持传统的同步机制（"通过共享内存来通讯"）

	其中典型的一个就是 Mutex 互斥锁。
*/
func main() {
	goRoutinuePrint(Incr)
	//goRoutinuePrint(IncrMutex)
}

/*
	开两个协程同时去调用带有 a++ 操作的函数 Incr
	有几率触发两次打印结果一致，因为没有做互斥操作
*/
func goRoutinuePrint(doSomething func() int) {
	go func() {
		fmt.Println(doSomething())
	}()
	go func() {
		fmt.Println(doSomething())
	}()
	time.Sleep(time.Duration(1) * time.Second)
}

var a = 0

func Incr() int {
	a++
	time.Sleep(time.Duration(1))
	return a
}

var lock = &sync.Mutex{}

func IncrMutex() int {
	// 进行操作的时候
	lock.Lock()
	defer lock.Unlock()
	a++
	time.Sleep(time.Duration(1))
	return a
}
