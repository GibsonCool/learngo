package main

import (
	"fmt"
	"reflect"
)

/*
	切片(Slice):
			Slice 本身是没有数据的，是对底层数组(array)的一个view（映射）
			因此对 view 上的操作都会真实反映到底层 array 中

*/

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("arr 的类型", reflect.TypeOf(arr))
	fmt.Println("s1 的类型", reflect.TypeOf(s1))
	fmt.Println("After updateSlice(s1)******************************************")
	//这里传入切片类型，就可以达到数组传入指针引用的效果。
	updateSlice(s1)
	updateSlice(s2)
	//从结果可以看出，两次操作的结果都反应到了 arr  中
	fmt.Println(s1, s2, arr)

	//切片可以在结果上重复切片。最终效果都会作用到arr上，就相当于arr建立了多个 view 对view 进行操作
	fmt.Println("After Reslice******************************************")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	s2[1] = 6666
	fmt.Println(arr)

	fmt.Println("Exding Slice ******************************************")

	/*
		slice的实现包含三个东西： ptr(指向slice开头的元素)  len(表示这个长度 s[] 取值不能超过这个长度)  cap(表示ptr到整个底层数组结束的长度)
		slice可以向后扩折，不可以向前扩展，
		s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)

		所以s2在s1的基础上再次切片，并且超出了s1的长度但没有超出cap的长度任然可以取到值
	*/
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:6]
	fmt.Println(arr)
	fmt.Printf("s1=%v  len(s1)=%d   cap(s1)=%d \n", s1, len(s1), cap(s1))
	s2 = s1[3:5]
	//s2 = s1[3:7]  这已经超出了cap的长度会出错
	fmt.Printf("s2=%v  len(s2)=%d   cap(s2)=%d \n", s2, len(s2), cap(s2))

	fmt.Println("add item to  Slice ******************************************")
	fmt.Println(arr) // [0 1 2 3 4 5 6 7]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3 , s4, s5===>", s3, s4, s5) //s3,s4,s5===> [5 6 10] [5 6 10 11] [5 6 10 11 12]
	/*
		1、添加元素如果超过cap，系统会重新分配更大的底层数组，原arr如果无人使用垃圾回收机制会回收

			从arr 结果可以看出来，s4,s5的操作并没有在arr上因为 append的时候已经超出了cap(arr)的长度，
			所以系统会用一个更长的数组来copy一份原数据进行操作。然后切片对于新数组的操作不会映射到旧数组上

		2、由于值传递的关系，必须接受 append 的返回值

	*/
	fmt.Println(arr) //[0 1 2 3 4 5 6 10]

}

func updateSlice(s []int) {
	s[0] = 100
}
