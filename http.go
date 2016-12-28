package golibs

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//获取url对应的内容，返回信息：StatusCode，body，err
func Get(requestUrl string) (int, string, error) {
	response, err := http.Get(requestUrl)
	defer response.Body.Close()
	if err != nil {
		return 0, "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, "", err
	}
	return response.StatusCode, string(body), nil
}

//获取url对应的内容，返回信息：StatusCode，body，err
func Post(requestUrl string, params url.Values) (int, string, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", requestUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return 1001, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("User-Agent", "Top4Net")
	response, err := client.Do(reqest)
	if err != nil {
		return 1002, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 1003, "", err
	}
	return response.StatusCode, string(body), nil
}

//获取当前连接的Http方法
func Method(r *http.Request) string {
	return r.Method
}
