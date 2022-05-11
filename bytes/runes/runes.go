package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("abc中文")
	for i, v := range b {
		fmt.Println(i, string(v))
	}
	r := bytes.Runes(b)
	for i, v := range r {
		fmt.Println(i, string(v))
	}
}
