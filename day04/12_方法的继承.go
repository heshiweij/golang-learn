package day04

import "fmt"

type Person7 struct {
}

type Student7 struct {
	Person7
}

// 绑定方法
func (this *Person7) getInfo() {
	fmt.Println("hello person")
}

// 方法的继承
func StructFuncInhert() {
	var s Student7

	// 方法的继承
	s.getInfo()

}
