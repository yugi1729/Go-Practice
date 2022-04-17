package main 

import (
	"fmt"
	"strconv"
)

type Node struct{
	value int
	next *Node
}

type LinkedList struct{
	head *Node
	tail *Node
}

func (list *LinkedList) addNode(value int){
	newNode := Node{value, nil}
	if list.head == nil{
		list.head = &newNode
		list.tail = &newNode
	}else{
		list.tail.next = &newNode
		list.tail = &newNode
	}
}

func (list *LinkedList) printList(){
	current := list.head
	output := ""
	for current != nil{
		if current.next != nil{
			output = output + strconv.Itoa(current.value) + "->"
		}else {
			output = output + strconv.Itoa(current.value)
		}
		current = current.next
	}
	fmt.Println(output)
}

func (list *LinkedList) deleteNode(){
	current := list.head
	for current.next != list.tail{
		current = current.next
	}
	current.next = nil
	list.tail = current
}

// driver function
func main(){
	list := LinkedList{nil, nil}
	list.addNode(1)
	list.addNode(2)
	list.addNode(3)
	list.addNode(4)
	list.addNode(5)
	list.printList()
	list.deleteNode()
	list.printList()
}