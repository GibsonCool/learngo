package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))

	case score < 60:
		g = "F"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func main() {
	const filename = "01基础语法/abc.txt"

	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//可以通过if的条件赋值。简写表达式。  其中赋值的变量作用域值就只在这个if语句里
	if contents2, err2 := ioutil.ReadFile(filename); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s\n", contents2)
	}

	//无法使用contents2变量
	//fmt.Printf("%s\n", contents2)

	fmt.Println(
		grade(55),
		grade(86),
		grade(99),
		//grade(101),
	)
}
