package main

import (
	//"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("abcd.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	//defer file.Close()
	// read, err := os.ReadFile("file")
	// if err != nil {
	// 	fmt.Println("ther is a erroe in reading")
	// }
	// fmt.Println(read)
	data := []byte("The sky is blue.and im write from vs code")
	n, err := file.Write(data)
	if err != nil {
		fmt.Println("Error in writing to a file:", err)
		return
	}
	fmt.Println("No. of bytes written to file:", n)
	fmt.Println("Data written successfully!")
	// readfile := bytes.ToString(file)
	// fmt.Println(readfile)
}
