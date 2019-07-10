package main

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/scheduler"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/zhenai/parser"
	"runtime"
	"strings"
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

	//并发调度器
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		//Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
	}

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		},
	)
}

func currentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(" Can not get current file info")
	}
	lastIndex := strings.LastIndex(file, "/") + 1
	file = file[:lastIndex]
	return file
}
