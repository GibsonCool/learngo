package parser

import (
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13单任务版爬虫/crawler/model"
	"regexp"
)

const infoRe = `<div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`

/*
	个人信息详情页解析
*/
func ParseProfile(contexts []byte) engine.ParseResult {

	re := regexp.MustCompile(infoRe)
	allSubmatch := re.FindAllSubmatch(contexts, -1)
	profile := model.Profile{}
	for index, m := range allSubmatch {

		//fmt.Printf("person info index %d: %s\n", index, string(m[1]))
		/*
			网页改版后获取内容可能不同，需要从写编写解析方法
		*/
		value := string(m[1])
		switch index {
		case 0:
			profile.Marriage = value
		case 1:
			profile.Age = value
		case 2:
			profile.Xinzuo = value
		case 3:
			profile.Height = value
		case 4:
			profile.Weight = value
		case 5:
			profile.WorkPlace = value
		case 6:
			profile.Income = value
		case 7:
			profile.Occupation = value
		case 8:
			profile.Education = value
		}

	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result

}
