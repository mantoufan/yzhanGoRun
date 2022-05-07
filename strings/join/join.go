package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	/* 字符串的连接：JOIN */
	// Copy
	b := make([]byte, 2)
	bp := copy(b, "a")
	fmt.Println(bp, string(b))
	bp += copy(b[bp:], "b")
	fmt.Println(bp, string(b))
	// Join By BufferString
	JoinByBuffer := func(str []string, sep string) string {
		if len(str) == 0 {
			return ""
		}
		if len(str) == 1 {
			return str[0]
		}
		buffer := bytes.NewBufferString(str[0])
		for _, s := range str[1:] {
			buffer.WriteString(sep)
			buffer.WriteString(s)
		}
		return buffer.String()
	}
	fmt.Println(JoinByBuffer([]string{"你", "好"}, "们"))
	// Join By Copy
	JoinByCopy := func(str []string, sep string) string {
		if len(str) == 0 {
			return ""
		}
		if len(str) == 1 {
			return str[0]
		}
		n := len(sep) * (len(str) - 1)
		for i := 0; i < len(str); i++ {
			n += len(str[i])
		}
		b := make([]byte, n)
		bp := copy(b, str[0])
		for _, s := range str[1:] {
			bp += copy(b[bp:], sep)
			bp += copy(b[bp:], s)
		}
		return string(b)
	}
	fmt.Println(JoinByCopy([]string{"你", "好"}, "们"))
	// Join Provide By strings
	fmt.Println(strings.Join([]string{"你", "好"}, "们"))
}
