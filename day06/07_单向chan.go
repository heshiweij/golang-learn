package day06

import "fmt"

// 单向管道
func UniDirectChan() {
	// 单向管道的定义
	//defination()

	// 生产者和消费者模式
	ProducerAndCustomer()
}

// 只写
func producer(chW chan<- int) {
	for i := 0; i < 10; i++ {
		chW<- i
	}

	close(chW) // 写完后关闭管道
}

func customer(chR <-chan int) {
	for data := range chR {
		fmt.Println("data => ", data)
	}
}

// 生产者和消费者模式
func ProducerAndCustomer() {
	ch := make(chan int)

	go producer(ch)

	customer(ch)
}

// 单向管道的定义
func defination() {
	// 定义一个无缓冲的双向 chan
	var ch chan int = make(chan int, 0)

	// 只读 ch
	var chR <-chan int = make(<-chan int)

	// 只写 ch
	var chW chan<- int = make(chan<- int)

	// 双向管道可以隐式转换为单向的 chan
	chR = ch
	chW = ch

	// right: 语法正确，但是会发生 deadlock
	//<-chR
	//chW <- a

	// wrong：语法错误
	//<-chW
	//chR<- 1

	fmt.Println(chR, chW) // 输出地址：管道是引用传递，和 map、slice 一样
}
