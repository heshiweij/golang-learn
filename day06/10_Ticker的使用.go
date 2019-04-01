package day06

import (
	"fmt"
	"time"
)

// Ticker 的使用
func UsingTicker() {
	ticker := time.NewTicker(time.Second)

	i := 0
	for {
		<-ticker.C

		i++
		fmt.Println("i = ", i)

		if i == 5 {
			ticker.Stop()
			//break
		}
	}
}
