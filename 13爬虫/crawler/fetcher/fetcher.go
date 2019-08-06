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
	client := &http.Client{}

	request, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		return nil, e
	}

	//去该网页上扒取的请求头信息，随时可能失效，防止页面403
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code : %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
