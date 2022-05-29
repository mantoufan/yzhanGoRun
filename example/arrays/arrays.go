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

	// 遍历数组：打印索引 + 元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 数组长度：未初始化的元素，值为 0
	var b [3]int = [3]int{1}
	fmt.Println(b)

	// 数组长度是数组类型一部分
	c := [...]int{0}
	d := [...]int{0, 1}
	fmt.Printf("Type: %T\n", c)
	fmt.Printf("Type: %T\n", d)

	// 不同类型不能赋值
	// c = d 编译错误
}
