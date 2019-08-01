package persist

import (
	"context"
	"crawler/engine"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

//获取itemSaver的input 通道
func ItemSaver(esIndex string, esUrl string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	//1.创建elasticSearch client,有可能es没启动,如果没启动，直接导致程序退出
	//modify for test
	//client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esUrl))
	//if err != nil {
	//	return out, err
	//}

	//2.goroutine 保存数据
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item "+
				"#%d: %v", itemCount, item)

			itemCount++

			//modify for test
			//err := Save(client, esIndex, item)
			//if err != nil {
			//	log.Printf("Item Saver: error "+
			//		"saving item %v:%v", item, err)
			//}
		}
	}()

	return out, nil
}

//存储数据
func Save(client *elastic.Client, esIndex string, item engine.Item) error {
	//判断类型
	if item.Type == "" {
		return errors.New("must supply item type")
	}

	//es type 和 id值在item项中，index值通过外部设置
	//index -> database
	//type -> table
	//id -> id
	indexService := client.Index().
		Index(esIndex).
		Type(item.Type).
		BodyJson(item)

	//处理id为空的情况
	if item.ID != "" {
		indexService.Id(item.ID)
	}

	//发起请求
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
