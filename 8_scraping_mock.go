package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	// CPUのコア数を取得
	cpuCount = runtime.NumCPU()

	// runtime側の並列数
	maxProcsCount = cpuCount * 2

	// goroutinのworker数
	// workerCount = cpuCount * 16
	workerCount = 1

	// 処理したいタスクの数を設定
	taskCount = 100
)

func init() {
	// runtime側の並列数を設定
	runtime.GOMAXPROCS(maxProcsCount)
}

// 0ms ~ 200msの間でレスポンスを返す
func mockAPIServer() bool {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	return true
}

// 擬似スクレイピング用の関数
func mockFetchURL(wg *sync.WaitGroup, q chan int) {
	// defer文はfetchURL関数のブロックを抜ける際に実行される
	defer wg.Done()

	for {
		// channelがcloseされると ok が false になる
		index, ok := <-q
		if !ok {
			return
		}

		// URLを叩いてみる
		_ = mockAPIServer()

		// Output result
		fmt.Println("Done : ", index)
	}
}

func main() {
	// 開始時間を記録
	start := time.Now()

	var wg sync.WaitGroup

	// int型のchannelを用意
	q := make(chan int)

	// goroutin workerの生成
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go mockFetchURL(&wg, q)
	}

	// workerCountの値が同時処理数でタスクを処理する
	for index := 0; index < taskCount; index++ {
		fmt.Println("Queueing : ", index)

		// indexをqに送信
		q <- index
	}

	// channelでの処理の最後は必ずchannelをcloseする
	close(q)

	// すべてのgoroutineが終了するのをまつ
	wg.Wait()

	// 終了時間を計測
	end := time.Now()

	// 実行時間を算出
	fmt.Printf("\n\n実行時間: %f秒\n\n", (end.Sub(start)).Seconds())
}
