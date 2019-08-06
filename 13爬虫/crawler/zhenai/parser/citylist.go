package parser

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/engine"
	"regexp"
)

/*
	城市列表信息解析器
*/

func ParseCityList(contents []byte) engine.ParseResult {
	//通过正则匹配获取网页内容需要提取出的信息
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//fmt.Printf("%s\n",m)

		//for _,subM :=range  m {
		//	fmt.Printf("%s",subM)
		//	fmt.Println( )
		//}

		//fmt.Printf("City: %s          URL: %s\n", m[2], m[1])

		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				//城市下用户列表解析器
				//ParseFunc: ParseCity,

				// 由于网站改版了，直接在城市这级就获取解析用户信息
				ParseFunc: ParseCityProfile,
			},
		)

	}
	return result
}
