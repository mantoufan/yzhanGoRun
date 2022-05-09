package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/* 标题处理 */
	// 每单词首字母大写：Title(s string) string
	fmt.Println(strings.Title("article title"))
	// 全大写 ToTitle(s string) string
	fmt.Println(strings.ToTitle("article title"))
	// 全大写，特殊字符转对应特殊大写字母
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "ā"))
}
