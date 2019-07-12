package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver : got item #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) {
	client, e := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)

	if e != nil {
		panic(e)
	}

	response, e := client.Index().
		// 数据库名
		Index("dating_profile").
		// 表明
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if e != nil {
		panic(e)
	}

	fmt.Println(response)
}
