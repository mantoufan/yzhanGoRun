package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*
		  Reader 类型：用于 Read 数据，实现 io 包下：
		  Reader, ReaderAt, RuneReader,
			RuneScanner, ByteReader, ByteScanner,
			ReadSeeker, Seeker, WriterTo 多个接口
	*/
	// 读取数据至 b，改变进度下标
	// func (r *Reader) Read(b []byte) (n int, err error)
	r := bytes.NewReader([]byte("abc中文"))
	b := make([]byte, 3)
	r.Read(b)
	fmt.Println(string(b))
	r.Read(b)
	fmt.Println(string(b))
	// 重新写入数据 b
	// func (r *Reader) Reset(b []byte)
	r.Reset([]byte("abc"))
	r.Read(b)
	fmt.Println(string(b))
	// 读取 1 个字节，改变进度下标
	// func (r *Reader) ReadByte() (byte, error)
	r.Reset([]byte("abc中文"))
	char, e := r.ReadByte()
	fmt.Println(string(char), e)
	// 返回前 1 个字节，已到达 0 返回错误，改变进度下标
	e = r.UnreadByte()
	fmt.Println(e)
	e = r.UnreadByte()
	fmt.Println(e) // 已到头，报错
	r.Reset([]byte("中文"))
	// 读取 1 个字符，改变进度下标
	// func (r *Reader) ReadRune() (ch rune, size int, err error)
	ch, size, e := r.ReadRune()
	fmt.Println(string(ch), size, e)
	// 返回前 1 个字符，已到达 0 返回错误，改变进度下标
	e = r.UnreadRune()
	fmt.Println(e)
	e = r.UnreadRune()
	fmt.Println(e)
	// 读取数据至 Writer
	// func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
	w := &bytes.Buffer{}
	r.WriteTo(w)
	fmt.Println(w.String())
	// 读取 off 开始的数据至 b，不改变进度下标
	// func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
	b = make([]byte, 3)
	r.ReadAt(b, 3)
	fmt.Println(string(b))
	// 根据 whence 值，修改并返回进度下标
	// func (r *Reader) Seek(offset int64, whence int) (int64, error)
	// whence == 0，进度下标改为 offset
	fmt.Println(r.Seek(1, 0)) // to 1
	// whence == 1，进度下标改为 当前进度 + offset
	fmt.Println(r.Seek(2, 1)) // to 1 + 2
	// whence == 2，进度下标改为 长度 + offset
	fmt.Println(r.Seek(-3, 2)) // to 6 - 3
}
