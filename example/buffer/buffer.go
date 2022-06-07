package main

import "fmt"

func main() {
	/* 通道缓冲
	   默认：通道无缓冲，只有对应接收方 <- 准备好，才发送 ->
		 可缓存通道：没有对应接受方，缓存限定数量得值
	*/
	// 创建通道，最多允许缓存 2 个值
	channel := make(chan string, 3)

	// 发送值到通道，开启缓冲
	channel <- "0"
	go func() {
		channel <- "1" // 1 或 2 都可能在前
	}()
	go func() {
		channel <- "2" // 1 或 2 都可能在前
	}()
	channel <- "3"

	// 接收值从通道
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
