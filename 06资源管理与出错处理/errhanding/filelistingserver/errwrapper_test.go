package main

import (
	"errors"
	"imooc.com/doublex/learngo/06资源管理与出错处理/errhanding/filelistingserver/filelisting"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return filelisting.UserError("用户自定义错误")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknow(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

type testStruct struct {
	h    appHandler
	code int
	msg  string
}

var tests = []testStruct{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "用户自定义错误"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknow, 500, "Internal Server Error"},
}

/*
	测试HTTP服务 err 统一处理逻辑是否过关

	使用假的 Request/Response
*/
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		//接触 httptest 库提供方法所需要的 reader  writer
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://imooc.com", nil)

		//最终测试调用进过 errWrapper 处理包装返回的统一错误处理匿名方法
		f(response, request)

		verifyResponse(response.Result(), tt, t)

	}
}

/*
	通过模拟真正的的起服务测试 http
*/
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt, t)
	}
}

func verifyResponse(response *http.Response, tt testStruct, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	//不符合预期结果，输出
	if response.StatusCode != tt.code || body != tt.msg {
		t.Errorf(
			"预期（%d, %s），实际（%d, %s）",
			tt.code, tt.msg,
			response.StatusCode, body,
		)
	}
}
