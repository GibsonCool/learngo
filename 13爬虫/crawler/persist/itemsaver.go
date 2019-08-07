package persist

import (
	"context"
	"github.com/olivere/elastic"
	"imooc.com/doublex/learngo/13爬虫/crawler/model"
	"log"
)

func ItemSaver(dataIndex string, dataType string) (chan interface{}, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out

			itemCount++
			if value, ok := item.(model.Profile); ok {
				err := save(client, value, dataIndex, dataType)
				if err != nil {
					log.Printf("ItemSaver : error saving item %v:%v", item, err)
				}
			}
			log.Printf("ItemSaver Got item #%d:%v", itemCount, item)
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item model.Profile, dataIndex string, dataType string) error {

	indexService := client.Index().Index(dataIndex).Type(dataType)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Id(item.Id).BodyJson(item).Do(context.Background())

	if err != nil {
		return err
	}
	return err
}
