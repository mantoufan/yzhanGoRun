package main

import (
	"fmt"
	"time"
)

func main() {
	/* 常规的通道发送和接收数据是阻塞
	   使用 default / select 实现非阻塞发送、接收
		 非阻塞的多路 select */
	messages := make(chan string, 2)
	//signals := make(chan bool)

	go func() {
		time.Sleep(time.Second * 1)
		messages <- "msg"
	}()

	// 阻塞接收：需要等待 1 秒
	msg := <-messages
	fmt.Println("block received", msg)

	// 非阻塞接收
	select {
	case msg := <-messages:
		fmt.Println("non-block received", msg)
	default:
		fmt.Println("no message received")
	}

	// 阻塞发送

	go func() {
		msg := <-messages
		fmt.Println("block received", msg)
	}()

	messages <- msg

	// 非阻塞发送
	select {
	case msg := <-messages:
		fmt.Println("non-block received", msg)
	default:
		fmt.Println("no message received")
	}
}
