package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* 整型转字符串 */
	// func Itoa(i int) string
	fmt.Println(strconv.Itoa(-15))
	// func FormatInt(i int64, base int) string
	// Itoa = FormatInt(i, 10)
	fmt.Println(strconv.FormatInt(-15, 10))
	fmt.Println(strconv.FormatInt(-15, 2))
	fmt.Println(strconv.FormatInt(-15, 16))
	fmt.Println(strconv.FormatInt(-35, 36))
	// func FormatUint(1 uint64, base int) string
	fmt.Println(strconv.FormatUint(15, 10))
	fmt.Println(strconv.FormatUint(15, 2))
	fmt.Println(strconv.FormatUint(15, 16))
	fmt.Println(strconv.FormatUint(35, 36))
}
