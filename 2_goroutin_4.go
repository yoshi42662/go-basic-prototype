package main

import (
	"fmt"
	"time"
)

// 逐次処理とgoroutinの比較
func main() {
	// 逐次処理
	for i := 0; i < 10; i++ {
		fmt.Println("sequential - ", i)
	}

	// goroutin
	for count := 0; count < 10; count++ {
		go func(cnt int) {
			fmt.Println("goroutin - ", cnt)
		}(count)
	}

	// 1秒待つ
	time.Sleep(1 * time.Second)
}
