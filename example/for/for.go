package main

import "fmt"

func main() {
	// 最基础：单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// 初始，条件，后序
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	// 不带条件 for 循环将一直重复执行
	// 直到循环体内使用 break 或 return
	for {
		fmt.Println("loop")
		break
	}
	// continue 跳到下一个循环迭代
	for i := 0; i <= 3; i++ {
		if i == 0 {
			continue
		}
		fmt.Println(i)
	}
}
