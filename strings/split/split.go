package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/* Fields */
	// 按空格或连续空格分隔
	slice := strings.Fields(" foo bar baz ")
	fmt.Println(slice)
	/* FieldsFunc */
	// 回调函数分隔字符串
	slice = strings.FieldsFunc(" foo bar  baz", func(char rune) bool {
		return unicode.IsSpace(char)
	})
	fmt.Println(slice)
	/* Split */
	// 按 sep 分割
	slice = strings.Split("foo,bar,baz", ",")
	fmt.Println(slice)
	/* Split */
	// 按 sep 分割，保留 sep
	slice = strings.SplitAfter("foo,bar,baz", ",")
	fmt.Println(slice)
	/* SplitN */
	// 按 sep 分割 N 次
	slice = strings.SplitN("foo,bar,bza", ",", 2)
	fmt.Println(slice)
	/* SplitN */
	// 按 sep 分割 N 次，保留 sep
	slice = strings.SplitAfterN("foo,bar,bza", ",", 2)
	fmt.Println(slice)
}
