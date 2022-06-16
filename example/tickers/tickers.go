package main

import (
	"fmt"
	"time"
)

func main() {
	/* 打点器：固定时间重复执行 */

	// 每 500 ms 给通道发送一次值
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 打点器：停止
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
