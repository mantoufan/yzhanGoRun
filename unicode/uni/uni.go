package main

import (
	"fmt"
	"unicode"
)

func main() {
	/* Unicode 码点
	   unicode 是基本的字符判断函数
	   go 语言所有代码都是 UTF8
	   单个字符（单引号括起来）是 unicode */

	/* 数字 */
	fmt.Println(unicode.IsControl('1'))
	// func IsDigit(r rune) bool 是否阿拉伯数字字符，即 0 - 9
	fmt.Println(unicode.IsDigit('1'))
	fmt.Println(unicode.IsDigit('Ⅷ'))
	// func IsNumber(r rune) bool 是否数字字符
	fmt.Println(unicode.IsNumber('Ⅷ'))

	/* 字母 */
	// func IsLetter(r rune) bool 是否字母
	fmt.Println(unicode.IsLetter('a'))
	// func IsLower(r rune) bool 是否小写字母
	fmt.Println(unicode.IsLower('a'))
	// func IsUpper(r rune) bool 是否大写字母
	fmt.Println(unicode.IsUpper('A'))
	// func IsTitle(r rune) bool 是否 Title Case
	// What is Titlecase Letter: https://www.compart.com/en/unicode/category/Lt
	fmt.Println(unicode.IsTitle('ῌ'))
	fmt.Println(unicode.IsTitle('ǈ'))

	/* 符号 */
	/* 控制字符 */
	// func IsControl(r rune) bool 是否控制字符
	fmt.Println(unicode.IsControl('\u0015'))
	fmt.Println(unicode.IsControl('\ufe35'))
	/* 标记字符 */
	// func IsMark(r rune) bool 是否标记字符
	fmt.Println(unicode.IsMark('1'))
	fmt.Println(unicode.IsMark('̀')) // 声调
	/* 可打印字符 */
	// func IsPrint(r rune) bool 是否可打印字符
	fmt.Println(unicode.IsPrint('1'))
	fmt.Println(unicode.IsPrint('̀')) // 声调
	/* 标点符号 */
	// func IsPunct(r rune) bool 是否 Unicode 标点符号
	fmt.Println(unicode.IsPunct(','))
	fmt.Println(unicode.IsPunct('\\'))
	fmt.Println(unicode.IsPunct('/'))
	fmt.Println(unicode.IsPunct('&'))
	/* 数学、货币、修饰（音调）和其它有效符号 */
	// func IsSymbol(r rune) bool 是否是一个象征性符号
	fmt.Println(unicode.IsSymbol('+'))
	fmt.Println(unicode.IsSymbol('$'))
	fmt.Println(unicode.IsSymbol(',')) // False
	fmt.Println(unicode.IsSymbol('̀')) // 声调 False
	/* 空格 */
	fmt.Println(unicode.IsSpace(' ')) // 半角空格
	fmt.Println(unicode.IsSpace('　')) // 全角空格

	/* 汉字 */
	// func Is(rangeTab *RangeTable, r rune) bool r 是否为 rangeTab 类型的字符
	fmt.Println(unicode.Is(unicode.Han, '中'))
	// func In(r rune, ranges ...*RangeTable) bool r 是否为 ranges 任意一个类型的字符
	fmt.Println(unicode.In('中', unicode.Han))

	/* 组合：是否在 汉字 / 空格 中 */
	fmt.Println(unicode.In('中', unicode.Han, unicode.Space))
	fmt.Println(unicode.In(' ', unicode.Han, unicode.Space))
	fmt.Println(unicode.In('　', unicode.Han, unicode.Space))
}
