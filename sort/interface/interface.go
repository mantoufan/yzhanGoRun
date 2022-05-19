package main

import (
	"fmt"
	"sort"
)

/* Interface */
// 对数据集合（包括自定义数据类型的集合）排序
// 需要实现 sort.Interface 接口
type Interface interface {
	// 获取数据集合元素个数
	Len() int
	// 如果 i 索引的数据 < j 索引的数据，返回 true 且不会调用 Swap()
	// 数据升序排序
	Less(i, j int) bool
	// 交换 i 和 j 索引两个元素位置
	Swap(i, j int)
}

// 排序：func Sort(data Interface)
// 是否排序：func IsSorted(data Interface) bool
func IsSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// 自定义 struct 排序
type StuScore struct {
	name  string // 姓名
	score int    // 分数
}

type StuScores []StuScore

// Len()
func (s StuScores) Len() int {
	return len(s)
}

// Less()
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"a", 90},
		{"b", 99},
		{"c", 95},
	}
	fmt.Println(stus)
	sort.Sort(stus)
	fmt.Println(sort.IsSorted(stus))
	fmt.Println(stus)
}
