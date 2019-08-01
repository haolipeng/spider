package parser

import (
	"crawler/engine"
	"regexp"
)

//过滤城市列表的正则表达式
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

var re = regexp.MustCompile(cityListRe)

//解析城市列表
func ParseCityList(contents []byte) engine.ParseResult {
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParserCity,
		})
	}

	return result
}
