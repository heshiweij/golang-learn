package day06

import (
	"fmt"
	"time"
)

// Timer 的使用
func UsingTimer() {
	// 01
	normalTimer()

	// 02
	simpleTimer()

	// 03
	sleep01()

	// 02
	sleep02()
}

func sleep02() {
	// 直接返回管道
	<-time.After(2 * time.Second)
}

func sleep01() {
	time.Sleep(2 * time.Second)
}

func simpleTimer() {
	<-time.NewTimer(2 * time.Second).C
}

func normalTimer() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}
