package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func init() {

}

func main() {

	if false {
		// 结构体的初始化
		var p1 = Person{
			"hsw",
			10,
		}

		var p2 = Person{
			Name: "hsw",
			Age:  10,
		}

		var p3 = &Person{ // 指针
			"hsw",
			10,
		}

		var p4 = &Person{ // 指针
			Name: "hsw",
			Age:  10,
		}

		fmt.Println(p1, p2, p3, p4)
	}

	if false {
		// 结构体是值传递

		var p Person = Person{"hsw", 10}
		func(p Person) { // inside
			p.Name = "wnm"
			p.Age = 20
			fmt.Println("inside: ", p.Name, p.Age)
		}(p)

		fmt.Println("inside: ", p.Name, p.Age) // outside

	}

	if false {
		// 匿名结构体
		// 01
		var p = struct {
			Name string
			Age  int
		}{
			Name: "hsw",
			Age:  10,
		}
		fmt.Println(p.Name, p.Age)

		// 02
		type Person struct {
			Name    string
			Age     int
			Address struct {
				lng  float64
				lat  float64
				Name string
			}
		}

		var p2 = Person{
			Name: "hsw",
			Age:  10,
			Address: struct {
				// 没有名字的结构体，只能这样指定
				lng  float64
				lat  float64
				Name string
			}{lng: 1, lat: 2, Name: "Shanghai"},
		}

		fmt.Println(p2)

		// 03
		var p3 = Person{
			Name: "hsw",
			Age:  10,
		}
		p3.Address.Name = "Shanghai" // 外部初始化匿名结构体字段更加的简单
		p3.Address.lat = 120.10
		p3.Address.lng = 30.23

		fmt.Println(p3)

	}

	if false {
		// 匿名字段
		type Person struct {
			string
			int
		}

		var p = Person{"hsw", 10} // 匿名字段赋值：顺序必须对应
		fmt.Println(p)

		var p2 = Person{ // 也可以使用类型名称来作为 key 指定赋值
			string: "hsw",
			int:    20,
		}
		fmt.Println(p2)
	}

	if false {
		// 结构体的比较

		type Person1 struct {
			Name string
			Age  int
		}

		var p1 = Person{"hsw", 10}
		var p2 = Person{"hsw", 10}
		var p3 = Person{"hsw", 20}
		//var p4 = Person1{"hsw", 11}

		fmt.Println(p1 == p2) // true
		fmt.Println(p2 == p3) // false
		fmt.Println(p1 == p3) // false
		//fmt.Println(p1 == p4) // wrong: 类型必须一致

	}

	if true {
		// 结构体组合
		type humen struct {
			sex int
		}

		type person struct {
			humen
			name string
			age  int
		}

		var p = person{name: "hsw", age: 10, humen: humen{1}}
		fmt.Println(p)

		var p2 = person{humen{1}, "hsw", 10}
		fmt.Println(p2)

		// 使用
		fmt.Println(p2.name, p2.age, p2.sex) // p2.sex 可以直接使用
	}


}
