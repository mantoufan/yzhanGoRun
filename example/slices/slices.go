package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	/* 切片：可变长度的复合序列 */
	/*
		  slice 声明时，不指定 len / cap，则 len / cap 值默认声明时的数组长度为准
			len：长度，slice 包含元素个数，可变
			cap：容量，slice 当前可接纳元素个数，根据 len 动态变化
	*/
	// 声明 a
	var a []int
	fmt.Printf("%d %d %v\n", len(a), cap(a), a)

	// 声明初始化 a
	a = []int{0}
	fmt.Printf("%d %d %v\n", len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("%d %d %v\n", len(a), cap(a), a)

	// 声明 b ：类型 + len + cap
	b := make([]int, 1, 2)
	fmt.Printf("%d %d %v\n", len(b), cap(b), b)

	// copy(to, from) 第一参数：目标 slice 第二参数：源 slice
	f, t := []int{0, 1}, []int{2, 3, 4, 5}
	copy(f, t)
	fmt.Println(f)

	// 比较 slice
	c, d := []int{0, 1}, []int{0, 1}
	fmt.Println(c, d, sliceEqual(c, d))
	c, d = []int{0, 1}, []int{0, 2}
	fmt.Println(c, d, sliceEqual(c, d))

	// 是否为空 slice
	e, d := []int{}, nil
	fmt.Println(sliceIsEmpty(e))
	fmt.Println(sliceIsEmpty(d))

	// 插入 slice
	g := []int{0, 1, 3}
	fmt.Println(sliceInsert(g, 2, 2))

	// 删除 slice
	fmt.Println(sliceRmove(g, 1))

	// 翻转 slice
	h := []int{2, 1, 0}
	fmt.Println(sliceReverse(h))

	// 字符串 slice
	fmt.Println(intsToString([]int{0, 1, 2}))
}

// 比较 slice 是否相等
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// 是否为空 slice
func sliceIsEmpty(a []int) bool {
	return len(a) == 0 || a == nil
}

// 插入 slice：在指定索引，插入元素
func sliceInsert(slice []int, i int, n int) []int {
	// 将元素插入指定位置
	sub := []int{n}
	for _, e := range slice[i:] {
		sub = append(sub, e)
	}
	copy(slice[i:], sub)
	slice = append(slice, sub[len(sub)-1])
	return slice
}

// 删除 slice：删除指定索引的元素
func sliceRmove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 翻转 slice
func sliceReverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Slice 转 字符串
func intsToString(slice []int) string {
	var buf bytes.Buffer
	for _, b := range slice {
		buf.WriteString(strconv.Itoa(b))
	}
	return buf.String()
}
