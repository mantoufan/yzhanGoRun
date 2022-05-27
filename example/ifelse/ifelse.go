package main

import "fmt"

func main() {
	// 这里是一个基本例子
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 只用 if 语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 条件语句之前可以有一个声明语句
	// 声明可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
