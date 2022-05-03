package main

import "fmt"

func main() {
	/*
		3.1 变量
		var a int // 没有赋值，默认为 0
		var a int  = 1 // 声明时赋值
		var a = 1 // 声明时赋值 */
	/*
		a := 1
		msg := "Hello World!"
		fmt.Println(a)
		fmt.Println(msg)
	*/

	/*
		3.2 简单类型
		控制：nil
		整型类型：int int8 int16 int32 int64 uint8 uint16 ...
		浮点数类型：float32, float64
		字节类型：byte （等价于 uint8）
		字符串类型：string
		布尔值类型：boolean
	*/
	/*
		var a8 int8 = 10
		var a1 byte = 'a'
		var b float32 = 12.2
		var str = "Hello World"
		ok := false
		fmt.Println(a8)
		fmt.Println(a1)
		fmt.Println(b)
		fmt.Println(str)
		fmt.Println(ok)
	*/

	/*
		3.3 字符串
		UTF8 编码
	*/
	/*
		str1 := "Golang"
		str2 := "Go语言"
		fmt.Println(reflect.TypeOf(str2[2]).Kind())
		fmt.Println(str1[2], string(str1[2]))
		fmt.Printf("%d %c \n", str2[2], str2[2])
		fmt.Println("len(str2): ", len(str2))
	*/
	/*
		字符串转 rune 数组
		字符串中每个字节都用 int32 表示，可处理中文
	*/
	/*
		str2 := "Go语言"
		runeArr := []rune(str2)
		fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
		fmt.Println(runeArr[2], string(runeArr[2]))
		fmt.Println("len(runeArr): ", len(runeArr))
	*/
	/*
	  3.4 数组 array 和 切片 slice
	*/
	var arr [5]int // 一维数组
	arr[0] = 1
	fmt.Println(arr)

	var arr2 [5][5]int // 二维数组
	arr2[0][0] = 1
	fmt.Println(arr2)

	ar := [5]int{0, 0, 0, 0, 0}
	ar[0] = 1
	fmt.Println(ar)

	ar2 := [5][5]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	ar2[0][0] = 1
	fmt.Println(ar2)
}
