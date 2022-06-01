package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	/* 结构体：字段集合（带类型） 组织数据 */
	// 创建
	fmt.Println(student{"Anne", 16})
	// 创建并指定字段名字
	fmt.Println(student{name: "Anne", age: 16})
	// 省略字段被初始化为零值
	fmt.Println(student{name: "Anne"})
	st := student{age: 16}
	fmt.Println(st.name == "")
	// & 前缀生成结构体指针
	fmt.Println(&student{"Anne", 16})
	// . 访问字段
	s := student{"Anne", 16}
	fmt.Println(s.name)
	fmt.Println(s.age)
	// 结构体指针
	sPointer := &s
	sCopy := s
	// . 访问字段，指针会自动解引用
	fmt.Println(sPointer.name)
	fmt.Println(sPointer.age)
	// 修改结构体
	sPointer.name = "Bob"
	s.age = 26
	// 拷贝指针的结果
	fmt.Println(sPointer)
	fmt.Println(s)
	// 拷贝的结果
	fmt.Println(sCopy.name)
	fmt.Println(sCopy.age)
}
