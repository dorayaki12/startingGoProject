package main

import (
	"fmt"
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break // 受信できなくなったら終了
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "is done")
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 3)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	// 複数のcaseが成立する場合はランダムに選択される
	select {
	case <-ch1:
		fmt.Println("ch1から受信")
	case <-ch2:
		fmt.Println("ch2から受信")
	case <-ch3:
		fmt.Println("ch3から受信")
	default:
		fmt.Println("ここには到達しない")
	}
}
