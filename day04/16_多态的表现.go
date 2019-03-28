package day04

import "fmt"

type Runnable interface {
	run()
}

type Person10 struct {
	name string
}

type Dog struct {
	name string
}

func (this Person10) run() {
	fmt.Println("person run")
}

func (this Dog) run() {
	fmt.Println("dog run")
}

func makeRun(i Runnable) {
	if  i != nil {
		i.run()
	}
}

// 多态的表现
func PolyDefine() {
	// 01
	p := Person10{"person10"}
	makeRun(p)

	d := Dog{"dog"}
	makeRun(d)

	p1 := Person10{"p10"}


	fmt.Println("============")

	// 02
	var slice = make([]Runnable, 0, 3)
	slice = append(slice, p, d, p1)

	for _, data := range slice {
		data.run()
	}
}
