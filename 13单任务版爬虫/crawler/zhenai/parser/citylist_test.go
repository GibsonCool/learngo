package parser

import (
	"fmt"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {

	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", contents)
}
