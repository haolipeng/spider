package parser

import (
	_ "crawler/model"
	"fmt"
	_ "fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//准备各种各样的输入，这里需要网络链接，如果测试机器网络不同，或测试的网页不存在了
	//contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("zhenai.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	//verify result
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s\n", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests;but have %d", resultSize, len(result.Requests))
	}

	fmt.Println("test successed")
}
