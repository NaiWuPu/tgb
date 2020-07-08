package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"180.76.233.214:2379"},
		DialTimeout:          5 * time.Second,
	})

	// watch 操作 用来获取未来更改的通知
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
	}
	fmt.Println("connect to etcd success")

	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "key", "value")
	cancel()

	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "key")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd faild, err:%v\n", err)
		return
	}

	for _, v := range resp.Kvs {
		fmt.Printf("%s:%s\n", v.Key, v.Value)
	}

}
