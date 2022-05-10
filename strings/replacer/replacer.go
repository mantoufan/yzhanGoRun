package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 结构：替换多组字符串对（old-new) */
	r := strings.NewReplacer("a", "b", "A", "B")
	fmt.Println(r.Replace("aA"))
}
