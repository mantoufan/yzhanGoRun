package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 1, 3, 2, 4}
	/* 升序排列 */
	sort.Ints(nums)
	fmt.Println(nums)
	/* 降序排列 */
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	/* 原理 */
	// 定义了一个 reverse 结构类型，嵌入 Interface 接口
	/*
		type reverse struct {
			Interface
		}
	*/
	// reverse 结构类型的 Less() 方法拥有嵌入的 Less() 方法相反的行为
	// Len() 和 Swap() 方法则会保持嵌入类型的方法
	/*
		func (r reverse) Less(i, j int) bool {
			return r.Interface.Less(j, i)
		}
	*/
	// 返回新的实现 Interface 接口的数据类型
	/*
		func Reverse(data Interface) Interface {
			return &reverse{data}
		}
	*/
}
