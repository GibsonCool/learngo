package main

import (
	"fmt"
	"imooc.com/doublex/learngo/03面向对象/tree"
)

func main() {
	var root tree.Node
	fmt.Println(root)

	//结构体实例的创建
	root = tree.Node{}
	fmt.Println(root)

	//指针引用，所以加&取地址值
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	fmt.Println(root)

	//还可以通过内建函数 new(structName) 来创建结构体
	// * 不论地址还是结构本身，一律可以使用 . 来访问成员 *
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)

	// 可以看到结构体方法和普通方法的效果一样，调用方式有差而已。
	// 这个跟 kotlin 的扩展函数就相当于对 java 的对象进行了一个方法包装传入该对象本身作为参数类似
	root.Print()
	tree.Print(root)

	//调用值传递的方法无法修改数据
	root.SetValue(10)
	root.Print()

	//调用指针传递的方法可以修改数据。并且和值传递方法调用无差别
	root.SetValuePointer(33)
	root.Print()

	//nil指针也可调用方法，但编辑器一般会有黄线警告
	var testNil *tree.Node
	testNil.SetValuePointer(3)
	testNil = &root
	testNil.SetValuePointer(300)
	testNil.Print()

	fmt.Println()
	root.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	fmt.Println("================================")
	q := Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.isEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.isEmpty())

	fmt.Println("================================")
	c := root.TraverseWithChannel()
	maxValueNode := 0
	for node := range c {
		if node.Value > maxValueNode {
			maxValueNode = node.Value
		}
	}
	fmt.Printf("maxValue : %d\n", maxValueNode)

}

/*
	扩展已有类型：
		定义别名

		使用组合
*/

//通过组合扩展
type myTreeNode struct {
	node *tree.Node
}

//扩展遍历数节点的方法，遍历顺序改为  左->右->中
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	myNode.node.Left.Traverse()

	myNode.node.Right.Traverse()

	myNode.node.Print()
}

//通过定义别名  对slice的操作进行包装扩展为一个队列操作
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}
