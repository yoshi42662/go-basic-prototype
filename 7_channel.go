package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	// 送信
	go func() { messages <- "ping" }()

	// 受信
	msg := <-messages

	fmt.Println(msg)
}
