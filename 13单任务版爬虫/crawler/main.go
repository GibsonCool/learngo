package main

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/persist"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/scheduler"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/zhenai/parser"
	"runtime"
	"strings"
)

/*
	爬取珍爱网
*/
func main() {

	//initUrl := "http://www.zhenai.com/zhenghun"
	initUrl := "http://www.zhenai.com/zhenghun/shanghai"

	//并发调度器
	e := engine.ConcurrentEngine{
		// 简单调度器
		//Scheduler: &scheduler.SimpleScheduler{},
		// 并发队列调度器
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(
		engine.Request{
			Url:        initUrl,
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
