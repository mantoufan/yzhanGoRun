package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* 遍历 */
	// 遍历 slice
	nums, sum, is := []int{0, 1, 2}, 0, []int{}
	for i, num := range nums {
		sum += num
		is = append(is, i)
	}
	fmt.Println(sum, is)

	// 遍历 map
	m := map[int]int{0: 0, 1: 1, 2: 2}
	sum, is = 0, []int{}
	for i, num := range m {
		sum += num
		is = append(is, i)
	}
	fmt.Println(sum, is)

	// 遍历 string
	sum, is = 0, []int{}
	for i, c := range "012" {
		n, _ := strconv.Atoi(string(c))
		sum += n
		is = append(is, i)
	}
	fmt.Println(sum, is)
}
