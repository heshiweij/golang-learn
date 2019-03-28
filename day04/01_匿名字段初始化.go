package day04

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	addr string
}

// 匿名字段初始化
func InitializeAnony() {
	// 顺序初始化
	var s = Student{Person{"hsw", 10}, "addr"}
	fmt.Printf("%+v \n", s)

	// 指定字段初始化
	var s2 = Student{Person: Person{"hsw", 10}, addr: "sy"}
	fmt.Printf("%+v \n", s2)

	// 指定字段初始化
	var s3 = Student{
		Person: Person{"hsw", 10},
		addr:   "sy", // 竖着写必须加 ","
	}
	fmt.Printf("%+v \n", s3)

	// 指定字段初始化
	var s4 = Student{
		Person: Person{
			name: "hsw",
			age:  10,
		},
		addr: "sy", // 竖着写必须加 ","
	}
	fmt.Printf("%+v \n", s4)



}
