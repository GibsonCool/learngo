package engine

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		// step1: 根据 URL 获取对应内容
		body, err := fetcher.Fetch(r.Url)
		log.Printf("Fetcher url is :%s ", r.Url)
		// 如果这次解析错误，直接跳过，比抛异常，保证整个爬虫后续可以正常运行
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		// step2: 根据当前 request 的解析函数解析获取到的内容
		parseResult := r.ParserFunc(body)

		// step3: 将解析结果中的新 request是 加入到队列中
		requests = append(requests, parseResult.Requests...)

		// 输出一下解析结果中的内容看是否正确
		for _, item := range parseResult.Items {
			log.Printf("Got item  %s", item)
		}
	}
}
