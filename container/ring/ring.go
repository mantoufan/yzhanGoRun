package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(3)
	r2 := ring.New(3)

	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 4; i <= 6; i++ {
		r2.Value = i
		r2 = r2.Prev()
	}

	// 连接两个环
	// func (r *Ring) Link(s *Ring) *Ring
	r.Link(r2)

	sum := 0
	// 循环进行操作
	// func (r *Ring) Do(f func(inerface{}))
	r.Do(func(p interface{}) {
		sum += p.(int)
	})
	fmt.Println("sum = ", sum)

	// 环长度
	// func (r *Ring) Len() int
	fmt.Println("len", r.Len())

	// 从当前元素开始，删除 n 个元素
	// func (r *Ring) Unlink(n int) *Ring
	fmt.Println(r.Next().Value)
	fmt.Println(r.Prev().Value)

	r.Unlink(3)
	r.Do(func(p interface{}) {
		fmt.Println("do：", p.(int))
	})

	// 指针从当前元素向后移动或向前移动
	// func (r *Ring) Move(n int) *Ring
	r.Move(1)
	r.Do(func(p interface{}) {
		fmt.Println("move：", p.(int))
	})
}
