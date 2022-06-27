package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	/* 字符串 */
	// 包含
	p("Contains：", s.Contains("test", "es"))
	// 计数
	p("Count：", s.Count("test", "t"))
	// 前缀
	p("HasPrefix：", s.HasPrefix("test", "te"))
	// 后缀
	p("HasSuffix：", s.HasSuffix("test", "st"))
	// 索引位置
	p("Index：", s.Index("test", "e"))
	// 重复字符
	p("Repeat：", s.Repeat("a", 5))
	// 替换字符
	p("Replace：", s.Replace("foo", "o", "0", -1))
	// 替换字符，替换指定次数
	p("Replace：", s.Replace("foo", "o", "0", 1))
	// 控制
	p("Split：", s.Split("a-b-c-d-e", "-"))
	// 小写
	p("ToLower：", s.ToLower("TEST"))
	// 大写
	p("ToUpper：", s.ToUpper("test"))

	// 字符串长度
	p("Len：", len("test"))
	// 访问字符串中的 Char
	p("Char：", "test"[1])
}
