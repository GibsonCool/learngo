package main

import (
	"fmt"
	"math"
)

/*
	函数定义：
		支持返回多个值，并且可以对值起名字
*/

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)

	}
}

func evalPlus(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)

	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func divName(a, b int) (q, r int) {
	return a / b, a % b
}

//支持函数式编程。函数是一等公民，函数也可以作为参数使用
func apply(op func(float64, float64) float64, a, b float64) float64 {
	return op(a, b)
}

//支持可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

//go中指针的概念  函数调用的时候可以有  值传递  和 引用传递(内存地址值传递）

//这种值传递无法交换两个数的值
func swap(a, b int) {
	b, a = a, b
}

//值传递可以通过这种方法返回接受来叫从新赋值交换
func swapReturn(a, b int) (bb, aa int) {
	return b, a
}

//或者直接传递指针的方式直接操作复制
func swapPoin(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println(eval(2, 3, "*"))

	i, i2 := div(5, 3)
	fmt.Println(i, i2)

	q, r := divName(5, 4)
	fmt.Println(q, r)

	fmt.Println(evalPlus(3, 4, "$"))

	fmt.Println(apply(math.Max, 10, 3))
	//或者直接在调用的时候写入一个匿名函数
	fmt.Println(apply(
		func(f float64, f2 float64) float64 {
			return math.Min(f, f2)
		}, 3, 5))
	fmt.Println(sum(3, 5, 6, 2, 32))

	a, b := 3, 5
	swap(a, b)
	fmt.Println(a, b)

	a, b = swapReturn(a, b)
	fmt.Println(a, b)

	a, b = 11, 22
	//通过指针调用
	swapPoin(&a, &b)
	fmt.Println(a, b)
}
