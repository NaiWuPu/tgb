package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type DoLog struct {
}

type LogRequest struct {
	Address string `json:"address"`
	Port    string `json:"port"`
	Key     string `json:"key"`
}

type LogResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var tempNum = 0

func (c *DoLog) ToLogin(req LogRequest, res *LogResponse) error {
	log.Println(req)
	if tempNum%2 == 0 {
		res.Code = 200
		res.Message = "success"
	} else {
		res.Code = 500
		res.Message = "error"
	}
	tempNum++
	return nil
}

func main() {
	_ = rpc.Register(new(DoLog)) // 注册rpc服务
	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	log.Println("start connection")

	go demoResponse()

	fmt.Printf("%s\n", "start connection")
	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}
		go func(conn net.Conn) { // 并发处理客户端请求
			fmt.Printf("%s,my add:%v peer add %v\n", "new client in coming", conn.LocalAddr(), conn.RemoteAddr())
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

func demoResponse() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := LogRequest{"127.0.0.1", "443", "key"}
	var res = new(LogResponse)
	p, e := json.Marshal(req)
	fmt.Println(string(p), e)
	err = conn.Call("DoLog.ToLogin", req, res) // 乘法运算
	if err != nil {
		log.Fatalln("log error: ", err)
	}
	log.Println(*res)

	err = conn.Call("DoLog.ToLogin", req, &res)
	if err != nil {
		log.Fatalln("log error: ", err)
	}
	log.Println(*res)
}
