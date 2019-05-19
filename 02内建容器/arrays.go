package main

import "fmt"

/*
	数组的定义和使用(遍历)
*/
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 3, 4, 3, 6}

	var grid [4][3]int

	fmt.Println(arr1, arr2, arr3)

	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for index, value := range arr3 {
		fmt.Println(index, value)
	}

	sum := 0
	//go中变量如果不适用编译器会提示错误，语法严格，可以使用 '_' 代替去除警告
	for _, v := range arr3 {
		sum += v
	}
	fmt.Printf("sum: %d", sum)

	fmt.Println("********************************")

	printArray(arr1)
	fmt.Println("********************************")

	printArray(arr3)
	fmt.Println("********************************")

	//这里无法调用：函数定义的参数类型是 type [5]int   但是传入的参数类型却是 type [3]int
	//printArray(arr2)

	fmt.Println(arr1, arr2, arr3)
	fmt.Println("********************************")

	printArrayPointer(&arr1)
	//printArrayPointer(&arr2)
	fmt.Println("********************************")

	fmt.Println(arr1)
	fmt.Println("********************************")
}

// 数组是值类型，在函数中使用的时候回copy值
// []int  和 [5]int  是完全不一样的东西，前者是切片  后者才是数组

func printArray(arr [5]int) {
	//在这里将第一个数值改为100 并不会影响到原传入数组中的值，因为数组是值传递，会copy一份
	arr[0] = 100
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

//可以传入指针来达到应用传递的效果
func printArrayPointer(arr *[5]int) {
	//在这里将第一个数值改为100 并不会影响到原传入数组中的值，因为数组是值传递，会copy一份
	arr[0] = 100
	for index, value := range arr {
		fmt.Println(index, value)
	}
}
