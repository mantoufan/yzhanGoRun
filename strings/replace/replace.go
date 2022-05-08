package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 替换 */
	// func Replace(s, pattern, replacement string, n int) string
	// func ReplaceAll(s, pattern, replacement string) string
	// ReplaceAll 内部调用了 Replace(s, pattern, replacement, -1)
	fmt.Println(strings.Replace("a a a", "a", "b", 1))
	fmt.Println(strings.Replace("a a a", "a", "b", -1))
	fmt.Println(strings.ReplaceAll("a a a", "a", "b"))
}
