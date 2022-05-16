package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* 字符串转浮点数 */
	/* 字符串 → 浮点数 */
	// func ParseFloat(s string, bitSize int) (f float64, err error)
	f0, e0 := strconv.ParseFloat("1.23", 0)
	fmt.Println(f0, e0)
	f1, e1 := strconv.ParseFloat("1.23a", 0)
	fmt.Println(f1, e1)
	f2, e2 := strconv.ParseFloat("1e-1024", 2)
	fmt.Println(f2, e2)
	/* 浮点数 → 字符串 */
	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	fmt.Println(strconv.FormatFloat(1.23, 'e', 1, 32))
	fmt.Println(strconv.FormatFloat(1.23, 'E', 1, 32))
	fmt.Println(strconv.FormatFloat(1.23, 'g', 1, 32))
	fmt.Println(strconv.FormatFloat(1.23, 'G', 1, 32))
	s := strconv.FormatFloat(1.23, 'e', 1, 32)
	fmt.Println(s)
	fmt.Println(strconv.ParseFloat(s, 32))
	/* 向 dst 添加 float64 */
	// func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int)
	dst := []byte("")
	dst = strconv.AppendFloat(dst, 1.23, 'e', 1, 32)
	fmt.Print(string(dst))
}
