package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

//可选择不同的输出方式
///////////////////////////////seed variable//////////////////////////////////////
const shanghai_seed = "http://www.zhenai.com/zhenghun/shanghai"

/////////////////////////elastic variable/////////////////////////////////
var (
	elastic_database = "dating_profile"
	elastic_url      = "http://192.168.227.134:9200"
)

func main() {

	//TODO:Item server要有不同的输出方式，如数据库，表格，网络传输
	itemChan, err := persist.ItemSaver(elastic_database, elastic_url)
	if err != nil {
		panic(err)
	}

	//构造并发引擎
	e := engine.ConcurrentEngine{
		&scheduler.QueuedScheduler{},
		1,
		itemChan,
	}

	e.Run(engine.Request{
		Url:        shanghai_seed,
		ParserFunc: parser.ParserCity,
	})
}
