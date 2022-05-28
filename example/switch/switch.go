package main

import "fmt"

func main() {
	switchCase(1)
	switchCase(2)
	switchCase(3)
	switchCaseWithCondition(1)
	switchCaseWithCondition(2)
	switchCaseWithCondition(3)
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum("1", "2", "3"))
}

// 常量：匹配多个值
func switchCase(i int) {
	switch i {
	case 1, 2:
		fmt.Println("1 - 2")
	case 3:
		fmt.Println("3 ->")
		fallthrough
	case 4:
		fmt.Println("<- 4")
	}
}

// 条件
func switchCaseWithCondition(i int) {
	switch {
	case i == 1 || i == 2:
		fmt.Println("1 - 2")
	case i == 3:
		fmt.Println("3 ->")
		fallthrough
	case i == 4:
		fmt.Println("<- 4")
	}
}

// 空接口
func sum(param ...interface{}) interface{} {
	switch param[0].(type) {
	case int:
		return sumInt(param)
	case string:
		return sumString(param)
	default:
		fmt.Println("不支持相加的类型")
		return nil
	}
}

func sumInt(param []interface{}) (sum int) {
	for _, v := range param {
		sum += v.(int)
	}
	return
}

func sumString(param []interface{}) (sum string) {
	for _, v := range param {
		sum += v.(string)
	}
	return
}
