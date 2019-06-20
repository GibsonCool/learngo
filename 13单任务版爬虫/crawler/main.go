package main

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/zhenai/parser"
)

/*
	爬取珍爱网
*/
func main() {

	engine.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		},
	)
}
