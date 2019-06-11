package queue

import "fmt"

/*
	go 语言写的示例代码，并会对示例输入结果检测，
	文末注释的内容在 godoc -http 6060 后，在浏览器访问可以看到示例注释文档内容
*/
func ExampleQueuePlus_Pop() {
	q := QueuePlus{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	//1
	//2
	//false
	//3
	//true
}
