package main

import "fmt"

func main() {
	/* 指针：通过引用传递值或数据结构 */
	i := 1
	zeroval(i)
	// 值保持不变
	fmt.Println(i)
	// & 获取指针
	fmt.Println(&i)
	// 值被改变
	zeroptr(&i)
	fmt.Println(i)
}

// 传递值
func zeroval(i int) {
	i = 0
}

// 传递引用
func zeroptr(i *int) {
	*i = 0 // 解引用，改变真实地址值
}
