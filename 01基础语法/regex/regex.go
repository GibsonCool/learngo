package main

import (
	"fmt"
	"regexp"
)

/*
	正则表达式
*/

const text = `
My email is cxxcexo@163.com
email1 is abc@def.org
email2 is   kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	match := re.FindAllString(text, -1)
	fmt.Println(match)

	fmt.Println()

	match2 := re.FindAllStringSubmatch(text, -1)
	for _, m := range match2 {
		fmt.Println(m)
	}
}
