package main

import "fmt"

//Node
type Node struct {
	key   int
	left  *Node
	right *Node
}

//Insert
func (n *Node) insertBinary(key int) {
	if key > n.key {
		//move right
		if n.right != nil {
			n.right.insertBinary(key)
		} else {
			n.right = &Node{key: key}
		}
	} else if key < n.key {
		//move left
		if n.left != nil {
			n.left.insertBinary(key)
		} else {
			n.left = &Node{key: key}
		}
	}
}

//Search
func (n *Node) searchBinary(key int) bool {
	if n == nil {
		return false
	}

	if key > n.key {
		//move right
		return n.right.searchBinary(key)
	} else if key < n.key {
		//move left
		return n.left.searchBinary(key)
	}
	return true
}

func BinaryTraversal(root *Node) error {

	//Pre Order Traversal
	// fmt.Print(root.key)
	// fmt.Print(" ")

	if root.left != nil {
		BinaryTraversal(root.left)
	}

	//In Order Traversal
	fmt.Print(root.key)
	fmt.Print(" ")

	if root.right != nil {
		BinaryTraversal(root.right)
	}

	//Post Order Traversal
	// fmt.Print(root.key)
	// fmt.Print(" ")

	return nil
}

func MainBinary() {
	tree := &Node{key: 20}
	tree.insertBinary(25)
	tree.insertBinary(15)
	tree.insertBinary(30)
	tree.insertBinary(13)
	tree.insertBinary(18)

	BinaryTraversal(tree)
	fmt.Println(tree.searchBinary(145))
}
