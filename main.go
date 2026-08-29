package main

import (
	"fmt"
	"unsafe"
)

// root = 1
// left child = index * 2
// right child = index * 2 + 1
var tree []int32
var initialized = false

// Figure out how to return array

//go:wasmexport treeInit
func initTree(levels int32) unsafe.Pointer {
	var treeSize int32

	for i := range levels + 1 {
		treeSize += 2 ^ i
	}

	tree = make([]int32, treeSize+1)

	initialized = true
	return unsafe.Pointer(&tree[0])
}

//go:wasmexport treeInsert
func insert(key int32) {
	if !initialized {
		return
	}

	// 0을 빈 값으로 사용하기 위함
	if key == 0 {
		return
	}

	curr := 1
	for tree[curr] != 0 {
		if key > tree[curr] {
			curr = curr*2 + 1
		} else {
			curr = curr * 2
		}
	}

	tree[curr] = key
}

//go:wasmexport treeSearch
func search(key int32) int32 {
	if !initialized {
		return 0
	}

	if key == 0 {
		return 0
	}

	var curr int32 = 1
	for tree[curr] != 0 && tree[curr] != key {
		if key > tree[curr] {
			curr = curr*2 + 1
		} else {
			curr = curr * 2
		}
	}

	if tree[curr] == 0 {
		return 0
	}

	return curr
}

func main() {
	fmt.Println("test")
}
