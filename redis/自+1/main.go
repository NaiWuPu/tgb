package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 自增操作
func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "180.76.233.214:6380",
		Password: "Ulb10321", // no password set
		DB:       0,          // use default DB
	})
	s := client.Get("va").Val()
	fmt.Println(s)
	if s == "" {
		fmt.Println("yes")
	}
	s, err := client.Get("va").Result()
	fmt.Println(s)
	fmt.Println(err)

}

func main() {
	incr()
}
