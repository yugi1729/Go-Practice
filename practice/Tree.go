package main

import (
	"fmt"
)

type Node struct {
	data int
	left *Node
	right *Node
}

func addNode(root *Node, data int) *Node {
	if root == nil {
		root = &Node{data, nil, nil}
		return root
	} else if data < root.data {
		root.left = addNode(root.left, data)
		return root
	} else {
		root.right = addNode(root.right, data)
		return root
	}
}

func displayTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	displayTree(root.left)
	displayTree(root.right)
}



func main(){
	root := &Node{8, nil, nil}
	addNode(root, 8)
	addNode(root, 9)
	addNode(root, 1)
	addNode(root, 2)
	displayTreeLevel(root)
}

