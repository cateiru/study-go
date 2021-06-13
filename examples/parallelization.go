package main

import (
	"fmt"
	"time"
)

func main() {
	const total = 6 // 合計実行数
	const sync = 3  // 同時実行数

	sig := make(chan string, sync)
	res := make(chan string, total)

	// end of close
	defer close(sig)
	defer close(res)

	fmt.Println("Start!!")

	for i := 0; i < total; i++ { // 同時実行
		go wait(sig, res, fmt.Sprintf("no%d", i))
	}
	for {
		if len(res) >= total {
			break
		}
	}
}

func wait(sig chan string, res chan string, name string) {
	sig <- fmt.Sprintf("sig %s", name)
	time.Sleep(6 * time.Second)

	fmt.Printf("%s: end wait\n", name)
	res <- fmt.Sprintf("sig %s", name)
	<-sig
}
