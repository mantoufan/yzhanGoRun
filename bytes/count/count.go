package main

import (
	"bytes"
	"fmt"
)

func main() {
	/* 出现次数无重叠 */
	fmt.Println(bytes.Count([]byte("abababa"), []byte("aba")))
}
