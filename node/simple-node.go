package main

import (
	//"container/list"
	"fmt"
)

type node struct {
	next *node
	data int
}
type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) endAppend(value int) {
	newNode := node{data: value}
	if l.head == nil {
		l.head = &newNode
		l.length++
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &newNode
	l.length++
	return
}
func (l *linkedlist) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
	return
}
func (l *linkedlist) printlist() {
	if l.head == nil {
		return
	} else {
		currentnode := l.head
		for currentnode != nil {
			fmt.Print(currentnode.data, " ")
			currentnode = currentnode.next
		}
	}
	return
}
func (l *linkedlist) deletWithvalue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
	}
	previoustoDelete := l.head
	for previoustoDelete.next.data != value {
		if previoustoDelete.next.data == value {
			return
		}
		previoustoDelete = previoustoDelete.next
	}
	previoustoDelete.next = previoustoDelete.next.next
	l.length--
}
func main() {
	mylist := &linkedlist{}
	fmt.Println("appending the elements in the front of the elements")
	mylist.prepend(1)
	mylist.prepend(2)
	mylist.prepend(3)
	mylist.prepend(8)
	mylist.printlist()
	fmt.Println("appending the elemets in the end")
	mylist.endAppend(5)
	mylist.endAppend(6)
	mylist.endAppend(7)
	mylist.endAppend(8)
	mylist.printlist()
	fmt.Println("deletinig the elements")
	mylist.deletWithvalue(8)
	for mylist.head != nil {
		fmt.Print(mylist.head.data, " ")
		mylist.head = mylist.head.next
	}
}
