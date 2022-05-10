package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	/* 读取字符串 */
	// Len(): 返回未读长度
	// Size()：返回整体长度
	s := strings.NewReader("abc")
	fmt.Println(s.Len())
	fmt.Println(s.Size())
	// 读取 2 个字符
	buf := make([]byte, 2)
	s.Read(buf)
	fmt.Println(string(buf))
	// Len(): 返回未读长度
	fmt.Println(s.Len())
	// Seek: 指针指向下次读取起始位置
	// func (r *Reader) Seek(offset int64, whence int) (int64, error) {}
	// whence = io.SeekStart 从起始位置往后
	// whence = io.SeekCurrent 从当前位置往后
	// whence = io.SeekEnd 从结束位置往前
	fmt.Println(s.Seek(0, io.SeekStart))
	fmt.Println(s.Seek(1, io.SeekCurrent))
	fmt.Println(s.Seek(-1, io.SeekEnd))
	// ReadAt 指针不动，从指定位置读取
	// func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {}
	buf = make([]byte, 1)
	s.ReadAt(buf, 2)
	fmt.Println(string(buf))
	// ReadByte() 指针 + 1，向后读 1 byte
	buf2, _ := s.ReadByte()
	fmt.Println(string(buf2))
	fmt.Println(s.Len())
	// UnreadByte() 指针 -1
	s.UnreadByte()
	fmt.Println(s.Len())
	// ReadRune 读中文
	// func (r *Reader) ReadRune() (ch rune, size int, err error) {}
	s = strings.NewReader("a中文")
	ch, size, _ := s.ReadRune()
	fmt.Println(string(ch))
	fmt.Println(size)
	ch, size, _ = s.ReadRune()
	fmt.Println(string(ch))
	fmt.Println(size)
	fmt.Println(s.Len())
	// UnreadRune() 指针 -字符长度：中文 3 非中文 1
	s.UnreadRune()
	fmt.Println(s.Len())
}
