package main

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13爬虫/crawler/zhenai/parser"
)

func main() {
	initUrl := "http://www.zhenai.com/zhenghun"

	engine.Run(engine.Request{
		Url:       initUrl,
		ParseFunc: parser.ParseCityList,
	})
}
