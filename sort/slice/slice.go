package main

import (
	"fmt"
	"sort"
)

func main() {
	/* 排序与查找 */
	/* 传入比较函数，比较两个变量的大小 */
	// func Slice(slice interface{}, less func(i, j int) bool)
	students := []struct {
		Name string
		Age  int
	}{
		{"a", 3},
		{"b", 1},
		{"c", 1},
		{"d", 2},
	}
	sort.Slice(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	})
	fmt.Println(students)
	// 稳定排序
	// func SliceStable(slice interface{}, less func(i, j int) bool)
	students = []struct {
		Name string
		Age  int
	}{
		{"a", 3},
		{"b", 1},
		{"c", 1},
		{"d", 2},
	}
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	})
	fmt.Println(students)
	// 判断是否有序
	// func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
	fmt.Println(sort.SliceIsSorted(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	}))
	// 二分查找指定元素
	// func Search(n int, f func(int) bool) int
	index := sort.Search(len(students), func(i int) bool {
		return students[i].Age >= 1 // 升序，第一个满足条件最小索引
	})
	fmt.Println(index)
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Age > students[j].Age
	})
	fmt.Println(students)
	index = sort.Search(len(students), func(i int) bool {
		return students[i].Age <= 1 // 降序，第一个满足条件的最小索引
	})
	fmt.Println(index)
}
