package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
	爬取珍爱网
*/
func main() {
	url := "http://www.zhenai.com/zhenghun"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
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

	bytes, err := ioutil.ReadAll(newReader)
	if err != nil {
		panic(err)
	}

	findCityList(bytes)
}

// 通过正则匹配找到所有城市列表信息
func findCityList(bytes []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(bytes, -1)

	for _, m := range matches {
		fmt.Printf("city:%s  url:%s \n", m[2], m[1])
	}

	fmt.Println("match count :", len(matches))
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		panic(e)
	}
	encoding, name, certain := charset.DetermineEncoding(bytes, "")
	fmt.Println(encoding, name, certain)
	return encoding
}
