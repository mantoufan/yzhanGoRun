package main

import "fmt"

func main() {
	/* 闭包 */
	// 创建闭包
	nextI := closures()
	fmt.Println(nextI())
	fmt.Println(nextI())
	fmt.Println(nextI())

	newNextI := closures()
	fmt.Println(newNextI())
	fmt.Println(newNextI())
	fmt.Println(newNextI())
}

func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
