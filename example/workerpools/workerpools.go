package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	/* 线程池
	   2 个通道
	   工作：工作通道
		 结果：结果通道 */
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动 3 个线程
	for workerId := 1; workerId < 3; workerId++ {
		go worker(workerId, jobs, results)
	}

	// 发送 9 个 jobs
	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)

	// 最后，我们收集所有这些任务的返回值
	for i := 1; i <= 9; i++ {
		<-results
	}
}
