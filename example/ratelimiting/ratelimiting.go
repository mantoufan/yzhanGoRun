package main

import (
	"fmt"
	"time"
)

func main() {
	/* 速率限制：GO 协程、通道和打点器限制速度 */

	/* 通过打点器阻塞通道限制速率 */
	// 声明一个缓冲为 5 的通道
	requests := make(chan int, 5)

	// 将请求发送给通道
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	// 关闭通道
	close(requests)

	// 声明一个打点器：每 200ms 执行一次
	limiter := time.Tick(time.Millisecond * 200)

	// 每次请求前阻塞 limiter 通道
	for req := range requests {
		<-limiter
		fmt.Println("request-01", req, time.Now())
	}

	/* 通道打点器自身临时限制速率 */
	burstyLimiter := make(chan time.Time, 3)

	// 缓冲 3 个请求, 快速执行
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		ticker := time.NewTicker(time.Millisecond * 2000)
		for t := range ticker.C {
			burstyLimiter <- t
		}
	}()

	// 模拟 5 个接入请求，只有前 3 个可以快速执行
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request-02", req, time.Now())
	}
}
