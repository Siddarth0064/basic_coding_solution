package main

import "fmt"

func main() {
	val := 1234
	resultStr := toString(val)
	fmt.Println(resultStr)
	fmt.Printf("%T", resultStr)
}
func toString(a int) string {
	str := ""
	for a > 0 {
		r := a % 10
		str = string((r)+'0') + str
		a = a / 10
	}
	return str
}
