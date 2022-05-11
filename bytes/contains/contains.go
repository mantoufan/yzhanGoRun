package main

import (
	"bytes"
	"fmt"
)

func main() {
	/* 是否存在某个子 slice */
	// func Contains(b, subslice []byte) bool
	fmt.Println(bytes.Contains([]byte("abc"), []byte("a")))
	// 内部实现
	Contains := func(b, subslice []byte) bool {
		return bytes.Index(b, subslice) != -1
	}
	fmt.Println(Contains([]byte("abc"), []byte("a")))
}
