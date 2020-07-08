package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 基于sarama 第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	// tailf 包使用
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完整数据 leader 和 follow 都确认
	// 指定分区 轮询方式
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个 partition

	config.Producer.Return.Successes = true // 成功交付消息在 chanel 返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}

	msg.Topic = "web_log"

	msg.Value = sarama.StringEncoder("this is a test log")

	// 链接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
