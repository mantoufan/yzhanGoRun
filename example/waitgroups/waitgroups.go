package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/* WaitGroup：等待多个协程完成 */
	var wg sync.WaitGroup
	for id := 1; id <= 5; id++ {
		wg.Add(1) // WaitGroup 计数器 + 1
		go worker(id, &wg)
	}
	// 阻塞，直到 WaitGroup 计数器恢复为 0
	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// 通知 WaitGroup 当前协程工作已经完成
	wg.Done()
}
