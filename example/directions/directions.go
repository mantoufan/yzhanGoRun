package main

import "fmt"

func main() {
	/* 通道方向：
	   指定通道作为函数参数
		 只写通道 chan<-
		 只读通道 <-chan */
	writer, reader := make(chan string, 1), make(chan string, 1)
	// send msg to writer
	send(writer, "msg")
	// read msg from writer to reader
	receive(writer, reader)
	fmt.Println(<-reader)
}

// chan<- 仅接收，只写
func send(writer chan<- string, msg string) {
	writer <- msg
}

// <-chan 仅发送，只读， chan<- 仅接收，只写
func receive(reader <-chan string, writer chan<- string) {
	msg := <-reader
	writer <- msg
}
