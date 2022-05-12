package main

import (
	"bytes"
	"fmt"
)

func main() {
	/* 读写操作 */
	b := bytes.NewBufferString("abc中文")
	// 返回字节数组：[offset, delim || 末尾]，改进度指针
	// func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
	frontBytes, _ := b.ReadBytes('b')
	fmt.Println(string(frontBytes))
	// 返回字符串：[offset, delim || 末尾]，改进度指针
	frontString, _ := b.ReadString('c')
	fmt.Println(frontString)
	// n = 0 重制
	// n > 0 舍弃 offset + n 后数据
	// func (b *Buffer) Truncate(n int)
	fmt.Println(b.Len())
	b.Truncate(0)
	fmt.Println(b.Len())
	b.WriteString("abc中文")
	b.ReadByte()
	fmt.Println(b.String())
	b.Truncate(2)
	fmt.Println(b.String())
}
