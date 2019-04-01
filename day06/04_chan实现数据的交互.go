package day06

import "fmt"

// chan 实现数据的交互
func ChanDataShare() {
	ch := make(chan int)

	defer fmt.Println("主线程结束")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i = ", i)
		}
		ch <- 666
	}()

	val := <-ch
	fmt.Println("主协程接收到了信号：", val)
}
