package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Node struct {
	key                 int
	left, right, parent *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) insert(key int) {
	if tree.root == nil {
		tree.root = &Node{key: key}
		return
	}

	curr := tree.root
	var prev *Node
	for curr != nil {
		prev = curr

		if key > curr.key {
			curr = curr.right
		} else {
			curr = curr.left
		}
	}

	newNode := &Node{key: key}
	if key > prev.key {
		prev.right = newNode
	} else {
		prev.left = newNode
	}

	newNode.parent = prev
}

func (tree BinaryTree) search(key int) (bool, []*Node) {
	curr := tree.root
	path := []*Node{curr}

	for curr != nil && curr.key != key {
		if key > curr.key {
			curr = curr.right
		} else if key < curr.key {
			curr = curr.left
		}

		path = append(path, curr)
	}

	if curr == nil {
		return false, path
	}

	return true, path
}

func (tree BinaryTree) getSuccessor(node *Node) *Node {
	if node.right == nil {
		return nil
	}

	curr := node.right
	for curr.left != nil {
		curr = curr.left
	}

	return curr
}

// TODO: Change it for visualization like insert function
func (tree *BinaryTree) remove(node *Node) {
	successor := tree.getSuccessor(node)

	if successor == nil {
		if node.left == nil {
			if node.key > node.parent.key {
				node.parent.right = nil
			} else {
				node.parent.left = nil
			}

			return
		}

		node.key = node.left.key

		node.left = node.left.left
		if node.left != nil {
			node.left.parent = node
		}

		node.right = node.left.right
		if node.right != nil {
			node.right.parent = node
		}

		return
	}

	node.key = successor.key

	if successor.key > successor.parent.key {
		successor.parent.right = successor.right
	} else {
		successor.parent.left = successor.right
	}

	if successor.right != nil {
		successor.right.parent = successor.parent
	}
}

func printTree(node *Node, level int) {
	if node == nil {
		return
	}

	fmt.Printf("%s%d\n", strings.Repeat("-", level), node.key)
	printTree(node.left, level+1)
	printTree(node.right, level+1)
}

func main() {
	tree := BinaryTree{}

	for range 20 {
		tree.insert(rand.Intn(100))
	}

	printTree(tree.root, 0)
}
