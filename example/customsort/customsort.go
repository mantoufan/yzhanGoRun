package main

import (
	"fmt"
	"sort"
)

type ByDescendingLength []string

func (s ByDescendingLength) Len() int {
	return len(s)
}

func (s ByDescendingLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByDescendingLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func main() {
	/* 实现 sort.Interface 的 Len Less 和 Swap 方法
	   按字符串长度降序排列 */
	strs := []string{"a", "bb", "ccc"}
	sort.Sort(ByDescendingLength(strs))
	fmt.Println(strs)
}
