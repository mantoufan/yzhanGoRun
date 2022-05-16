package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* 输出带双引号的字符串 */
	fmt.Println(`a"b"c`)
	fmt.Println("a\"b\"c")
	fmt.Println("a" + strconv.Quote("b") + "c")
	fmt.Println(strings.Join([]string{"a", strconv.Quote("b"), "c"}, ""))
	var buffer bytes.Buffer
	buffer.WriteString("a")
	buffer.WriteString(strconv.Quote("b"))
	buffer.WriteString("c")
	fmt.Println(buffer.String())
}
