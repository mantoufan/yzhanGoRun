package main

import (
	"fmt"
	"sort"
)

func main() {
	/* sort 包原生支持 []int []float64 []string 三种内建数据类型切片的排序操作 */
	/* 一、IntSlice 类型及 []int 排序 */
	/*
		type IntSlice []int
		func (p IntSlice) Len() int {
			return len(p)
		}
		func (p IntSlice) Less(i, j int) bool {
			return p[i] < p[j]
		}
		func (p IntSlice) Swap(i, j int) {
			p[i], p[j] = p[j], p[i]
		}
	*/
	// IntSlice 类型定义了 Sort() 方法，包装了 sort.Sort() 函数
	/*
		func (p IntSlice) Sort() {
			Sort(p)
		}
	*/
	// IntSlice 类型定义了 SearchInts() 方法，包装了 SearchInts() 函数
	/*
		func (p IntSlice) Search(x int) int {
			return SearchInts(p, x)
		}
	*/
	// sort.Ints() 方法使用了 IntSlice 类型
	/*
		func Ints(a []int) {
			Sort(IntSlice(a))
		}
	*/
	// 升序排列整数切片
	s := []int{4, 3, 1, 5, 2}
	sort.Ints(s)
	fmt.Println(s)
	// 降序排列整数切片
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	// 升序排列整数切片
	sort.Ints(s)
	fmt.Println(s)
	fmt.Println(sort.SearchInts(s, 2))
	// 判断是否升序
	fmt.Println(sort.IntsAreSorted(s))
	fmt.Println(sort.IsSorted(sort.IntSlice(s)))
	// 判断是否降序
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	fmt.Println(sort.IsSorted(sort.Reverse(sort.IntSlice(s))))
	/* 二、Float64Slice 类型及 []float64 排序 */
	/*
		type Float64Slice []float64
		func (p Float64Slice) Len() int {
			return len(p)
		}
		func (p Float64Slice) Less(i, j int) bool {
			return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j])
		}
		func (p Float64Slice) Swap(i, j int) {
			p[i], p[j] = p[j], p[i]
		}
		func (p Float64Slice) Sort() {
			Sort(p)
		}
		func (p Float64Slice) Search(x float64) int {
			return SearchFloat64(p, x)
		}
	*/
	// 升序排列
	// func Float64ss(a []float64)
	fs := []float64{.4, .3, .1, .5, .2}
	sort.Float64s(fs)
	fmt.Println(fs)
	// 降序排列
	sort.Sort(sort.Reverse(sort.Float64Slice(fs)))
	fmt.Println(fs)
	// 判断是否升序
	// func Float64AreSorted(a []float64) bool
	fmt.Println(sort.Float64sAreSorted(fs))
	fmt.Println(sort.IsSorted(sort.Float64Slice(fs)))
	// 判断是否降序
	fmt.Println(sort.IsSorted(sort.Reverse(sort.Float64Slice(fs))))
	// func SearchFloat64s(a []float64, x float64) int
	sort.Sort(sort.Float64Slice(fs))
	fmt.Println(fs)
	fmt.Println(sort.SearchFloat64s(fs, .3))

	/* 三、StringSlice 类型及 []string 排序 */
	/*
		  type StringSlice []string
			func (p StringSlice) Len() int {
				return len(p)
			}
			func (p StringSlice) Less(i, j int) bool {
				return p[i] < p[j]
			}
			func (p StringSlice) Swap(i, j int) {
				p[i], p[j] = p[j], p[i]
			}
			func (p StringSlice) Sort() {
				Sort(p)
			}
			func (p StringSlice) Search(x string) int {
				return SearchStrings(p, x)
			}
	*/
	// 升序排列
	// func Strings(a []string)
	ss := []string{"c", "b", "a"}
	sort.Strings(ss)
	fmt.Println(ss)
	// 降序排列
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Println(ss)
	// 是否升序排列
	// func StringsAreSorted(a []string) bool
	sort.Sort(sort.StringSlice(ss))
	fmt.Println(sort.StringsAreSorted(ss))
	// 是否降序排列
	fmt.Println(sort.IsSorted(sort.Reverse(sort.StringSlice(ss))))
	// func SearchStrings(a []string, x string) int
	fmt.Println(ss)
	fmt.Println(sort.SearchStrings(ss, "c"))
}
