package main

import "fmt"
import "reflect"

func main() {
	// var a int // 没有赋值，默认为 0 
	// var a int  = 1 // 声明时赋值
	// var a = 1 // 声明时赋值
	// a := 1
	// msg := "Hello World!"
	// fmt.Println(a)
	// fmt.Println(msg)

	// var  a8 int8 = 10
	// var  a1 byte = 'a'
	// var   b float32 = 12.2
	// var str = "Hello World"
	// ok := false
	// fmt.Println(a8)
	// fmt.Println(a1)
	// fmt.Println(b)
	// fmt.Println(str)
	// fmt.Println(ok)

	// str1 := "Golang"
	// str2 := "Go语言"
	// fmt.Println(reflect.TypeOf(str2[2]).Kind())
	// fmt.Println(str1[2], string(str1[2]))
	// fmt.Printf("%d %c \n", str2[1], str2[1])
	// fmt.Println("len(str2): ", len(str2))

	str2 := "Go语言"
	runeArr := []rune(str2)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	fmt.Println(runeArr[2], string(runeArr[2]))
	fmt.Println("len(runeArr): ", len(runeArr))
}