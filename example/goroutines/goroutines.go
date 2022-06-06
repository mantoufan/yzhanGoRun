package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	/* 协程 gorontine ：轻量级的线程 */
	f("阻塞式调用")

	// Gorontine 协程并发调用
	go f("协程并发调用")
	// 匿名函数启用 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("匿名函数并发调用")

	// 协程在独立 Go 协程异步运行
	// 程序直接运行到这行 Scanln 需要输入
	var name string
	fmt.Println("请输入名称：")
	fmt.Scanln(&name)
	fmt.Println("名称：" + name + " 程序退出")
}
