package parser

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/singleCrawler/engine"
	"regexp"
)

/*
	城市用户信息列表解析器
*/
func ParseCity(contents []byte) engine.ParseResult {

	//通过正则匹配获取网页内容需要提取出的信息
	re := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {

		result.Items = append(result.Items, "User  "+string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				//城市下用户列表解析器
				ParseFunc: engine.NilParser,
			},
		)
	}
	return result
}
