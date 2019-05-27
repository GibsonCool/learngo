package main

import "fmt"

/*
	map使用示例：
		寻找最长不含有重复字符的子串
*/

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcdeadf"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("pewwek"))
	fmt.Println(lengthOfNonRepeatingSubStr("tesgadfb"))

	//这里的结果是不对的
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))

	//采用 rune 类型转换后可以不用考虑中英文的情况，结果就准确
	fmt.Println(lengthOfNonRepeatingSubStr2("tesgadfb"))
	fmt.Println(lengthOfNonRepeatingSubStr2("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr2("一二三二一"))
}

//
func lengthOfNonRepeatingSubStr(s string) int {
	//通过map来记录每个字符的位置
	lastOccurred := make(map[byte]int)

	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		//从map中取出字符所在的位置，
		lastI, ok := lastOccurred[ch]

		//如果存在这个字符，并且超过了start.说明出现了重复字符，起始位置设置为这个重复字符上次出现的位置往后移一个
		if ok && lastI >= start {
			start = lastI + 1
		}

		//当前位置到起始位置的长度大于之前最长长度，则从新赋值最长长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

//支持中文
func lengthOfNonRepeatingSubStr2(s string) int {
	//通过map来记录每个字符的位置
	lastOccurred := make(map[rune]int)

	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		//从map中取出字符所在的位置，
		lastI, ok := lastOccurred[ch]

		//如果存在这个字符，并且超过了start.说明出现了重复字符，起始位置设置为这个重复字符上次出现的位置往后移一个
		if ok && lastI >= start {
			start = lastI + 1
		}

		//当前位置到起始位置的长度大于之前最长长度，则从新赋值最长长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
