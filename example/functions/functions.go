package main

import "fmt"

func main() {
	/* 函数 */
	// 连续参数，同样类型，仅声明最后一个参数
	r := plus(1, 2, 3)
	fmt.Println(r)

	// 多返回值
	a, b := multiReturn()
	fmt.Println(a, b)

	// 使用空白标识符
	_, c := multiReturn()
	fmt.Println(c)

	// 可变参数
	d := sum(1, 2, 3)
	fmt.Println(d)

	nums := []int{1, 2, 3}
	e := sum(nums...)
	fmt.Println(e)
}

// 连续参数，同样类型，仅声明最后一个参数类型
func plus(a, b, c int) int {
	return a + b + c
}

// 多返回值
func multiReturn() (int, int) {
	return 1, 2
}

// 可变参数
func sum(nums ...int) (total int) {
	total = 0
	for _, num := range nums {
		total += num
	}
	return
}
