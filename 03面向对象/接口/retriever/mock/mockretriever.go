package mock

import "fmt"

type Retriever struct {
	Contents string
}

/*
	接口的实现是隐式的，不需要像 Java 一样实现接口

	只需要实现接口中的方法
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
