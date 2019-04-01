package day06

import "fmt"

// 关闭chan
func CloseChan() {
	//usingFor()

	usingForRange()
}

// 使用 for range：自动判断管道关闭并退出循环
func usingForRange() {
	ch := make(chan int)

	defer fmt.Println("主线程停止")

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		// 关闭管道
		//	- 关闭后，不能再写数据
		//	- 关闭后，主线程立马会接收到，从而退出程序
		close(ch)
	}()

	for val := range ch {
		fmt.Println("读取到了：", val)
	}
}

// 使用 for：需要手动判断管道是否被关闭
func usingFor() {
	ch := make(chan int)

	defer fmt.Println("主线程停止")

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		// 关闭管道
		//	- 关闭后，不能再写数据
		//	- 关闭后，主线程立马会接收到，从而退出程序
		close(ch)
	}()

	for {
		// 有数据读取，没有数据阻塞
		if val, ok := <-ch; ok {
			// 管道没有被关闭
			fmt.Println("读取到数据：", val)

		} else {
			// 退出循环，进而退出所有协程
			fmt.Println("break")
			break
		}
	}
}
