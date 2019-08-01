package fetcher

import (
	"fmt"
	"testing"
)

const userPrivateInfo = "http://album.zhenai.com/u/1690271375"

func TestFetch(t *testing.T) {
	result, err := FetchWithUserAgent(userPrivateInfo)
	if err != nil {
		fmt.Println("fetch method error")
		return
	}

	fmt.Printf("%s\n", result)
}
