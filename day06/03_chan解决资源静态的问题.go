package day06

import (
	"fmt"
)

var ch = make(chan int)

// 协程资源竞争问题
func RoutineChannel() {
	// 同时开启两个协程，处理同一个资源，会产生资源竞争问题
	go person03()
	go person04()

	for {
		;
	}

}

func Printer01(content string) {
	for _, data := range content {
		fmt.Printf("%c", data)
	}

	fmt.Println()
}

func person03() {
	Printer01("hello")
	// 当执行完毕后，向管道推入 666
	ch <- 666
}

func person04() {
	// 无数据时，阻塞，有数据时，取出并继续执行
	a := <-ch
	Printer01("world")
	fmt.Println(a)
}
