package main

import (
	"fmt"
	"os"
)

func main() {
	/* 延迟调用 */
	f := createFile("/tmp/defer.txt")
	defer closeFile(f) // 先写入再关闭
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("0 create")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("1 write")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("2 close")
	f.Close()
}
