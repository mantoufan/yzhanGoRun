package main

import (
	"fmt"
	"time"
)

func main() {
	/* 定时器
	   未来某一时刻的独立事件
		 告诉定时器需要等待，提供一个用于通知的通道
		 这里的定时器将等待 2 秒
	*/
	timer1 := time.NewTimer(time.Second * 1)

	// 直到定时器通道 timer1.C 明确发送
	// 定时器将一直阻塞
	<-timer1.C
	fmt.Println("Timer 1 expired")

	/* 停止定时器 */
	timer2 := time.NewTimer(time.Second * 10)

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
