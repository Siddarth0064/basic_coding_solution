package main

import "fmt"

func main() {
	userp := person{
		name: "siddarth",
		age:  23,
	}

	fmt.Println(userp)
	userp.update("somashekar")
	fmt.Println(userp)

}

type person struct {
	name string
	age  int
}

func (p *person) update(s string) {
	p.name = s
}
