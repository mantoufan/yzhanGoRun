package main

import "fmt"

func main() {
	/* 递归 */
	fmt.Println(recursion(3))
	// 斐波那契数列
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
}

func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n-1)
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
