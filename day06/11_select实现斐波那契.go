package day06

import (
	"fmt"
)

// select实现斐波那契
func UsingSelectFibonacci() {

	channel := make(chan int)
	quit := make(chan bool)

	go func() {

		x, y := 1, 1
		for {

			fmt.Println("读取了")

			select {
			case channel <- x:
				x, y = y, x+y
			case <-quit:
				return
			}
		}

	}()

	// 只取 8 个
	for i := 0; i < 1; i++ {
		fmt.Println("取出：", <-channel) // chan为空则阻塞
	}
	quit <- true // 可以停止
}

func fibonacci(channel chan int, quit chan bool) {
	// 产生斐波那契

}
