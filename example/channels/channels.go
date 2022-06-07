package main

import "fmt"

func main() {
	/* 通道（Channels）是连接多个 Go 协程的管道 */
	// 创建通道
	messages := make(chan string)

	// 发送：值到通道，默认：阻塞
	go func() {
		messages <- "messages"
	}()

	// 接收：值从通道，默认：阻塞，不使用任何其它同步操作
	receivesMsg := <-messages
	fmt.Println(receivesMsg)
}
