package main

import "fmt"

func main() {
	/* 数组：固定长度且编号的序列 */
	// 数组 a 存放 3 个 int
	var a [3]int
	fmt.Println("a", a)

	// 设置或返回数组的值
	a[1] = 1
	fmt.Println("set:", a)
	fmt.Println("get:", a[1])

	// 返回数组长度
	fmt.Println("len:", len(a))

	// 声明并初始化数组
	a = [3]int{0, 1, 2}
	fmt.Println("a", a)

	// 多维数组：一维数组组合而成
	var twoDimensionalArray [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoDimensionalArray[i][j] = i*5 + j
		}
	}
	fmt.Println("twoDimensionalArray", twoDimensionalArray)
}
