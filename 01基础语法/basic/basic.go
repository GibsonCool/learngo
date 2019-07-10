package main

import (
	"fmt"
	"math"
)

/**
变量的定义
*/
func variableZeroValue() {
	var a int
	var s string

	fmt.Printf("a = %d   s = %s\n", a, s)
	fmt.Printf("a = %d   s = %q\n", a, s)
}

func variableInitiaValue() {
	var a, b int = 2, 3
	var s string = "test"
	fmt.Println(a, b, s)
}

//Go中也支持类型推导
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	var (
		d = 0
		f = "ttt"
	)
	fmt.Println(a, b, c, s, d, f)
}

func variableShorter() {
	//函数内变量的定义推荐这种方法、简介明了
	a, b, c, s := 3, 4, true, "def"
	c = false

	fmt.Println(a, b, c, s)
}

var aa = 5
var bb = "stsr"

//定义包内部的变量就不能使用 := 的方式,Go 中没有全局变量的概念
//cc := 33
var cc = 33

//简写方式
var (
	dd = 99
	ee = "def"
)

//Go类型中没有隐式类型转换，必须都使用强制类型转换
func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//常量定义 使用关键字 const
//Go中常量不全部大写，常量数值可作为各种类型使用
const filename = "abc.txt"
const (
	d = 5
	e = 9
)

func consts() {
	const a, b = 3, 4
	var c int
	//这里使用常量之后不需要强制转换为float ，因为使用常量并且没有固定类型，这里计算将相当于文本替换然后在类型推导
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c, d, e)
}

//Go中的枚举类型也是使用const 来定义
func enums() {
	const (
		red    = "red"
		yellow = "yellow"
		blue   = "blue"
		green  = "green"
	)

	//iota  -->const iota = 0 // Untyped int.
	//Go中通过iota实现自增长的枚举，iota会在下一行增长，
	const (
		cpp = iota
		java
		python
		golang
		kotlin
	)

	const (
		Apple, Banana = iota + 1, iota + 2
		Cherimoya, Durian
		Elderberry, Fig
	)

	//通过自增至iota创建表达式自增
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(red, yellow, blue, green)
	fmt.Println(cpp, java, python, golang, kotlin)
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, dd, ee)
	triangle()

	consts()

	enums()
}
