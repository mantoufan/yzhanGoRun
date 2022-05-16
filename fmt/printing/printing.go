package main

import (
	"fmt"
)

func main() {
	type class struct {
		teacher string
		student string
	}
	c := class{"teacher", "student"}
	/* Printf 格式化输出 */
	// 格式化输出结构 struct
	fmt.Printf("% + v\n", c)
	// 输出值的 Go 语言表示方法
	fmt.Printf("%#v\n", c)
	// 输出值得类型得 Go 语言表述方法
	fmt.Printf("%T\n", c)
	// 输出值的 true 或 false
	fmt.Printf("%t\n", true)
	fmt.Printf("%t\n", false)
	// 输出二进制表示
	fmt.Printf("%b\n", 1024)
	// 输出 Unicode 编码字符
	fmt.Printf("%c\n", 97)
	// 输出十进制表示
	fmt.Printf("%d\n", 0b10000000000)
	// 输出八进制表示
	fmt.Printf("%o\n", 8)
	// 输出十六进制 + 单引号
	fmt.Printf("%q\n", 15)
	// 输出十六进制 · 小写
	fmt.Printf("%x\n", 15)
	// 输出十六进制 · 大写
	fmt.Printf("%X\n", 15)
	// 输出 Unicode 表示
	fmt.Printf("%U\n", 97)
	// 无小数部分 · 指数的科学计数法
	fmt.Printf("%b\n", 12.34)
	// 科学计数法 · e 表示
	fmt.Printf("%e\n", 12.34)
	// 科学计数法 · E 表示
	fmt.Printf("%E\n", 12.34)
	// 有小数部分，无指数部分
	fmt.Printf("%f\n", 12.34)
	// 有小数部分，根据 e 或 f 表示
	fmt.Printf("%g\n", 12.34)
	// 有小数部分，根据 E 或 f 表示
	fmt.Printf("%G\n", 12.34)
	// 输出字符串 或 []byte
	fmt.Printf("%s\n", "abc")
	// 输出字符串 + 双引号
	fmt.Printf("%q\n", "abc")
	// 输出字符串，两字节十六进制表示，a-f 表示
	fmt.Printf("%x\n", "a")
	// 输出字符串，两字节十六进制表示，A-F 表示
	fmt.Printf("%X\n", "a")
	// 输出 0x 开头的十六进制表示
	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
}
