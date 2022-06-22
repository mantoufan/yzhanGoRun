package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	/* 互斥锁 */
	state := map[int]int{}

	// 用 mutex 同步访问 state
	mutex := &sync.Mutex{}

	// 记录操作次数
	var opt int64 = 0

	// 模拟重复读取 state
	for i := 0; i < 100; i++ {
		total := 0
		go func() {
			mutex.Lock()
			total += state[rand.Intn(3)]
			mutex.Unlock()
			atomic.AddInt64(&opt, 1)
			runtime.Gosched()
		}()
	}

	// 模拟重复写入 state
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			state[rand.Intn(3)] = rand.Intn(3)
			mutex.Unlock()
			atomic.AddInt64(&opt, 1)
			runtime.Gosched()
		}()
	}

	// 子弹飞一会儿
	time.Sleep(time.Second)

	// 最终操作数
	optCopy := atomic.LoadInt64(&opt)
	fmt.Println(optCopy)

	// 最终状态
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
