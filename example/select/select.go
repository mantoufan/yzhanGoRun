package main

import (
	"fmt"
	"time"
)

func main() {
	/* 通道选择器
	   同时等待多个通道操作 */
	chan1 := make(chan string)
	chan2 := make(chan string)

	// 模拟并行异步 RPC 远程过程调用函数
	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- "1"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "2"
	}()

	/* Select 同时等待两个值
	   1 秒后，收到 chan1 值
	   2 秒后，收到 chan2 值 */
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("received", msg1)
		case msg2 := <-chan2:
			fmt.Println("received", msg2)
		}
	}
}
