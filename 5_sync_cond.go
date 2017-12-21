package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Condを使用したgoroutin実行
func main() {
	// channel
	c := make(chan bool)

	m := new(sync.Mutex)
	cond := sync.NewCond(m)

	// cond.Wait()を仕込んだgoroutinを5つ用意
	for i := 0; i < 5; i++ {
		j := i

		go func() {
			cond.L.Lock()
			defer cond.L.Unlock()

			cond.Wait()
			fmt.Println(j)
			c <- true
		}()
	}

	// 3秒待機
	time.Sleep(3 * time.Second)

	// sync.Condの実行シグナルをコール

	// パターン1 -> Signal()を使用する場合
	for i := 0; i < 5; i++ {
		// goroutinの数だけ、Signal()を呼び出す必要がある
		go func() {
			cond.Signal()
		}()

		<-c
	}

	// パターン2 -> Broadcast()を使用する場合
	// cond.Broadcast()
	// <-c
	// <-c
	// <-c
	// <-c
	// <-c
}
