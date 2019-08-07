package engine

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {

	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		//结果这里先直接打印
		for _, item := range parseResult.Items {
			log.Printf("Got Item %s", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {

	log.Printf("Featching %s", r.Url)
	body, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Printf("Fetcher : error fetching url %s : %v", r.Url, e)
		// 如果某个请求错误了，直接跳过
		return ParseResult{}, e
	}

	return r.ParseFunc(body), nil
}
