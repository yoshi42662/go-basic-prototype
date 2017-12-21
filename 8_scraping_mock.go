package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	// runtime側の並列数を設定
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	// runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	// 開始時間を記録
	start := time.Now()

	var wg sync.WaitGroup

	q := make(chan int)

	// CPUのコア数を取得
	cpuCount := runtime.NumCPU()
	fmt.Println("CPU count is " + fmt.Sprintf("%d", cpuCount))

	// ワーカー数の設定
	workerCount := cpuCount * 2
	// workerCount = 1

	// ワーカーの作成
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go fetchURL(&wg, q)
	}

	// 処理したいタスクの数を設定
	taskCount := 100

	// workerCountの値が同時処理数でタスクを処理する
	for index := 0; index < taskCount; index++ {
		fmt.Println("Queueing : ", index)
		q <- index
	}

	// channelでの処理の最後は必ずchannelをcloseする
	close(q)

	// すべてのgoroutineが終了するのをまつ
	wg.Wait()

	// 終了時間を計測
	end := time.Now()

	// 実行時間を算出
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func fetchURL(wg *sync.WaitGroup, q chan int) {
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

func mockAPIServer() bool {
	// Do something a bit heavy

	// 0ms ~ 200msの間でレスポンスを返す
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	return true
}
