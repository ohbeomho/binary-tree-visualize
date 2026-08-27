package main

type Node struct {
	key int
	left, right *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) insert(key int) {
	if tree.root == nil {
		tree.root = &Node{key: key}
		return
	}

	curr := &tree.root
	for *curr != nil {
		if key > (*curr).key {
			curr = &(*curr).right
		} else {
			curr = &(*curr).left
		}
	}

	*curr = &Node{key: key}
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

func (tree *BinaryTree) remove(node *Node) {
	// TODO: Implement it
}

func main() {
	tree := BinaryTree{}
}
