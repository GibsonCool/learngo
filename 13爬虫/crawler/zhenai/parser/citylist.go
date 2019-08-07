package parser

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/engine"
	"regexp"
	"strconv"
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

		result.Items = append(result.Items, "City "+string(m[2]))
		url := string(m[1])

		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: url,
				//城市下用户列表解析器
				//ParseFunc: ParseCity,

				// 由于网站改版了，直接在城市这级就获取解析用户信息
				ParseFunc: ParseCityProfile,
			},
		)

		// 多查询5页的数据，
		page := 2
		for {

			pageUrl := url + "/" + strconv.Itoa(page)
			//log.Printf("page url :%s",pageUrl)
			//
			//if resp,e :=http.Get(pageUrl);e!=nil || http.StatusOK != resp.StatusCode {
			//		break
			//}
			result.Requests = append(
				result.Requests,
				engine.Request{
					Url:       pageUrl,
					ParseFunc: ParseCityProfile,
				},
			)
			page++
			if page > 5 {
				break
			}
		}

		//if index==1{
		//	break
		//}
	}
	return result
}
