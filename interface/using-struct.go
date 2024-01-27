package main

import "fmt"

func main() {
	p := person{
		office:   "teksystem",
		vechical: "car",
	}
	fmt.Println(p)
	p.driving()
	p.update("microsoft")
	fmt.Println(p)
	p.driving()

}

type person struct {
	office   string
	vechical string
}
type traveling interface {
	driving()
}

func (pp person) driving() {
	fmt.Println("driving through", pp.vechical, "to ", pp.office, "office")
}
func (up *person) update(s string) {
	up.office = s

}
