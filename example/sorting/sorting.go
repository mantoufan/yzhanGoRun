package main

import (
	"fmt"
	"sort"
)

func main() {
	/* 原地排序 */
	// 整数排序：升序
	ints := []int{3, 1, 2}
	sort.Ints(ints)
	fmt.Println("Ints：", ints)

	// 整数排序：升序：是否排序
	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted：", sorted)
	sorted = sort.IsSorted(sort.IntSlice(ints))
	fmt.Println("Sorted：", sorted)

	// 整数排序：降序
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Ints Reverse：", ints)

	// 整数排序：降序：是否排序
	sort.IsSorted(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Ints Sorted：", sorted)
}
