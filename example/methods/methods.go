package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return (r.width + r.height) << 1
}

func (r *rect) setX2() {
	r.width *= 2
	r.height *= 2
}
func main() {
	/* 方法：结构体定义函数 */
	r := rect{1, 2}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
	// 指针调用：
	// 1 避免方法调用时产生拷贝
	// 2 方法可以改变结构体
	rPointer := &r
	fmt.Println(rPointer.area())
	fmt.Println(rPointer.perimeter())

	// 修改指针：
	r.setX2()
	fmt.Println(rPointer.area())
	fmt.Println(rPointer.perimeter())
}
