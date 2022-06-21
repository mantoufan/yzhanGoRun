package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	/* 原子计数：Atomic Counters */

	// 声明一个无符号整型计数器
	var ops uint64 = 0

	// 模拟并发更新
	for i := 0; i < 50; i++ {
		go func() {
			for {
				/* AddUint64 计数器自动增加
				   使用 & 语法取出 ops 的内存地址 */
				atomic.AddUint64(&ops, 1)

				// 让出 CPU 时间片 默认 coroutines 在同一个线程
				runtime.Gosched()
			}
		}()
	}

	// 等待 1 秒，让 ops 自动加操作执行一会儿
	time.Sleep(time.Second)

	// 安全读取，当前值拷贝到 opsCopy
	opsCopy := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsCopy)
}
