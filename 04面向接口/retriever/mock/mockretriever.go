package mock

import "fmt"

type Retriever struct {
	Contents string
}

/*
	接口的实现是隐式的，不需要像 Java 一样实现接口

	只需要实现接口中的方法

	***
		所以很重要的一点思维概念，go 中的接口实现是由使用者决定我需不需要实现那些接口，需要什么样的功能
		而不是由接口决定你实现类。所以我认为接口应该是越简单越好，粒度越小越好，然后确实遇到需要复杂的
		多个功能的，可以进行多样化组合成新接口类型。  组合优于继承的思想很不错。这里和flutter中的控件
		组合大于继承的思想一样，使用起来很nice
	***
*/
func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

// 实现系统的接口 Stringer 提供类似 Java 的 toString()的功能
func (r *Retriever) String() string {
	return fmt.Sprintf("mock Retriever : {Contents = %s}", r.Contents)
}
