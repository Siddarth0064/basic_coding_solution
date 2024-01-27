package main

import (
	"fmt"
	"sort"
)

func main() {
	strArr := []string{"apple", "cat", "zeebra", "ball", "elephant"}
	fmt.Println(strArr, " befor sorting")
	sort.Strings(strArr)
	fmt.Println(strArr, " after sorting array")
}
