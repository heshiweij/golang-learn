package day05

import (
	"errors"
	"fmt"
)

// error 接口的使用
func ErrorUsing() {
	// 01 error 的基本概念和使用
	// errors 是包名
	// errorString 是一个实现了 error interface{ Error() string } 的结构体
	// Error() 方法返回了 string
	err := errors.New("出现错误了")

	if err != nil {
		fmt.Println("直接输出 err：", err) // ?? TODO 为啥这里直接能输出一个结构体
		fmt.Println("调用 err.Error() 输出", err.Error())
		//fmt.Printf("%T", err) // errors.errorString 结构体类型
	}

	// 02 需要自己处理的错误
	if result, err := getInfo(4, 2); err != nil {
		fmt.Println("计算出现错误了：", err)
	} else {
		fmt.Println("计算正确，结果为：", result)
	}

	// 03 panic 函数的使用
	// 错误处理方式，必须在可能发生的错误之前调用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	result := gatherInfo(10, 0)
	fmt.Println(result)
}

// 定义一个可处理返回错误的函数
func getInfo(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数不能为 0")
	} else {
		result = a / b
	}

	return
}

// 定义一个使用 panic 抛出异常的函数
func gatherInfo(a, b int) (result int) {
	if b == 0 {
		panic("除数不能为 0")
	}

	result = a / b
	return
}
