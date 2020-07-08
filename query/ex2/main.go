package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
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
	// 长连接
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	log.Println(err)
	//}
	tr := &http.Transport{
		Proxy:                  nil,
		DialContext:            nil,
		Dial:                   nil,
		DialTLS:                nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      true,
		DisableCompression:     false,
		MaxIdleConns:           0,
		MaxIdleConnsPerHost:    0,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        0,
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  0,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      false,
	}
	client := http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	resp, err := client.Do(req)
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



