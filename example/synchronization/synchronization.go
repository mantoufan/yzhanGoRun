package main

import (
	"fmt"
	"time"
)

func main() {
	/* 协程 · 同步化 */
	// 创建通道
	done := make(chan bool, 1)
	// 将通道传入协程
	go worker(done)
	// 接收通知前，协程一直阻塞
	<-done
}

func worker(done chan bool) {
	fmt.Println("working……")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送通知
	done <- true
}
