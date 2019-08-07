package main

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13爬虫/crawler/persist"
	"imooc.com/doublex/learngo/13爬虫/crawler/scheduler"
	"imooc.com/doublex/learngo/13爬虫/crawler/zhenai/parser"
)

func main() {
	//initUrl := "http://www.zhenai.com/zhenghun/shanghai"
	initUrl := "http://www.zhenai.com/zhenghun"

	saver, err := persist.ItemSaver("dating_profile", "zhenai")

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    saver,
	}

	e.Run(engine.Request{
		Url:       initUrl,
		ParseFunc: parser.ParseCityList,
	})
}
