package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 实现接口
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func area(g geometry) float64 {
	return g.area()
}

func perimeter(g geometry) float64 {
	return g.perimeter()
}

func main() {
	/* 接口是命名的方法签名（signatures）的集合 */
	r := rect{1, 2}
	c := circle{1}
	fmt.Println(area(r))
	fmt.Println(area(c))
	fmt.Println(perimeter(r))
	fmt.Println(perimeter(c))
}
