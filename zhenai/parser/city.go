package parser

import (
	"crawler/engine"
	"regexp"
)

const cityNextPage = `href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`
const city = `<a href="(http://album.zhenai.com/u/([0-9]+))"[^>]*>([^<]+)</a>`

var cityRe = regexp.MustCompile(city)
var cityNextPageRe = regexp.MustCompile(cityNextPage)

func ParserCity(contents []byte) engine.ParseResult {
	//获取城市的名称和对应的url
	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		id := string(m[2])
		name := string(m[3])

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParserProfile(c, id, url, name)
			},
		})

	}

	//获取下一页内容
	matches = cityNextPageRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
	}

	return result
}
