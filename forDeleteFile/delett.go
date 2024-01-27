package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "a.txt"
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Println("problem in file ")
		return
	}
	fmt.Println("succsucesfuly removed ", filename)

}
