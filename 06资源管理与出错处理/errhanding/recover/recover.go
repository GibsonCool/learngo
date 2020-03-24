package main

import (
	"fmt"
)

/*
	panic:    		//类似于java throw Exception 抛异常

		1.停止当前函数执行

		2.一直向上返回，执行每一层 defer  (前面说过 defer 是在程序结束时候按照栈先进后出一个一个执行，异常了也是程序结束所以会执行，不难理解)

		3.如果没有遇到 recover 。程序退出



	recover:

		1.仅在 defer 调用中使用

		2.获取 panic 的值

		3.如果无法处理，可重新 panic   // 有点类似 Java 中 try()cathch() 后在 catch() 中继续 throw exception

*/
func main() {
	tryRecover()
}

func tryRecover() {

	// defer 中跟一个匿名函数处理，并且调用
	defer func() {
		r := recover()
		//func recover() interface{} recover返回任意类型所以进行类型转换看是否是 error
		if err, ok := r.(error); ok {
			fmt.Println("Error occurent: ", err)
		} else {
			//如果不是，从新panic
			panic(fmt.Sprintf("不知道怎么处理：%v", r))
		}
	}()

	/*
		error
	*/
	//panic(errors.New("自定义一个错误"))
	b := 0
	a := 5 / b
	fmt.Println(a)

	/*
		not error
	*/
	//panic(123)

}
