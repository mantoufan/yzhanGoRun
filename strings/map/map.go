package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/* 替换：遍历字符，调用回调函数，将返回值替换原字符 */
	// func Map(mapping func(rune) rune, s string) string
	// 返回 -1 丢弃字符
	fmt.Println(strings.Map(func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // 大写保留
			return r
		case r >= 'a' && r <= 'z': // 小写转大写
			return r - 32
		case unicode.Is(unicode.Han, r): // 汉字换成空格
			return ' '
		}
		return -1 // 其它字符丢弃
	}, "abcABC\\汉字."))

}
