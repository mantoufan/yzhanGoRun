package main
import (
	"fmt"
	"strconv"
)

func main () {
	/*
	 * strconv 包定义了两个 error 类型变量
	 * ErrRange 值超过类型表示最大范围
	 * ErrSyntax 语法错误
	 */
	type NumError struct {
		Func string // failing function (ParseBool ParseInt ParseUint ParseFloat)
		Num string // input
		Err error // reason the conversion faild (ErrRange, ErrSyntax)
	}
	/* Error() 接口 */
	func (e *NumError) Error() string {
		return "strconv." + e.Func + ":" + "parsing" + Quote(e.Num) + ":" + e.Err.Error()
	}
	/* 快捷函数，用于构造 NumError 对象 */
	func syntaxError(fn, str string) *NumError {
		return &NumError{fn, str, ErrSyntax}
	}
	func rangeError(fn, str string) *NumError {
		return &NumError{fn, str, ErrRange}
	}
}