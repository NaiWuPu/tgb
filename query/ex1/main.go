package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{}
	urlObj, _ := url.Parse("http://baidu.com")
	data.Set("q","nihao")
	data.Set("p","1")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil{
		log.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(b))
}