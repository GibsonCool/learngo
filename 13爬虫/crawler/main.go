package main

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13爬虫/crawler/scheduler"
	"imooc.com/doublex/learngo/13爬虫/crawler/zhenai/parser"
)

func main() {
	initUrl := "http://www.zhenai.com/zhenghun"

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:       initUrl,
		ParseFunc: parser.ParseCityList,
	})
}
