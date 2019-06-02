package main

import (
	"bufio"
	"fmt"
	"imooc.com/doublex/learngo/03面向对象/tree"
	"io"
	"strings"
)

func main() {

	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println("========================")
	f2 := fibonacci2()
	printFileContent(f2)

	fmt.Println("========================")
	//使用函数来遍历二叉树
	root := tree.Node{
		Left: &tree.Node{
			Value: 2,
			Right: tree.CreateTreeNode(3),
		},

		Value: 6,

		Right: &tree.Node{
			Left:  tree.CreateTreeNode(7),
			Value: 9,
			Right: tree.CreateTreeNode(11),
		},
	}

	root.Traverse()
	fmt.Println()
	root.Print()
	fmt.Println()

	//通过传入函数统计二叉树的量
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}

/*
	用闭包实现 斐波拉契
*/
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/*
	go  中的一个特点
	函数也能实现接口
*/
func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 给需要返回的函数定义为类型 intGen
type intGen func() int

// 然后让 intGen 实现 Reader 功能的接口支持通过Read来读取值，并作为 fibonacci2 的返回值
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	//防止无限调用读取
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//调用已经实现过接口的类来读取我们的返回值，简化操作
	return strings.NewReader(s).Read(p)
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	//省略初始条件,递增条件实现while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
