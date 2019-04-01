package day06

import (
	"fmt"
	"runtime"
	"time"
)

// 协程的使用
func RoutineUsing() {
	//mainThread()

	//routineModel()

	//routineExitWhenMainExit()

	//routineNotStopWhenMainGoSche()

	//routineGoExitUsing()
}

func test() {
	// 01 输出 test defer
	//defer fmt.Println("defer")
	//fmt.Println("test")

	//02 输出 defer
	//defer fmt.Println("defer")
	//return
	//fmt.Println("test")

	// 03 输出 defer（函数中执行了 goexit，函数的 defer 还是会执行）
	defer fmt.Println("defer2")
	runtime.Goexit()
	fmt.Println("test")
}

// 使用 Go.Goexit() 终止当前所在的协程
func routineGoExitUsing() {

	go func() {
		fmt.Println("world1")
		test()
		fmt.Println("world2")
	}()

	for {
		;
	}
}

// 当主线程让出 CPU 时间片，则协程不会退出
func routineNotStopWhenMainGoSche() {
	go func() {
		for {
			fmt.Println("hello routine")
			time.Sleep(time.Second)
		}
	}()

	// 主线程让出时间片
	//	- 主线程会停下来，允许协程执行一段时间，然后主线程重新获取时间片，执行结束
	runtime.Gosched()

	fmt.Println("hello main")

}

// 主线程退出后，协程也退出
func routineExitWhenMainExit() {
	go func() {
		for {
			fmt.Println("hello routine")
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("hello main")
}

// 协程模式
func routineModel() {

	go func() {
		for {
			fmt.Println("hello routine")
			time.Sleep(time.Second)
		}
	}()

	for {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}

}

// 主线程阻塞
func mainThread() {
	// 单线程模式下，死循环会阻塞代码
	for {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}

	// 改代码不会被执行
	for {
		fmt.Println("hello2")
		time.Sleep(time.Second)
	}
}
