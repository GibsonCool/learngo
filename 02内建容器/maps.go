package main

import "fmt"

/*
	Map 的定义：
		map[keyType] valueType {keyValue,valueValue}

		valueType也可以是Map类型，比如：
		map[int] map[int]string{}

	Map 的操作

		创建：make(map[string]int)

		获取元素：m[key]

		key不存在时，获取Value类型的初始值

		用value，ok := m[key] 来判断是否存在key

		用 delete 删除一个key



	Map 的遍历（内部是hashMap，所以可以是无序的)

		使用range遍历key,或者遍历key,value键值对

		不保证遍历顺序，如需要顺序，需手动对key排序

		使用len获取元素个数
*/
func main() {
	m := map[string]string{
		"name":    "DoubleX",
		"course":  "golang",
		"site":    "test",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("map 的遍历======================================")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("map 的取值======================================")
	courseNameTest := m["course"]
	courseName, ok := m["course"]
	fmt.Println(courseNameTest, courseName, ok)

	//如果取值是一个不存在的key
	errorName, ok := m["error"]
	fmt.Println(errorName, ok)
	if errorName, ok := m["error"]; ok {
		fmt.Println(errorName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("map 的删除======================================")
	//使用内建函数 delete
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	//name, ok := m["name"]	//这里不能再用:= 赋值了。会因为上面name,ok变量已经通过:=赋值过了，变量已经存在直接使用就可以
	name, ok = m["name"]
	fmt.Println(name, ok)

}
