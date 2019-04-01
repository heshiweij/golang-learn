package day06

import (
	"fmt"
	"time"
)

// 有无缓存Chan的区别
func RoutineChenCache() {
	// 无缓存 chan
	//noChanCache()

	// 有缓存 chan
	chanCache()
}

// 有缓存 chan
func chanCache() {
	ch := make(chan int, 3)

	// len 为元素的个数，cap 为缓冲区大小
	fmt.Println("The len: ", len(ch), " The cap: ", cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("add ch ", i)
			// 有缓存chan，当推入数据后，不会被阻塞；缓冲区满了才会阻塞
			ch <- i
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 3; i++ {
		// 当缓冲区为空，会被阻塞
		val := <-ch
		fmt.Println("取出：", val)
	}
}

// 无缓存 chan
func noChanCache() {
	ch := make(chan int) // 相当于 make(chan int, 0)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("add ch ", i)
			// 无缓存chan，当推入数据后，会被阻塞
			ch <- i
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 3; i++ {
		// 当缓冲区为空，会被阻塞
		val := <-ch
		fmt.Println("取出：", val)
	}
}
