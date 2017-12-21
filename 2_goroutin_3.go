package main

import (
	"fmt"
	"time"
)

// for文のカウンタ変数をそのままgoroutinに突っ込んでみる
func main() {
	for count := 0; count < 10; count++ {
		go func() {
			fmt.Println("goroutin 1 task count - ", count)
		}()
	}

	// 3秒待つ
	time.Sleep(3 * time.Second)
}
