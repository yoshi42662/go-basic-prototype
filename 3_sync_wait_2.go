package main

import (
	"fmt"
	"sync"
)

// sleepを挟まずに、WaitGroupを使ったgoroutin実行
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		// WaitGroupのカウンタをインクリメント
		wg.Add(1)

		go func(wg *sync.WaitGroup, count int) {
			// defer文はこのgoroutinの最後に実行される
			// DoneでWaitGroupのカウンタをデクリメント
			defer wg.Done()

			// for文の変数jをそのまま使用
			fmt.Println("goroutin 1 task count - ", count)
		}(&wg, i)
	}

	// WaitGroupのカウンタが0になるまで待つ
	wg.Wait()

	fmt.Println("finish without waiting.")
}
