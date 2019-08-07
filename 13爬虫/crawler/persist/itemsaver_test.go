package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"imooc.com/doublex/learngo/13爬虫/crawler/model"
	"testing"
)

func TestItemSaver(t *testing.T) {

	profile := model.Profile{
		Id:                "1111111",
		Name:              "痴情",
		Gender:            "男士",
		CurrentResidence:  "新疆阿勒泰",
		Age:               59,
		SalaryOrEducation: "3001-5000元",
		WeddingStatus:     "离异",
		Height:            175,
		Remarks:           "",
		InfoLink:          "http://album.zhenai.com/u/1890873467",
	}

	err := save(profile)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index("dating_profile").Type("zhenai").Id("1111111").Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", result.Source)
	var actual model.Profile
	err = json.Unmarshal(result.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != profile {
		t.Errorf("got %v: expected %v", actual, profile)
	}
}
