package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/* 大小写转换 */
	// func ToLower(s string) string
	fmt.Println(strings.ToLower("A"))
	// func ToLowerSpecial(c unicode.SpecialCase, s string) string
	fmt.Println(strings.ToLowerSpecial(unicode.AzeriCase, "Ā"))
	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("a"))
	// func ToUpperSpecial(c unicode.SpecialCase, s string) string
	fmt.Println(strings.ToUpperSpecial(unicode.AzeriCase, "ā"))
}
