package main

import (
	"fmt"
	"time"
)

// goroutingを試してみる
func main() {
	for k := 0; k < 100; k++ {
		go func(count int) {
			fmt.Println("goroutin task count - ", count)
		}(k)
	}

	// 3秒待つ
	time.Sleep(3 * time.Second)
}
