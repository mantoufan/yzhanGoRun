package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 字符串重复几次 */
	// func Repeat(s string, count int) string
	// count = 0，返回空
	// 报错：count < 0 或返回值长度超过 string 上限
	fmt.Println(strings.Repeat("a", 0))
	fmt.Println(strings.Repeat("b", 1))
}
