package day06

import (
	"fmt"
	"time"
)

// Timer 的停止和重置
func TimerStoppingReset() {
	stopTimer()

}

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
