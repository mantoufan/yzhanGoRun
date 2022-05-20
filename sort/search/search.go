package main

import (
	"fmt"
	"sort"
)

func main() {
	/* 二分查找
	  func Search(n int, f func(int) bool) int
	  用“二分查找”算法查找：
		使 f(x)(0 <= x < n) 返回 true 的最小值 i
		不存在 i，返回 n
	*/
	pivot, a := 3, []int{0, 1, 2, 3, 4, 5}
	index := sort.Search(len(a), func(i int) bool {
		return a[i] >= pivot
	})
	fmt.Println(index)
	/* 猜数字 */
	GuessingGame()
}
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d?", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
