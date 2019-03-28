package day04

import "fmt"

// 成员的操作
func OperateAnony() {
	var s Student = Student{
		Person{
			name: "he",
			age:  10,
		},
		"s",
	}

	// 普通成员的操作
	fmt.Println(s.name) // 调用的是 student 的 name，如果不存在则调用 Person 的 name

	// 同名成员的操作
	fmt.Println(s.Person.name) // 调用的是 s.Person 的 name
}
