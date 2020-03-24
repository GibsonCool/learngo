package main

import "fmt"

func main() {

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}

/*
	<----------- 函数式编程 ------------>

	函数是一等公民：可以作为变量、参数、返回值

	高阶函数：因为函数可以作为变量，那么函数又接收变量，一个函数就可以接受另一个函数当成变量作为参数，也可以将一个函数作为变量值返回

	闭包：
		将 innerAdder 函数（函数体）作为 adder 函数的返回值

		innerAdder 函数中包含局部变量 v  和 自由变量 sum

		innerAdder 函数保存了自由变量 sum 引用

*/
func adder() func(int) int {
	sum := 0
	innerAdder := func(v int) int {
		sum += v
		return sum
	}
	return innerAdder
}
