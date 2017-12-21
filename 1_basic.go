// パッケージ名 / package name
package main

// パッケージのインポート / importing package
import (
	"fmt"  // fmt.MethodNameで使用可能 / available by fmt.MethodName
	"time" // time.MethodNameで使用可能 / available by time.MethodName
)

// 定数 / constant
const (
	ExportedConst   = 1
	unexportedConst = 2
)

// 変数 / variable
var (
	ExportedVar   = 1
	unexportedVar = 2
)

// イニシャライザ関数 / initializer
func init() {
	// 遅延実行 defer
	defer func() {
		fmt.Println("2. defer at init()")
	}()

	// 出力 / output
	fmt.Println("1. initializing")
}

// メイン関数 / main function
func main() {
	// ローカル変数の宣言代入(型は自動判定)
	localVar := "this is localVar!"

	// ローカル変数の再代入
	localVar = "hello world!"

	// 出力
	fmt.Println("3.", localVar)

	// main関数のdefer 1
	defer func() {
		fmt.Println("7. first defer at main")
	}()

	// main関数のdefer 2
	defer func() {
		fmt.Println("6. second defer at main")
	}()

	// ゴルーチン
	go func() {
		fmt.Println("4. executing from goroutin")
	}()

	// 3秒待機
	time.Sleep(2 * time.Second)

	fmt.Println("5. finished main func!")
}
