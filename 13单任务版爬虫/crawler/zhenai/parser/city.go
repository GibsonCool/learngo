package parser

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))

		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseProfile,
			})

		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
