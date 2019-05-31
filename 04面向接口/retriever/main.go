package main

import (
	"fmt"
	mock2 "imooc.com/doublex/learngo/04面向接口/retriever/mock"
	queue3 "imooc.com/doublex/learngo/04面向接口/retriever/queue"
	real3 "imooc.com/doublex/learngo/04面向接口/retriever/real"
	"time"
)

const url = "http://www.imooc.com"

//接口的定义
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

//接口的使用
func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "DoubleX",
		"course": "golang",
	})
}

/*
	接口的组合，可以使用任何接口组合为新接口
*/
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(rp RetrieverPoster) string {
	result := rp.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return rp.Get(url) + "   " + result
}

/*
	接口变量里面有什么？

		①实现者的类型   ②实现者的值（可以是值copy的实现者的一份。也可以是指针也就是地址值指向实现者）

	查看接口变量

		①Type Assertion   ②Type Switch

*/
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch: ")

	switch v := r.(type) {
	case *mock2.Retriever:
		fmt.Println("Contents:", v.Contents)

	case *real3.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {

	var r Retriever
	mockRetriever := mock2.Retriever{Contents: "this is a fake imooc.com"}
	r = &mockRetriever
	fmt.Println(download(r))
	//fmt.Printf("%T %v\n", r, r) //	mock.Retriever {this is a fake imooc.com}
	inspect(r)

	r = &real3.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Printf("%T %v\n", r, r) //	*real.Retriever &{Mozilla/5.0 1m0s}
	//fmt.Println(download(r))
	inspect(r)

	// Type assertion  直接取得接口中的真实类型。判断实现者是否是 mock.Retriever 类型
	if realRetriever, ok := r.(*mock2.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Printf("not a mock retriever. This is a %T\n\n", realRetriever)
	}

	// 通过 interface{} 可以让queue接收任何类型的对象值
	queue := queue3.QueuePlus{}
	queue.Push(1)
	queue.Push("abc")
	queue.Push(true)
	queue.Push(nil)
	fmt.Printf("Is Empty: %t    queue length:%d\n", queue.IsEmpty(), len(queue))
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	//fmt.Println(queue.Pop())
	fmt.Println()
	fmt.Println("------try a session 接口组合的实现使用-------")
	fmt.Println(session(&mockRetriever))
}
