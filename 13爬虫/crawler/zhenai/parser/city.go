package parser

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/engine"
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
	for index, m := range matches {

		result.Items = append(result.Items, "User  "+string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				//城市下用户列表解析器
				ParseFunc: engine.NilParser,
			},
		)
		//TODO: 需要删除，测试用来控制条数
		if index == 3 {
			break
		}
	}
	return result
}
