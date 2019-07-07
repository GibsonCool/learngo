package main

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/scheduler"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/zhenai/parser"
)

/*
	爬取珍爱网
*/
func main() {

	// 简单调度器
	//engine.SimpleEngine{}.Run(
	//	engine.Request{
	//		Url:        "http://www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	},
	//)

	// 并发调度器
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		},
	)
}
