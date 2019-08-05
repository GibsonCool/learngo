package parser

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/singleCrawler/engine"
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
	for index, m := range matches {
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
				ParseFunc: ParseCity,
			},
		)
		if index == 3 {
			break
		}
	}
	return result
}
