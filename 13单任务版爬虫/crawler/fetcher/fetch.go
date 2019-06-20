package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	抽离根据 URL 获取内容的方法

*/
func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error : statu code ", resp.StatusCode)
	}

	//	通常网页的编码格式是不确定的
	//	通过一个第三方库来获取该网页的编码格式
	encoding := determineEncoding(resp.Body)

	// 内容通过制定编码格式转换为一个新的 reader
	newReader := transform.NewReader(resp.Body, encoding.NewDecoder())

	return ioutil.ReadAll(newReader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		log.Printf("Fetcher error: %v", e)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	//fmt.Println(encoding, name, certain)
	return encoding
}
