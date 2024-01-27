package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "im am siddarth"
	find := "as"
	resultint := searchStr(str, find)
	if resultint != -1 {
		fmt.Println("its presnt")
	} else {
		fmt.Println("its presnt")
	}
}
func searchStr(s, f string) int {

	sw := strings.Split(s, " ")
	fw := strings.Split(f, " ")
	sl := len(sw)
	fl := len(fw)
	for i := 0; i < sl-fl; i++ {

	}

}
