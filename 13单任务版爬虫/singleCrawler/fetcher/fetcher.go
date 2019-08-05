package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	根据链接地址读取内容你那个返回
*/
func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code : %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
