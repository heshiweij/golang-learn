package main

import "fmt"

func init() {

}

type Runnable interface {
	run()
	say()
}

type Person struct {
	name string
}

func (p Person) run() {
	fmt.Println("run: ", p.name)
}

func (p Person) say() {
	fmt.Println("say: ", p.name)
}

func main() {
	if false {
		var p Person = Person{"haha"}
		p.run()
	}

	if false {
		var r Runnable = &Person{"haha"}
		r.run()
	}

	if true {
		var r Runnable = Person{"haha"}
		r.say()
		r.run()
	}

}
