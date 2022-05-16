package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* 字符串转布尔值 */
	/* 字符串 → 布尔值 */
	// func ParseBool(str string) (value bool, err error)
	b0, e0 := strconv.ParseBool("true")
	fmt.Println(b0, e0)
	b1, e1 := strconv.ParseBool("string")
	fmt.Println(b1, e1)
	/* 布尔值 → 字符串 */
	// func FormatBool(b bool) string
	b2 := strconv.FormatBool(true)
	fmt.Println(b2)
	b3 := strconv.FormatBool(false)
	fmt.Println(b3)
	/* 向 dst 添加字符串 */
	// func AppendBool(dst []byte, b bool)
	dst := []byte("Is Bool:")
	dst = strconv.AppendBool(dst, true)
	fmt.Println(string(dst))
}
