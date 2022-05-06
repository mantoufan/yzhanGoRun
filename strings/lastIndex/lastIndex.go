package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 字符或子串在字符串中出现位置 */
	/* Index */
	// 子串首次出现位置
	fmt.Println(strings.Index("aba", "a"))
	// 字节首次出现位置
	fmt.Println(strings.IndexByte("aba", 'a'))
	// 子串中任意 Unicode 代码点首次出现位置
	fmt.Println(strings.IndexAny("aba", "ade"))
	fmt.Println(strings.IndexAny("熊大熊二", "小熊"))
	// Unicode 代码点首次出现位置
	fmt.Println(strings.IndexRune("熊大熊二", '熊'))
	// 回调函数首次返回位置
	fmt.Println(strings.IndexFunc("aba熊大熊二", func(char rune) bool {
		if char == '熊' {
			return true
		}
		return false
	}))

	/* LastIndex */
	// 子串最后一次出现位置
	fmt.Println(strings.LastIndex("aba", "a"))
	// 字节首次出现位置
	fmt.Println(strings.LastIndexByte("aba", 'a'))
	// 子串中任意 Unicode 代码点首次出现位置
	fmt.Println(strings.LastIndexAny("aba", "ade"))
	fmt.Println(strings.LastIndexAny("熊大熊二", "小熊"))
	// 回调函数首次返回位置
	fmt.Println(strings.LastIndexFunc("aba熊大熊二", func(char rune) bool {
		if char == '熊' {
			return true
		}
		return false
	}))
}
