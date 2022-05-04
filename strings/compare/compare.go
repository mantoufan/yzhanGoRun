package main

import (
	"fmt"
	"strings"
)

func main() {
	a, b := "A", "a"
	// 不忽略大小写: < -1 = 0 > 1
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(strings.Compare(a, b))
	// 忽略大小写
	fmt.Println(strings.EqualFold(a, b))
}
