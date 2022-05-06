package main

import (
	"fmt"
	"strings"
)

func main() {
	// hasPrefix: 是否有前缀
	MyHasPrefix := func(s string, prefix string) bool {
		return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
	}
	fmt.Println(MyHasPrefix("abc", "a"))
	fmt.Println(MyHasPrefix("abc", ""))
	fmt.Println(MyHasPrefix("abc", "b"))
	fmt.Println(strings.HasPrefix("abc", "a"))
	fmt.Println(strings.HasPrefix("abc", ""))
	fmt.Println(strings.HasPrefix("abc", "b"))
	// hasSuffix: 是否有后缀
	MyHasSuffix := func(s string, suffix string) bool {
		return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
	}
	fmt.Println(MyHasSuffix("abc", "c"))
	fmt.Println(MyHasSuffix("abc", ""))
	fmt.Println(MyHasSuffix("abc", "b"))
	fmt.Println(strings.HasSuffix("abc", "c"))
	fmt.Println(strings.HasSuffix("abc", ""))
	fmt.Println(strings.HasSuffix("abc", "b"))
}
