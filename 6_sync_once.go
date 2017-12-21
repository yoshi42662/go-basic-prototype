package main

import (
	"fmt"
	"sync"
)

// onlyOnce()を1度だけ実行したい
func main() {
	var once sync.Once

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			// 1回のみonceFuncが実行される
			once.Do(onlyOnce)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

// 1回だけ実行したい関数を用意
func onlyOnce() {
	fmt.Println("Only once")
}
