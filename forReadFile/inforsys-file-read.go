package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("aaa.txt")
	if err != nil {
		panic(err)
	}
	tempSlice := make([]byte, 500)
	numberOfBytes, err := file.Read(tempSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The contents of the file are as follows: \n%s ",
		string(tempSlice[:numberOfBytes]))
	file.Close()
}
