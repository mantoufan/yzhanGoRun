package main

import (
	"fmt"
	"strings"
)

func main() {
	/* Contains */
	fmt.Println(strings.Contains("abc", "a"))
	/* Contains Any */
	// 只要第二参数中任意一个字符被包含，即返回 true
	fmt.Println(strings.ContainsAny("abc", "ade"))
	fmt.Println(strings.ContainsAny("abc", "def"))
	/* Contains Rune */
	// 判断代码点是否存在
	fmt.Println(strings.ContainsRune("abc", rune('a')))
	fmt.Println(strings.ContainsRune("abc", rune('d')))

	/* Contains 内部使用 Index 实现，判断 Index 返回值是否 >= 0 */
	fmt.Println(strings.Index("abc", "a"))
	fmt.Println(strings.IndexAny("abc", "ade"))
	fmt.Println(strings.IndexByte("abc", 'a'))
	fmt.Println(strings.IndexRune("中文", '中'))
	fmt.Println(strings.IndexFunc("中文", func(r rune) bool {
		if r == '中' {
			return true
		}
		return false
	}))
	fmt.Println(strings.LastIndex("abc", "a"))
	fmt.Println(strings.LastIndexAny("abc", "ade"))
	fmt.Println(strings.LastIndexByte("abc", 'a'))
	fmt.Println(strings.LastIndexFunc("中文", func(r rune) bool {
		if r == '中' {
			return true
		}
		return false
	}))
}
