package main

import "fmt"

func main() {
	fmt.Println("enter the number :")
	var num int
	fmt.Scanln(&num)
	resultStr := IntToString(num)
	resultRevers := ToRevers(resultStr)
	resultInt := StringToInt(resultRevers)
	fmt.Println(resultStr)
	fmt.Println(resultRevers)
	fmt.Println(resultInt)
}
func IntToString(n int) string {
	str := ""
	for n != 0 {
		r := n % 10
		str = string('0'+r) + str
		n = n / 10
	}
	return str
}
func ToRevers(ss string) string {
	s := []rune(ss)
	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	return string(s)
}
func StringToInt(s string) int {
	var num int
	for _, v := range s {
		num = num*10 + int(v-'0')
	}
	return num
}
