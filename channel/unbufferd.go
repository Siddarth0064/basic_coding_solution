package main

import "fmt"

type cal interface {
	adding()
}
type person struct {
	name   string
	age    int
	marks1 int
	marks2 int
}

func (p *person) update(name string) {
	p.name = name
}

func (p person) adding() {
	fmt.Println(p.marks1 + p.marks2)
	panic("panic crated in adding")

}
func main() {
	pp := person{
		name:   "sidd",
		age:    22,
		marks1: 23,
		marks2: 2,
	}
	pp1 := person{
		name:   "sooo",
		age:    33,
		marks1: 66,
		marks2: 22,
	}
	defer recover()
	fmt.Println(pp)
	fmt.Println(pp1)
	pp.adding()
	pp.update("dddkkk")
	fmt.Println(pp)
}

func recover() {
	fmt.Println("its a recoverd from recover")
}
