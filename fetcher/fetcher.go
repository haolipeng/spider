package fetcher

import (
	"bufio"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//获取url对应的utf-8内容
const userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"

//设置web页面抓取速率
var rateLimiter = time.Tick(1000 * time.Millisecond)

func FetchWithUserAgent(urlSeed string) ([]byte, error) {
	<-rateLimiter
	req, err := http.NewRequest(http.MethodGet, urlSeed, nil)
	if err != nil {
		log.Println("http NewRequest method error")
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			errors.New("Error Status code")
	}

	//使用<meta charset="utf-8">判断字符集类型gbk or utf-8
	bodyReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bodyReader)

	respReader := transform.NewReader(bodyReader, encode.NewDecoder())
	all, err := ioutil.ReadAll(respReader)
	return all, err
}

//判断文本内容的编码方式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetch Error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "charset")
	return e
}
