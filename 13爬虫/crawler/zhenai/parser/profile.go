package parser

import (
	"imooc.com/doublex/learngo/13爬虫/crawler/engine"
	"imooc.com/doublex/learngo/13爬虫/crawler/model"
	"regexp"
	"strconv"
)

//字节转换成整形
func BytesToInt(b []byte) int {

	x, _ := strconv.Atoi(string(b))
	return x
}

var nameUrlRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
var genderRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var currentResidenceRe = regexp.MustCompile(`<td><span class="grayL">居住地：</span>([^<]+)</td>`)
var ageRe = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([^<]+)</td>`)
var monthlySalaryRe = regexp.MustCompile(`<td><span class="grayL">月[^薪]+薪：</span>([^<]+)</td>`)
var weedingStatusRe = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td width="180"><span class="grayL">身[^高]+高：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="grayL">学[^历]+历：</span>([^<]+)</td>`)

func ParseCityProfile(contents []byte) engine.ParseResult {

	nameUrlMatches := nameUrlRe.FindAllSubmatch(contents, -1)
	genderMatches := genderRe.FindAllSubmatch(contents, -1)
	currentResMatches := currentResidenceRe.FindAllSubmatch(contents, -1)
	ageMatches := ageRe.FindAllSubmatch(contents, -1)
	monthlySalMatches := monthlySalaryRe.FindAllSubmatch(contents, -1)
	weedingStaMatches := weedingStatusRe.FindAllSubmatch(contents, -1)
	heightMatches := heightRe.FindAllSubmatch(contents, -1)
	educationMatches := educationRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	menIndex, womanIndex := 0, 0
	for index, m := range nameUrlMatches {
		gender := string(genderMatches[index][1])
		var salaryOrEducation string
		if gender != "" && gender == "男士" {
			salaryOrEducation = string(monthlySalMatches[menIndex][1])
			menIndex++
		} else {
			salaryOrEducation = string(educationMatches[womanIndex][1])
			womanIndex++
		}

		profile := model.Profile{
			Name:              string(m[2]),
			Gender:            string(genderMatches[index][1]),
			Age:               BytesToInt(ageMatches[index][1]),
			WeddingStatus:     string(weedingStaMatches[index][1]),
			CurrentResidence:  string(currentResMatches[index][1]),
			SalaryOrEducation: salaryOrEducation,
			Height:            BytesToInt(heightMatches[index][1]),
			InfoLink:          string(m[1]),
		}

		result.Items = append(result.Items, profile)
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:       string(m[1]),
				ParseFunc: engine.NilParser,
			},
		)

	}

	return result
}
