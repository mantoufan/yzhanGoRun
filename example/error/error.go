package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	/* 返回错误：独立、明确的返回值来返回错误信息
	   好处：1 明确知道哪个函数返回错误
	   好处：2 调用那些没有出错的函数的一样调用 */
	// 使用 errors 的 errors.New
	for _, i := range []int{100, 101} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 errored：", e)
		} else {
			fmt.Println("f2 succed：", r)
		}
	}
	// 使用自定义错误
	for _, i := range []int{100, 101} {
		if r, e := f2(i); e != nil {
			fmt.Println("f1 errored：", e)
		} else {
			fmt.Println("f2 succed：", r)
		}
	}
	// 获取自定义错误实例
	_, e := f2(101)
	fmt.Println("e：", e)
	fmt.Println("e.(*argError)：", e.(*argError))
	if instance, ok := e.(*argError); ok {
		fmt.Println("instance.num：", instance.num)
		fmt.Println("instance.prob：", instance.prob)
	}
}

// 使用 errors 的 errors.New("error msg")
func f1(num int) (int, error) {
	if num > 100 {
		return -1, errors.New(strconv.Itoa(num) + " is exceed 100")
	}
	return num, nil
}

// 使用 Error 方法自定义 error 类型
type argError struct {
	num  int
	prob string
}

func (e *argError) Error() string {
	return e.prob
}

func f2(num int) (int, error) {
	if num > 100 {
		return -1, &argError{num, strconv.Itoa(num) + " is exceed 100"}
	}
	return num, nil
}
