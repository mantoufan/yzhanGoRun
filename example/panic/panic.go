package main

import "os"

func main() {
	/* 抛出错误：意料之外的错误
	   输出：
		 ① 错误消息
		 ② Go 运行时的栈信息
		 ③ 返回一个非零的状态码
	*/
	panic("a problem")

	_, err := os.Create("/tmp/data/01")
	if err != nil {
		panic(err)
	}
}
