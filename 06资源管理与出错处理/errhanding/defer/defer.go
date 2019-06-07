package main

import (
	"bufio"
	"fmt"
	"imooc.com/doublex/learngo/05函数式编程/functional/fib"
	"os"
	"runtime"
	"strings"
)

func main() {

	tryDefer()
	//fmt.Println(currentFile())
	deferValue()
	//writeFile("fib.text")
	writeFileErrorHanding("fib.txt")
}

/*
	defer 确保调用在函数结束时发生

	多个 defer 存在的时候，defer 列表是有一个调用栈的，先进后出

	即使遇到 return 或 error 依然会先执行defer

	何时使用 defer 调用：
		Open/Close   Lock/Unlock  PrintHeader/PrintFooter
*/
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	//panic("error occurred")
	fmt.Println(4)

}

/*
	使用 defer 进行文件资源管理
*/
func writeFile(fileName string) {
	file, err := os.Create(currentFile() + fileName)
	if err != nil {
		panic(err)
	}
	//正常我们的读写打开 io  操作使用完毕后都需要关闭流，避免资源浪费
	//在 go 中直接使用 defer 标记函数在结束后自动关闭
	defer file.Close()

	//为了提高效率，先使用 bufferIo 写入缓存，然后在从缓存写入文件，io操作大多有这种相识的操作，比如Java
	writer := bufio.NewWriter(file)
	//同样使用 defer 标记写完后将缓存内容写入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}
}

/*
	 错误处理：
		和 Java 大同小异。try()catch()  尽可能的将错误细化处理，如果处理不了的在抛出来
*/
func writeFileErrorHanding(fileName string) {

	//O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	//O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	//加了 O_EXCL 如果文件已经存在会报错
	file, err := os.OpenFile(currentFile()+fileName, os.O_RDWR|os.O_EXCL|os.O_CREATE, 0666)

	//err= errors.New("自定义 error 返回")
	if err != nil {
		//panic(err)
		//fmt.Println(" ile already exit")
		/*
			    os.OpenFile() 注释中最后一句是： If there is an error, it will be of type *PathError.
				所以我们可以对错误异常进行更加详细的处理
		*/
		//Type assertion 类型断言
		if pathError, ok := err.(*os.PathError); !ok {
			//如果类型转换不是 pathError 就直接输出
			panic(err)
		} else {
			fmt.Printf("^%s, %s, %s\n",
				pathError.Op,
				pathError.Err,
				pathError.Path)
		}
		return
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println("i:", i)
		fmt.Fprintln(writer, f())
	}
}

/*
	获取当前执行文件的目录
*/
func currentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(" Can not get current file info")
	}
	lastIndex := strings.LastIndex(file, "/") + 1
	file = file[:lastIndex]
	return file
}

/*
	虽然 defer 语句会在程序结束后才执行
    但是参数却是在 defer 语句时就计算
*/
func deferValue() {
	fmt.Println("===================")
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 20 {
			return
		}
	}
}
