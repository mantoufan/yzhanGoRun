package main

import (
	"container/list"
	"fmt"
)

type Element struct {
	next, prev *Element    // 上一个元素和下一个元素
	list       *List       // 元素所在链表
	Value      interface{} // 元素
}
type List struct {
	root Element // 链表的根元素
	len  int     // 链表的长度
}

func main() {
	/* 链表是一个有 prev 和 next 指针的数组 */
	// func New() *List
	l := list.New()
	// func (l *List) PushBack(v interface{}) *Element // 在链表最后插入元素
	l.PushBack(2)
	// func (l *List) PushFront(v interface{}) *Element // 在链表头部插入元素
	l.PushFront(1)

	fmt.Printf("len: %v\n", l.Len())
	fmt.Printf("first: %v\n", l.Front().Value)
	fmt.Printf("second: %v\n", l.Front().Next().Value)

	l1 := list.New()
	l1.PushBack(3)
	l1.PushBack(4)
	// func (l *List) PushBackList(other *List) // 在链表最后插入接上新链表
	l.PushBackList(l1)

	fmt.Printf("len: %v\n", l.Len())
	fmt.Printf("third: %v\n", l.Front().Next().Next().Value)
	fmt.Printf("fourth: %v\n", l.Front().Next().Next().Next().Value)

	l2 := list.New()
	l2.PushBack(0)
	// func (l *List) PushFrontList(other *List) // 在链表头部插入接上新链表
	l.PushFrontList(l2)

	fmt.Printf("len: %v\n", l.Len())
	fmt.Printf("zero: %v\n", l.Front().Value)

	// 在某个元素后插入
	// func (l *List) InsertAfter(v interface{}, mark *Element) *Element // 在某个元素后插入
	l.InsertAfter(6, l.Back())
	fmt.Printf("sixth: %v\n", l.Back().Value)

	// 在某个元素前插入
	// func (l *List) InsertBefore(v interface{}, mark *Element) *Element // 在某个元素前插入
	l.InsertBefore(5, l.Back())
	fmt.Printf("fifth：%v\n", l.Back().Prev().Value)

	// 把 e 元素移动到 mark 之后
	// func (l *List) MoveAfter(e, mark *Element)
	l.MoveAfter(l.Front(), l.Back())
	fmt.Printf("zero -> sisth：%v\n", l.Back().Value)

	// 把 e 元素移动到 mark 之前
	// func (l *List) MoveBefore(e, mark *Element)
	l.MoveBefore(l.Back(), l.Front())
	fmt.Printf("sisth -> zero：%v\n", l.Front().Value)

	// 把 e 元素移动到链表最后
	// func (l *List) MoveToBack(e *Element)
	l.MoveToBack(l.Front())
	fmt.Printf("zero -> sisth：%v\n", l.Back().Value)

	// 把 e 元素移动到链表头部
	// func (l *List) MoveToFront(e *Element)
	l.MoveToFront(l.Back())
	fmt.Printf("sisth -> zero：%v\n", l.Front().Value)

	// 删除某个元素怒
	// func (l *List) Remove(e *Element) interface{}
	fmt.Printf("len：%v\n", l.Len())
	l.Remove(l.Front())
	fmt.Printf("len：%v\n", l.Len())

	// 链表初始化
	// func (l *List) Init() *List
	l.Init()
	fmt.Printf("len：%v\n", l.Len())
}
