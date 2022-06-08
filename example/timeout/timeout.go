package main

import (
	"fmt"
	"time"
)

func main() {
	/* 异步请求的超时处理
	   用 select 和 通道 实现 */
	c := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c <- "msg"
	}()

	// 等待 1 秒，没有收到，返回：1 秒超时
	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1 second")
	}

	go func() {
		time.Sleep(time.Second * 2)
		c <- "msg"
	}()

	// 等待 2 秒，已收到，返回：消息
	select {
	case msg := <-c:
		fmt.Println("receive:" + msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 1 second")
	}
}
