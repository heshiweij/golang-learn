package day04

import "fmt"

type Person2 struct {
	name string
	age  int
}

type Student2 struct {
	p Person2
	addr string
	int
	string
	float32
	name string
}

type Test struct {
	name string
	age int
}

// 非结构体匿名字段
func NoneStructAnonyField() {

	// 指定字段赋值
	var s Student2 = Student2{
		p: Person2{name:"", age:0},
		addr:"",
		int: 1,
		string: "string",
		float32: 1.2,
	}

	fmt.Printf("%+v \n", s)

	// 非结构体成员字段的访问
	//	- 直接通过匿名字段的类型访问
	fmt.Println(s.name, s.int, s.addr, s.float32)

	// 按顺序赋值
	var s2 Student2 = Student2{
		Person2{"", 0},
		"",
		1,
		"string",
		1.2,
		"fs",
	}

	fmt.Printf("%+v \n", s2)

	// 顺序赋值一定要全部赋值吗？
	//	- 是的
	t := Test{"", 1}
	fmt.Printf("%+v \n", t)

}
