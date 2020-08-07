package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"tgb/gzip/zip"
)

const BaseUploadPath = "./"

func main() {
	// 读取证书
	pool := x509.NewCertPool()
	caCertPath := "cert/cacert.pem"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err != nil {
		fmt.Println("LoadX509keypair err: ", err)
		return
	}
	// 包含证书请求
	//    tr := &http2.Transport{  // http2协议
	tr := &http.Transport{ // http1.1协议
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
			//InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}

	// 请求成功，收到返回体
	//resp, err := client.Get("https://localhost:8088")
	resp, err := client.Get("https://0000000000000000000000000000000000000000:8088/")
	if err != nil {
		fmt.Println("http get error: ", err)
		return
	}
	defer resp.Body.Close()
	// 读取返回文件
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(resp.Body) err")
	}
	// 读取文件头，并复制到目录
	err = ioutil.WriteFile(BaseUploadPath+resp.Header.Get("filename"), buff, 0755)
	if err != nil {
		fmt.Println("ioutil.WriteFile err")
	}
	// zip解压缩
	err = zip.DeCompress(BaseUploadPath+resp.Header.Get("filename"), BaseUploadPath)
	if err != nil {
		fmt.Println("zip compress err")
	}
}
