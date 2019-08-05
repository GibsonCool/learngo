package engine

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/singleCrawler/fetcher"
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

		log.Printf("Featching %s", r.Url)
		body, e := fetcher.Fetch(r.Url)
		if e != nil {
			log.Printf("Fetcher : error fetching url %s : %v", r.Url, e)
			// 如果某个请求错误了，直接跳过
			continue
		}

		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		//结果这里先直接打印
		for _, item := range parseResult.Items {
			log.Printf("Got Item %s", item)
		}

	}
}
