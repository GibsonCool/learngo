package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const Pattern = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

/*
	定义这个函数进行统一正常的业务处理，
	如有异常直接返回，交给 errWrapper 函数去进行统一错误处理
*/
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, Pattern) != 0 {
		return userError("path must start with" + Pattern)
	}

	path := request.URL.Path[len(Pattern):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
