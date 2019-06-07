package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"imooc.com/doublex/learngo/06资源管理与出错处理/errhanding/filelistingserver/filelisting"
	"net/http"
	"os"
)

/*
	如何实现统一的错误处理逻辑：(defer + panic + recover)

		因为 go 支持函数式编程，函数可以作为另一函数的参数或返回值

		所以这里使用 errWrapper() 接受 A 函数，并且输出 B 函数

		A 中去做正常的业务逻辑处理，在 errWrapper 中做统一错误处理并返回 B 函数

		这样错误处理就放到了 errWrapper 中。需要的地方将 原函数 fx 通过包装一次即可到达目的：errWrapper(fx)

	error VS  panic

		意料之中的：使用error。如：文件打不开，还可以细化自定义一些 userError   比如路径错误

		意料之外的：使用panic。如：数组越界
*/

func main() {

	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	serverError := http.ListenAndServe(":8888", nil)
	if serverError != nil {
		panic(serverError)
	}

}

// 定义 appHandler 函数类型（函数类型是表示所有包含相同参数和返回类型的函数集合）
// filelisting.HandleFileList 就是次函数类型的一个实例
type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Error("test Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		//先调用 handler(在本实例中也 就是 filelisting.HandleFileList) 进行正常业务处理
		err := handler(writer, request)
		if err != nil {
			//如果是自定义可预见的一些用户错误单独处理
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//如果正常业务处理中出错了，进行异常错误统一处理
			log.Error("Error handling request: %s\n", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			//进行最后的异常输出到前端
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

// 自定义一些可预见，可以提示给用户看的error
type userError interface {
	/*
		这里用到了接口的组合，直接把 error 接口拿过来用
		type error interface {
			Error() string
		}
	*/
	error //

	// 用于提供给用户看到错误信息
	Message() string
}
