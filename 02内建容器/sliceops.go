package main

import (
	"fmt"
)

/*

 */
func main() {
	fmt.Println("创建 slice ============================================")
	/*
		slice的创建方式
	*/
	//值定义变量无初始值，此时为nil
	var s []int //Zero value for slice is nil

	// 有初始值确定  len = cap
	s1 := []int{2, 4, 6, 8}

	//借助内建函数 make   len = cap
	s2 := make([]int, 16)

	// len != cap
	s3 := make([]int, 10, 32)

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i*2+1)
	}
	fmt.Println(s)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("拷贝 slice ============================================")
	//内建函数  copy
	printSlice(s2)
	//会将s1中的值拷贝到s2中
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("删除 slice 中的元素 8  ============================================")
	//利用切片和append
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("popping from front 从队列头部取出")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("popping from back  从队列尾部取出")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("%v  len=%d, cap=%d\n", s, len(s), cap(s))
}
