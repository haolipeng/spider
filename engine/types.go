package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //url对应Parser解析器
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string      //请求的url
	Type    string      //爬虫对应的业务 es type
	ID      string      //必要的 es id
	Payload interface{} //具体类型根据业务来决定，如珍爱网是Profile类型
}
