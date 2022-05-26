package main

import "fmt"

func main() {
	// var 声明变量
	var a = "initial"
	fmt.Println(a)

	// var 声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 自动推断变量类型
	var d = true
	fmt.Println(d)

	// 声明后，初始化为 零值
	var e int
	fmt.Println(e)

	// := 声明并初始化变量
	f := "short"
	fmt.Println(f)
}
