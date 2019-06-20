package parser

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 抽离城市列表信息解析器方法
func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParse,
			})
	}

	return result
}
