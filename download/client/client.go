package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
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

	//    tr := &http2.Transport{  // http2协议
	tr := &http.Transport{ // http1.1协议
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
			//InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}

	//resp, err := client.Get("https://localhost:8088")
	resp, err := client.Get("https://0000000000000000000000000000000000000000:8088/")
	if err != nil {
		fmt.Println("http get error: ", err)
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
}
