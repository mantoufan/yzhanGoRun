package main

import "fmt"

func main() {
	/* 遍历通道中的值 */
	// 遍历 queue 通道中的两个值
	queue := make(chan string, 2)
	queue <- "1"
	queue <- "2"

	// 如果不 close 后面的循环会阻塞，等待第三个值
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
