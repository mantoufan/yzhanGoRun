package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 子串出现次数 */
	// 使用 Rabin-Karp 哈希算法实现
	fmt.Println(strings.Count("aba", "aba"))
	fmt.Println(strings.Count("汉字", "汉"))
	// 无重叠的次数
	fmt.Println(strings.Count("ababa", "aba"))
	// ""始终返回 utf8.RuneCountlnString(s) + 1
	fmt.Println(strings.Count("汉字", ""))
}
