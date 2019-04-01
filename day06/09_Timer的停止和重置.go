package day06

import (
	"fmt"
	"time"
)

// Timer 的停止和重置
func TimerStoppingReset() {
	// 停止
	//stopTimer()

	// 重置
	resetTimer()

}

// 重置
func resetTimer() {
	timer := time.NewTimer(10 * time.Second)

	ok := timer.Reset(time.Second)

	if ok {
		fmt.Println("重置成功")
	}

	<-timer.C

	fmt.Println("结束")
}

// 停止
func stopTimer() {
	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C

		fmt.Println("时间到了")
	}()

	// 立即停止定时器
	//	- 加上这句后，不会输出 "时间到了"
	//timer.Stop()

	for {
		;
	}
}
