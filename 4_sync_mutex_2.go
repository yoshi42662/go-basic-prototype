package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex
var a string

// sync.Mutexを使用した場合の変数代入
func main() {
	// 変数aに値をセット
	a = "this is variable a"

	// Mutex Lock
	l.Lock()

	// goroutinで変数aの書き換え
	go f()

	// Mutex Lock
	l.Lock()
	/* ここでl.Lock()を呼び出すことで、var l sync.Mutex が Unlockされるまで待つようになる */

	// 変数aを出力
	fmt.Println(a)
}

func f() {
	// 変数aに値をセット
	a = "hello, world!"

	// Mutex Unlock
	l.Unlock()
}
