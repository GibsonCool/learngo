package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	rune 在 go 中就是其他语言的 char 类型

	字符串的一些操作通常都可以在 strings 这个包里面找到，常用的比如:

		修改操作：Fields,Split,Join
		统计操作：Contains,Index
		大小写：ToLower,ToUpper
		去空格：Trim,TrimRight,TrimLeft
*/
func main() {
	s := "Yes我爱慕课网!"              //UTF-8 编码是可变长字节  中文默认是3个字节
	fmt.Println(len(s))           //所以 length 是 19
	for i, b := range []byte(s) { // b is a byte
		fmt.Printf(" （%d %X） ", i, b)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune. rune = int32
		fmt.Printf(" （%d %X） ", i, ch)
	}
	fmt.Println()
	/*
	 （0 59）  （1 65）  （2 73）  （3 E6）  （4 88）  （5 91）  （6 E7）  （7 88）  （8 B1）  （9 E6）  （10 85）  （11 95）  （12 E8）  （13 AF）  （14 BE）  （15 E7）  （16 BD）  （17 91）  （18 21）
	 （0 59）  （1 65）  （2 73）  （3 6211）  （6 7231）  （9 6155）  （12 8BFE）  （15 7F51）  （18 21）

	  从上面的两个输出结果可以看出一个中文占了 3 个字节
	*/

	/*
		使用 utf8 库的一些转解码操作，帮我们解决底层数据的解析，便于使用
	*/
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 将 s 先转成 rune 在统计结果就可以忽律中英文不同字节的问题

	bytes := []byte(s)
	for len(bytes) > 0 {
		// utf8.DecodeRune  从 []byte() 中对 utf-8 字符进行 rune 类型解码并返回第一个 rune 字符 和 长度
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 上面的大多都是对于 string 到 rune 的底层一些解析操作。通常我们如果要使用 rune 是不需要这么复杂
	// 直接通过 rune 转换就行。需要注意的一点是： []rune(s) 转换其实是从新开了一段内存，然后按照 rune = int32
	// 一个字节占八位，也就是一个 rune 存放四个字节，无论中英文都是这样子。然后进行解析
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

}
