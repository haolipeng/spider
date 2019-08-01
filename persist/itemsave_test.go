package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

/////////////////////////elastic variable/////////////////////////////////
var (
	elastic_database = "dating_profile"
	elastic_url      = "http://192.168.227.134:9200"
)

func TestItemSaver(t *testing.T) {
	//Try to start up elastic search
	//here using docker go client
	//client连接ES
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(elastic_url))
	if err != nil {
		panic(err)
	}

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1320662004",
		Type: "zhenai",
		ID:   "1320662004",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "5000-8000",
			Gender:     "女",
			Xingzuo:    "白羊座",
			Occupation: "人事/行政",
			Marriage:   "未婚",
			House:      "已购房",
			Hukou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	//Save expected item
	err = Save(client, elastic_database, expected)
	if err != nil {
		panic(err)
	}

	//esId := "AWjMngKUnKIAbCn7xlh7"

	//Fetch item
	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.ID).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", resp.Source)

	//Verify result 校验结果
	var actual engine.Item
	json.Unmarshal(([]byte)(*resp.Source), &actual)

	//json序列化和反序列化
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual == expected {
		fmt.Println("itemSaver function test Passed")
	} else {
		fmt.Println("itemSaver function test failed")
	}
}
