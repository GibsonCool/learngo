package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	抽离根据 URL 获取内容的方法

*/
func Fetch(url string) ([]byte, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	// 添加请求头，模拟浏览器发送请求，否则个人详情页面请求会 403
	request.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36",
	)
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		fmt.Println("Error url :", url)
		fmt.Println("Error : statu code ", resp.StatusCode)
	}

	body := bufio.NewReader(resp.Body)
	//	通常网页的编码格式是不确定的
	//	通过一个第三方库来获取该网页的编码格式
	encoding := determineEncoding(body)

	// 内容通过制定编码格式转换为一个新的 reader
	newReader := transform.NewReader(body, encoding.NewDecoder())

	return ioutil.ReadAll(newReader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, e := r.Peek(1024)
	if e != nil {
		log.Printf("Fetcher error: %v", e)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	//fmt.Println(encoding, name, certain)
	return encoding
}
