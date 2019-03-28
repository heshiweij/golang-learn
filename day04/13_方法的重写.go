package day04

import "fmt"

type Person8 struct {
}

type Student8 struct {
	Person8
}

// 为父类绑定方法
func (this *Person8) getInfo() {
	fmt.Println("hello person")
}

// 为子类绑定方法
func (this *Student8) getInfo() {
	fmt.Println("hello student")
}

// 方法的重写
func StructFuncRewrite() {
	var p Person8
	p.getInfo() // hello person

	var s Student8
	s.getInfo()         // hello student（重写）
	s.Person8.getInfo() // hello person  指定调用父类的方法
}
