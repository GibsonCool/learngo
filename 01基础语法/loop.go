package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//Go中的循环只有for

/* 将一个整数转换为二进制字符串 */
func converToBin(v int) string {

	result := ""
	for ; v > 0; v /= 2 {
		result = strconv.Itoa(v%2) + result
	}
	return result
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContent(file)
}

// 使用系统接口 io.Reader 封装扩展方法
func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	//省略初始条件,递增条件实现while，当 scanner.Scan()返回false 循环结束
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	//省略初始条件，判断条件，递增条件，实现死循环
	for {
		fmt.Println("text")
	}
}

func main() {
	fmt.Println(
		converToBin(1),  //1
		converToBin(2),  //10
		converToBin(5),  //101
		converToBin(13), //1101
	)

	readFile("01基础语法/abc.txt")

	s := `to be or not to be 
		this is a question
		talk is cheep
		`
	printFileContent(strings.NewReader(s))

	//forever()
}
