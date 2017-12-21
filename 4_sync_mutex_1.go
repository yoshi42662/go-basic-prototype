package main

import (
	"fmt"
)

var a string

// sync.Mutexが無い場合の変数代入
func main() {
	// 変数aに値をセット
	a = "this is variable a"

	// goroutinで変数aの書き換え
	go f()

	// 変数aを出力
	fmt.Println(a)
}

func f() {
	// 変数aに値をセット
	a = "hello, world!"
}
