package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/* 是否符合 utf-8 编码 */
	// func Valid(p []byte) bool
	fmt.Println(utf8.Valid([]byte("中")[:2]))
	// func VaildRune(r rune) bool
	fmt.Println(utf8.ValidRune('中'))
	// func VaildString(s string) bool
	fmt.Println(utf8.ValidString("中"))

	/* 判断 Rune 所占字节数 */
	// func RuneLen(r rune) int
	b := make([]byte, 3)
	fmt.Println(utf8.RuneLen('中'))

	/* 判断字节串或字符串的 Rune 数 */
	// func RuneCount(p []byte) int
	fmt.Println(utf8.RuneCount([]byte("中")))
	fmt.Println(utf8.RuneCountInString("中"))

	/* 编码和解码到 Rune */
	// func EncodeRune(p []byte, r rune) int
	b = make([]byte, 3)
	fmt.Println(utf8.EncodeRune(b, '中'))
	// func DecodeRune(p []byte)(r rune, size int)
	fmt.Println(utf8.DecodeRune(b))
	// func DecodeRuneInSting(s string)(r rune, size int)
	fmt.Println(utf8.DecodeRuneInString("中"))
	// func DecodeLastRune(p []byte)(r rune, size int)
	fmt.Println(utf8.DecodeLastRune(b))
	// func DecodeLastRuneInString(s string)(r rune, size int)
	fmt.Println(utf8.DecodeLastRuneInString("中"))

	/* 是否为完整 rune */
	// func FullRune(p []byte) bool
	fmt.Println(utf8.FullRune(b))
	fmt.Println(utf8.FullRune(b[:1]))
	// func FullRuneInString(s string) bool
	fmt.Println(utf8.FullRuneInString("中"))

	/* 是否为 rune 的第一个字节 */
	// func RuneStart(b byte) bool
	fmt.Println(utf8.RuneStart(b[0]))
	fmt.Println(utf8.RuneStart(b[1]))
}
