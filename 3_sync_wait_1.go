package main

import (
	"fmt"
)

// sleepせずにgoroutinを実行してみる
func main() {
	for i := 0; i < 100; i++ {
		j := i

		go func() {
			// for文の変数jをそのまま使用 (Bad practice)
			fmt.Println("goroutin 1 task count - ", j)
		}()
	}

	fmt.Println("finish without waiting.")
}
