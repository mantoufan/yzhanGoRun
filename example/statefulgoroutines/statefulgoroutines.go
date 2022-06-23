package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	/* Stateful Go 协程 */
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// reads 和 writes 通道
	// 发布读和写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	/* 声明一个有 Stateful Go 协程 */
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	/* 100 个 Go 协程：读 操作 */
	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	/* 100 个 Go 协程：写 操作 */
	for i := 0; i < 100; i++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让 Go 协程跑 1s
	time.Sleep(time.Second)

	// 读取 ops 值
	readCopy := atomic.LoadUint64(&readOps)
	fmt.Println("readCopy:", readCopy)
	writeCopy := atomic.LoadUint64(&writeOps)
	fmt.Println("writeCopy:", writeCopy)
}
