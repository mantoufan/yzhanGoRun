package main

import (
	"fmt"
	"math"
)

func main() {
	// 声明常量
	const s string = "constant"
	fmt.Println(s)

	// const 可以出现在任何 var 出现的地方
	const n = 5e9

	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值型常量没有确定的类型，直到被给定
	// 显式的类型转化
	fmt.Println(int64(d))

	// 当上下文需要时，比如变量赋值或者函数调用
	// 一个数可以给定一个类型，举个例子
	// math.Sin 函数需要一个 float64 参数
	fmt.Println(math.Sin(n))
}
