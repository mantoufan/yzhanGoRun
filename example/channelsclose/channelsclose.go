package main

import "fmt"

func main() {
	/* Go 协程：关闭一个通道，意味着不能再向通道发送值 */
	/* 这个特性可以用来给通道的接收方传达工作已经完成的信息 */
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 循环的从 jobs 接收数据
			j, more := <-jobs
			fmt.Println(more)
			if more {
				fmt.Println("received job", j)
			} else {
				// 利用通道的阻塞机制，将这里变成同步
				fmt.Println("received all jobs")
				done <- true
				fmt.Println("done")
				return
			}
		}
	}()

	// 这里使用 jobs 发送 3 个任务到工作函数
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// 关闭 jobs
	close(jobs)

	// 通知全部任务结束
	<-done
}
