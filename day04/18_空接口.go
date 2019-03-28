package day04

import "fmt"

// 空接口
func InterfaceEmpty() {

	// 空接口可以接收任意类型的数据
	var i interface{}
	i = 1
	i = ""
	i = true
	i = func() {}
	i = new(Person10)
	i = &i
	i = map[int]string{1: "1", 2: "2", 3: "3"}
	i = Person10{""}
	i = &Person10{""}

	// 空接口的应用
	println(1,2,1.2,false,true,new(Person10), &i, func() {}, map[int]string{1:"ss"})

}

func println(args ...interface{}) {
	for i, data := range args {
		fmt.Println(i, data)
	}
}
