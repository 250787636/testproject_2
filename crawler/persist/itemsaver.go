package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"ky/ssp/crawler/engine"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d:%v", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error"+
					"saving item %v:%v",
					item, err)
			}
		}
	}()
	return out
}

// 开启elasticSearch
func save(item engine.Item) (err error) {
	client, err := elastic.NewClient(
		// Must trun off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
