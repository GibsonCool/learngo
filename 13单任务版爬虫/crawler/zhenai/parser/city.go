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
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)

		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				//巧妙利用函数式编程，闭包的优势链接上下不兼容的函数
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, name)
				},
			})

		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
