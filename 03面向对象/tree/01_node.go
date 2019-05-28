package tree

import "fmt"

/*
	面向'对象'：

		go 语言中只支持封装。不支持继承和多态(继承和多态功能过接口来做，面向接口编程)

		go 中也没有class（类）,只有struct--》结构体
*/

/*
	结构体的定义 :
		type  name struct{
			value1  valueType
			value2  valueType
			...
		}

	定义只需要定义好属性，成员就可以了，不需要所谓的'构造函数'，因为从后面的代码可以看出有很多创建结构体的方式。不需要构造函数


	结构体方法的定义：

		1、值接受者
			func (struetName StructType) funcName (Value ValueType){
				...
			}

		2、指针接受者
			func (struetName *StructType) funcName (Value ValueType){
				...
			}

	值接受者 VS 指针接受者：

		* 要改变内容必须使用指针接受者

		* 结构过大也考虑使用指针接受者，毕竟值接受者是要拷贝原数据过大影响性能和占用内存

		* 一致性：如有指针接受者，最好都是指针接受者，这只是一个好规范习惯。毕竟值接受者能做的指针接受者都能做
*/
type Node struct {
	Value       int
	Left, Right *Node
}

/*
	如果有特殊需求，可以自定义工厂函数，创建结构体
	注意返回了局部变量的地址
*/
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

/*
	结构体的方法定义:

		与 Java 类比不通的是，就是个方法的定义，不需要写在结构体内

		「显示定义和命名方法接受者」

		支持值接受者 和 指针接受者

		nil 指针也可以调用方法
*/
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//与上面相比，其实结构的方法和普通方法大同小异，只是一个需要方法接受者，然后调用方式不同
func Print(node Node) {
	fmt.Print(node.Value, " ")
}

//值传递--》数据会copy一份，对copy数据的操作并不会反应到原数据上
func (node Node) SetValue(value int) {
	node.Value = value
}

//指正传递(引用传递)--》传递是元数据的地址值。对数据的操作会直接反应到原数据上
func (node *Node) SetValuePointer(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node ,Ignore")
		return
	}
	node.Value = value
}

//遍历这个树节点  左->中->右
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
