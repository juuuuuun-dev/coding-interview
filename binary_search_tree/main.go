package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if n.value > value {
		if n.left == nil {
			n.left = &Node{value: value}
			return
		}
		n.left.insert(value)
		return
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
			return
		}
		n.right.insert(value)
		return
	}
}

// Inorder left -> root -> right
func (node *Node) inorder() {
	if node != nil {
		node.left.inorder()
		fmt.Println(node.value)
		node.right.inorder()
	}
}

func (node *Node) search(value int) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	} else if node.value > value {
		return node.left.search(value)
	} else if node.value < value {
		return node.right.search(value)
	}
	return false
}

// nodeのleftがleftを。なれけば自身を返す。つまりそのnodeからなる最小を返す
func (node *Node) min() *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (node *Node) remove(value int) *Node {
	if node == nil {
		return nil
	}
	if node.value > value {
		node.left = node.left.remove(value)
	} else if node.value < value {
		node.right = node.right.remove(value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		//
		tmp := node.right.min()
		node.value = tmp.value
		node.right = node.right.remove(tmp.value)
	}
	return node
}

func main() {
	node := Node{value: 3}
	node.insert(6)
	node.insert(5)
	node.insert(7)
	node.insert(1)
	node.insert(10)
	node.insert(2)
	node.inorder()
	node.remove(6)
	fmt.Println(node.search(6))
	node.inorder()
}
