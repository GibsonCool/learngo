package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

/*
	使用 http 来当做客户端发起请求

*/
var url = "http://www.imooc.com"

func main() {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	// 添加请求头，模拟手机端发送请求
	request.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Mobile Safari/537.36",
	)

	client := http.Client{
		// 创建重定向方法，如果有重定向会走入这个方法
		// 由于请求头中加入了模拟手机发起的请求，因此请求的网页会重定向到请求手机版的页面
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response)

}
