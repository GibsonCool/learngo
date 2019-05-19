package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	scanner := bufio.NewScanner(file)
	//省略初始条件,递增条件实现while
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

	readFile("abc.txt")

	forever()
}
