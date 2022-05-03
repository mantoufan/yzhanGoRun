package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	/* 匹配 */
	// regexp.Match + []byte()
	if ok, _ := regexp.Match("[0-9]", []byte("1")); ok {
		fmt.Println("ok")
	}
	// regexp.MatchString
	if ok, _ := regexp.MatchString("[0-9]", "1"); ok {
		fmt.Println("ok")
	}
	/* 解析 */
	// 返回 Regexp 对象指针
	rp, _ := regexp.Compile("[0-9]")

	/* 查找 */
	// 匹配第一个 byte
	first := rp.Find([]byte("abc123"))
	fmt.Println(string(first))

	// 匹配指定个数 byte
	alls := rp.FindAll([]byte("abc123"), 3)
	fmt.Println(alls)

	// 第一次匹配的开始位置和结束位置
	indexs := rp.FindIndex([]byte("abc123abc123"))
	fmt.Println(indexs)

	// 所有匹配的开始位置和结束位置
	indexAlls := rp.FindAllIndex([]byte("abc123abc123"), 2)
	fmt.Println(indexAlls)

	/* 捕获 */
	// 第一次匹配
	rp, _ = regexp.Compile("[a-z]+([0-9]+)[a-z]+(\\d+)")
	group := rp.FindSubmatch([]byte("abc123abc456"))
	fmt.Println(group)
	// 捕获分组 1 - 第一个括号
	fmt.Println(string(group[1]))
	// 捕获分组 2 - 第二个括号
	fmt.Println(string(group[2]))

	// 第一次匹配索引（开始，结束） 捕获分组 1 索引（开始，结束） 捕获分组 2 索引（开始，结束）
	groupIndex := rp.FindSubmatchIndex([]byte("abc123abc456"))
	fmt.Println(groupIndex)

	// 所有匹配
	rp, _ = regexp.Compile("[a-z]+([0-9]+)")
	groupAll := rp.FindAllSubmatch([]byte("abc123abc456"), 2)
	fmt.Println(groupAll)
	// 匹配1
	fmt.Println(string(groupAll[0][0]))
	// 匹配1 捕获1
	fmt.Println(string(groupAll[0][1]))
	// 匹配2
	fmt.Println(string(groupAll[1][0]))
	// 匹配2 捕获1
	fmt.Println(string(groupAll[1][1]))

	// 所有匹配索引（开始 结束） 匹配内捕获分组索引（开始 结束）
	groupAllIndex := rp.FindAllSubmatchIndex([]byte("abc123abc456"), 2)
	fmt.Println(groupAllIndex)

	/* 替换 */
	// 替换 byte
	rp, _ = regexp.Compile("[a-z]")
	rByte := rp.ReplaceAll([]byte("abc123"), []byte("z"))
	fmt.Println(rByte)

	// 替换字符串
	rString := rp.ReplaceAllString("abc123", "z")
	fmt.Println(rString)

	// 替换 byte 函数
	rByteFunc := rp.ReplaceAllFunc([]byte("abc123"), func(a []byte) []byte {
		return []byte(strings.ToUpper(string(a)))
	})
	fmt.Println(string(rByteFunc))

	// 替换字符串函数
	rStringFunc := rp.ReplaceAllStringFunc("abc123", strings.ToUpper)
	fmt.Println(string(rStringFunc))
}
