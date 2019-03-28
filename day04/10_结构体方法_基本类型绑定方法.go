package day04

import "fmt"

// 结构体方法 给基本类型绑定方法
func StructBasicTypeBindFunc() {
	var num long = 10
	result := num.add(11)
	fmt.Println("result: ", result)
}

type long int

// 给基本类型绑定方法
//	- 只能使用 type 别名后，才可以绑定
func (this long) add(num long) long {
	return this + num
}
