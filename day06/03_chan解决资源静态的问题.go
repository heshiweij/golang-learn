package day06

import (
	"fmt"
	"time"
)

var a = 1

// 协程资源竞争问题
func RoutineResCompete() {
	// 同时开启两个协程，处理同一个资源，会产生资源竞争问题
	//runtime.GOMAXPROCS(4)
	//go person01()
	//go person02()
	//
	//for {
	//	;
	//}

	// 02 协程可以直接访问和修改全局变量的数据，而不是拷贝

	go func() {
		time.Sleep(time.Second * 2)
		a = 2
	}()

	go test2()

	for {
		fmt.Println("a = ", a)
		time.Sleep(time.Second)
	}

}

func test2() {
	a = 3
}

func Printer(content string) {
	for _, data := range content {
		fmt.Printf("%c", data)
	}

	fmt.Println()
}

func person01() {
	Printer("hello")
}

func person02() {
	Printer("world")
}
