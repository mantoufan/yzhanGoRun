package main

import "fmt"

func main() {
	/* 关联数组 / 集合 */
	// 声明 map
	var m map[string]int
	m = map[string]int{
		"a": 0,
		"b": 1,
	}
	fmt.Println(m)
	// 设置键值对
	m["c"] = 2
	// 获取长度
	fmt.Println(len(m))
	// 内建 delete 移除键值对
	delete(m, "a")
	// 判断键是否存在
	_, ok := m["c"]
	fmt.Println(ok)
	// 遍历 map
	for key, value := range m {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}
