package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* 字符串转整型 */
	// func Atoi(s string) (i int, err error)
	fmt.Println(strconv.Atoi("1"))
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// bitSize 整数取值范围：0 8 16 32 64 分别表示 int(32 位 4 字节，64 位 8 字节) int8 int16 int32 int64
	// Atoi = ParseInt(s, 10, 0)
	fmt.Println(strconv.ParseInt("-0xf", 0, 0))
	fmt.Println(strconv.ParseInt("-f", 16, 0))
	fmt.Println(strconv.ParseInt("-0o11", 0, 0))
	fmt.Println(strconv.ParseInt("-11", 8, 0))
	fmt.Println(strconv.ParseInt("-0b11", 0, 0))
	fmt.Println(strconv.ParseInt("-11", 2, 0))
	// func ParseUint(s string, base int, bitSize int) (i int64, err error)
	fmt.Println(strconv.ParseUint("0xf", 0, 0))
	fmt.Println(strconv.ParseUint("f", 16, 0))
	fmt.Println(strconv.ParseUint("0o11", 0, 0))
	fmt.Println(strconv.ParseUint("11", 8, 0))
	fmt.Println(strconv.ParseUint("0b11", 0, 0))
	fmt.Println(strconv.ParseUint("11", 2, 0))
}
