package main

import "fmt"

func main() {
	str := "1234"
	result := toInt(str)
	fmt.Println(result)
}
func toInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = sum*10 + int(s[i]-'0')
	}
	return sum
}
