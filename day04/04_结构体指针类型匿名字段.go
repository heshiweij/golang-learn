package day04

import "fmt"

// 结构体指针类型匿名字段
type Student4 struct {
	*Person4
	score int
}

type Person4 struct {
	name string
	age  int
}

// 结构体指针类型的匿名字段
func StructAnonyPointerField() {
	// 指针类型匿名字段初始化
	var s = Student4{&Person4{"", 10}, 10}

	fmt.Println(s)

	var s2 Student4
	s2.Person4 = new(Person4)
	s2.name = "hsw" // 直接访问
	s2.Person4.name = "hsw" // 和上面一样
	s2.age = 10
	s2.score = 100

	fmt.Println(s2)

}
