package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 写入字符串 */
	b := strings.Builder{}
	// 写入 1 字节
	// func (b *Builder) WriteByte(c byte) error
	fmt.Println(b.WriteByte('a'))
	fmt.Println(b.String())
	// 写入 多 字符
	// func (b *Builder) Write(p []byte) (int, error)
	fmt.Println(b.Write([]byte("bc")))
	fmt.Println(b.String())
	// 写入 1 字符
	// func (b *Builder) WriteRune(r rune) (int, error)
	fmt.Println(b.WriteRune('中'))
	fmt.Println(b.String())
	// Len() 字符串长度
	fmt.Println(b.Len())
	// Cap() 返回 current_capacity
	fmt.Println(b.Cap())
	// Grow(n int) Cap 增加：current_capacity * 2 + n
	// n <= 8 - 6 不需要扩容
	b.Grow(2) // 8
	fmt.Println(b.Cap())
	// n > 8 - 6 需要扩容：8 * 2 + 3
	b.Grow(3)
	fmt.Println(b.Cap())
	// Rest() 清空
	b.Reset()
	fmt.Println(b.String())
}
