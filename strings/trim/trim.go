package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/* 字符串修剪 */
	s := "abc中文ABC"
	/* 连续，左侧右侧中匹配 cutset 任一字符删除 */
	// func Trim(s string, cutset string) string
	fmt.Println(strings.Trim(s, "acAC"))
	/* 连续，左侧中匹配 cutset 任一字符删除 */
	// func TrimLeft(s string, cutset string) string
	fmt.Println(strings.TrimLeft(s, "ac"))
	/* 连续，右侧中匹配 cutset 任一字符删除 */
	// func TrimRight(s string, cutset string) string
	fmt.Println(strings.TrimRight(s, "AC"))
	/* 连续，左侧中匹配前缀，删除 */
	// func TrimPrefix(s string, prefix string) string
	fmt.Println(strings.TrimPrefix(s, "ab"))
	/* 连续，右侧中匹配后缀，删除 */
	// func TrimSuffix(s string, prefix string) string
	fmt.Println(strings.TrimSuffix(s, "BC"))
	/* 连续，左侧右侧中间隔符，包括 \t \n \v \f \r ' ' U+0085 (NEL) */
	// func TrimSpace(s string) string
	s = "\v \nabc中文ABC\t \r \f"
	fmt.Println(strings.TrimSpace(s))
	s = "abc中文ABC"
	/* 连续，左侧右侧匹配 f 删除 */
	// func TrimFunc(s string, f func(rune) bool) string
	fmt.Println(strings.TrimFunc(s, func(r rune) bool {
		return !unicode.Is(unicode.Han, r)
	}))
	/* 连续，左侧匹配 f 删除 */
	// func TrimLeftFunc(s string, f func(rune) bool) string
	fmt.Println(strings.TrimLeftFunc(s, func(r rune) bool {
		return !unicode.Is(unicode.Han, r)
	}))
	/* 连续，右侧匹配 f 删除 */
	// func TrimRightFunc(s string, f func(rune) bool) string
	fmt.Println(strings.TrimRightFunc(s, func(r rune) bool {
		return !unicode.Is(unicode.Han, r)
	}))
}
