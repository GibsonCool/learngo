package fib

/*
	用闭包实现 斐波拉契
*/
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
