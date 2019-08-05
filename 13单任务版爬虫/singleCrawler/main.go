package main

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/singleCrawler/engine"
	"imooc.com/doublex/learngo/13单任务版爬虫/singleCrawler/zhenai/parser"
)

func main() {
	initUrl := "http://www.zhenai.com/zhenghun"

	engine.Run(engine.Request{
		Url:       initUrl,
		ParseFunc: parser.ParseCityList,
	})
}
